package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"dario.cat/mergo"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/ethpandaops/dora/config"
	"github.com/ethpandaops/dora/types"
)

// Config is the globally accessible configuration
var Config *types.Config

// ReadConfig will process a configuration
func ReadConfig(cfg *types.Config, path string) error {
	err := readConfigFile(cfg, path)
	if err != nil {
		return err
	}

	readConfigEnv(cfg)

	var chainConfig types.ChainConfig
	if cfg.Chain.ConfigPath == "" {
		switch cfg.Chain.Name {
		case "mainnet":
			err = yaml.Unmarshal([]byte(config.MainnetChainYml), &chainConfig)
		case "goerli", "prater":
			err = yaml.Unmarshal([]byte(config.PraterChainYml), &chainConfig)
		case "sepolia":
			err = yaml.Unmarshal([]byte(config.SepoliaChainYml), &chainConfig)
		case "holesky":
			err = yaml.Unmarshal([]byte(config.HoleskyChainYml), &chainConfig)
		case "gnosis":
			err = yaml.Unmarshal([]byte(config.GnosisChainYml), &chainConfig)
		case "creeper2":
			err = yaml.Unmarshal([]byte(config.Creeper2ChainYml), &chainConfig)
		default:
			return fmt.Errorf("tried to set known chain-config, but unknown chain-name")
		}
		if err != nil {
			return err
		}

	} else {
		var reader io.Reader
		if strings.HasPrefix(cfg.Chain.ConfigPath, "http://") || strings.HasPrefix(cfg.Chain.ConfigPath, "https://") {
			client := &http.Client{Timeout: time.Second * 120}
			resp, err := client.Get(cfg.Chain.ConfigPath)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("url: %v, result: %v %v", path, resp.StatusCode, resp.Status)
			}
			reader = resp.Body
		} else {
			f, err := os.Open(cfg.Chain.ConfigPath)
			if err != nil {
				return fmt.Errorf("error opening Chain Config file %v: %w", cfg.Chain.ConfigPath, err)
			}
			defer f.Close()
			reader = f
		}
		decoder := yaml.NewDecoder(reader)
		err = decoder.Decode(&chainConfig)
		if err != nil {
			return fmt.Errorf("error decoding Chain Config file %v: %v", cfg.Chain.ConfigPath, err)
		}
	}

	// load preset if PresetBase is set
	if chainConfig.PresetBase != "" {
		var chainPreset types.ChainConfig
		switch chainConfig.PresetBase {
		case "mainnet":
			err = yaml.Unmarshal([]byte(config.MainnetPresetYml), &chainPreset)
		case "minimal":
			err = yaml.Unmarshal([]byte(config.MinimalPresetYml), &chainPreset)
		case "gnosis":
			err = yaml.Unmarshal([]byte(config.GnosisPresetYml), &chainPreset)
		case "creeper2":
			err = yaml.Unmarshal([]byte(config.Creeper2PresetYml), &chainPreset)
		default:
			return fmt.Errorf("tried to use unknown chain-preset: %v", chainConfig.PresetBase)
		}
		if err != nil {
			return err
		}

		err := mergo.Merge(&chainPreset, chainConfig, mergo.WithOverride)
		if err != nil {
			return fmt.Errorf("error merging chain preset: %v", err)
		}
		cfg.Chain.Config = chainPreset
	} else {
		cfg.Chain.Config = chainConfig
	}

	cfg.Chain.Name = cfg.Chain.Config.ConfigName

	if cfg.Chain.GenesisTimestamp == 0 {
		switch cfg.Chain.Name {
		case "mainnet":
			cfg.Chain.GenesisTimestamp = 1606824023
		case "goerli", "prater":
			cfg.Chain.GenesisTimestamp = 1616508000
		case "sepolia":
			cfg.Chain.GenesisTimestamp = 1655733600
		case "creeper2":
			cfg.Chain.GenesisTimestamp = 1709016906
		default:
			cfg.Chain.GenesisTimestamp = uint64(cfg.Chain.Config.MinGenesisTime) + cfg.Chain.Config.GenesisDelay
		}
	}

	// default validator names
	if cfg.Frontend.ValidatorNamesYaml == "" && cfg.Frontend.ValidatorNamesInventory == "" {
		switch cfg.Chain.Name {
		case "sepolia":
			cfg.Frontend.ValidatorNamesYaml = "~internal/sepolia.names.yml"
		case "holesky":
			cfg.Frontend.ValidatorNamesYaml = "~internal/holesky.names.yml"
		case "creeper2":
			cfg.Frontend.ValidatorNamesYaml = "~internal/creeper2.names.yml"
		}
	}

	// endpoints
	if cfg.BeaconApi.Endpoints == nil && cfg.BeaconApi.Endpoint != "" {
		cfg.BeaconApi.Endpoints = []types.EndpointConfig{
			{
				Url:  cfg.BeaconApi.Endpoint,
				Name: "default",
			},
		}
	}
	for idx, endpoint := range cfg.BeaconApi.Endpoints {
		if endpoint.Name == "" {
			url, _ := url.Parse(endpoint.Url)
			if url != nil {
				cfg.BeaconApi.Endpoints[idx].Name = url.Hostname()
			} else {
				cfg.BeaconApi.Endpoints[idx].Name = fmt.Sprintf("endpoint-%v", idx+1)
			}
		}
	}
	if cfg.BeaconApi.Endpoints == nil || len(cfg.BeaconApi.Endpoints) == 0 {
		return fmt.Errorf("missing beacon node endpoints (need at least 1 endpoint to run the explorer)")
	}

	// execution endpoints
	if cfg.ExecutionApi.Endpoints == nil && cfg.ExecutionApi.Endpoint != "" {
		cfg.ExecutionApi.Endpoints = []types.EndpointConfig{
			{
				Url:  cfg.ExecutionApi.Endpoint,
				Name: "default",
			},
		}
	}
	for idx, endpoint := range cfg.ExecutionApi.Endpoints {
		if endpoint.Name == "" {
			url, _ := url.Parse(endpoint.Url)
			if url != nil {
				cfg.ExecutionApi.Endpoints[idx].Name = url.Hostname()
			} else {
				cfg.ExecutionApi.Endpoints[idx].Name = fmt.Sprintf("endpoint-%v", idx+1)
			}
		}
	}

	// blobstore
	if cfg.BlobStore.NameTemplate == "" {
		cfg.BlobStore.NameTemplate = "{hash}"
	}

	log.WithFields(log.Fields{
		"genesisTimestamp":       cfg.Chain.GenesisTimestamp,
		"configName":             cfg.Chain.Config.ConfigName,
		"depositChainID":         cfg.Chain.Config.DepositChainID,
		"depositNetworkID":       cfg.Chain.Config.DepositNetworkID,
		"depositContractAddress": cfg.Chain.Config.DepositContractAddress,
	}).Infof("did init config")

	return nil
}

func readConfigFile(cfg *types.Config, path string) error {
	if path == "" {
		return yaml.Unmarshal([]byte(config.DefaultConfigYml), cfg)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening config file %v: %v", path, err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("error decoding explorer config: %v", err)
	}
	return nil
}

func readConfigEnv(cfg *types.Config) error {
	return envconfig.Process("", cfg)
}
