// Copyright © 2024 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package badger

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// beaconBlockBodyYAML is the spec representation of the struct.
type beaconBlockBodyYAML struct {
	RANDAOReveal       string                        `yaml:"randao_reveal"`
	ETH1Data           *phase0.ETH1Data              `yaml:"eth1_data"`
	Graffiti           string                        `yaml:"graffiti"`
	ProposerSlashings  []*phase0.ProposerSlashing    `yaml:"proposer_slashings"`
	AttesterSlashings  []*electra.AttesterSlashing           `yaml:"attester_slashings"`
	Attestations       []*electra.Attestation                `yaml:"attestations"`
	Deposits           []*phase0.Deposit             `yaml:"deposits"`
	VoluntaryExits     []*phase0.SignedVoluntaryExit `yaml:"voluntary_exits"`
	ExecutionPayload   *deneb.ExecutionPayload       `yaml:"execution_payload"`
	BlobKZGCommitments []string                      `yaml:"blob_kzg_commitments"`
	ExecutionRequests  *electra.ExecutionRequests            `yaml:"execution_requests"`
}

// MarshalYAML implements yaml.Marshaler.
func (b *BeaconBlockBody) MarshalYAML() ([]byte, error) {
	blobKZGCommitments := make([]string, len(b.BlobKZGCommitments))
	for i := range b.BlobKZGCommitments {
		blobKZGCommitments[i] = b.BlobKZGCommitments[i].String()
	}

	yamlBytes, err := yaml.MarshalWithOptions(&beaconBlockBodyYAML{
		RANDAOReveal:       b.RANDAOReveal.String(),
		ETH1Data:           b.ETH1Data,
		Graffiti:           fmt.Sprintf("%#x", b.Graffiti),
		ProposerSlashings:  b.ProposerSlashings,
		AttesterSlashings:  b.AttesterSlashings,
		Attestations:       b.Attestations,
		Deposits:           b.Deposits,
		VoluntaryExits:     b.VoluntaryExits,
		ExecutionPayload:   b.ExecutionPayload,
		BlobKZGCommitments: blobKZGCommitments,
		ExecutionRequests:  b.ExecutionRequests,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}

	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (b *BeaconBlockBody) UnmarshalYAML(input []byte) error {
	// This is very inefficient, but YAML is only used for spec tests so we do this
	// rather than maintain a custom YAML unmarshaller.
	var unmarshaled beaconBlockBodyJSON
	if err := yaml.Unmarshal(input, &unmarshaled); err != nil {
		return errors.Wrap(err, "failed to unmarshal YAML")
	}
	marshaled, err := json.Marshal(unmarshaled)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}

	return b.UnmarshalJSON(marshaled)
}
