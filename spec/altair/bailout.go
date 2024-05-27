package altair

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

type BailOut struct {
	ValidatorIndex phase0.ValidatorIndex
}

type bailOutJSON struct {
	ValidatorIndex string `json:"validator_index"`
}

type bailOutYAML struct {
	ValidatorIndex uint64 `yaml:"validator_index"`
}

// MarshalJSON implements json.Marshaler.
func (s *BailOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(&bailOutJSON{
		ValidatorIndex: fmt.Sprintf("%d", s.ValidatorIndex),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *BailOut) UnmarshalJSON(input []byte) error {
	var bailOutJSON bailOutJSON
	if err := json.Unmarshal(input, &bailOutJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return s.unpack(&bailOutJSON)
}

func (s *BailOut) unpack(bailOutJSON *bailOutJSON) error {
	if bailOutJSON.ValidatorIndex == "" {
		return errors.New("validator index missing")
	}
	validatorIndex, err := strconv.ParseUint(bailOutJSON.ValidatorIndex, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for validator index")
	}
	s.ValidatorIndex = phase0.ValidatorIndex(validatorIndex)

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (s *BailOut) MarshalYAML() ([]byte, error) {
	yamlBytes, err := yaml.MarshalWithOptions(&bailOutYAML{
		ValidatorIndex: uint64(s.ValidatorIndex),
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}

	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (s *BailOut) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var bailOutJSON bailOutJSON
	if err := yaml.Unmarshal(input, &bailOutJSON); err != nil {
		return err
	}

	return s.unpack(&bailOutJSON)
}

// String returns a string version of the structure.
func (s *BailOut) String() string {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
