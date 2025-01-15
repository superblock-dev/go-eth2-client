#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0","${path_to_go_eth2_client}/spec/altair","${path_to_go_eth2_client}/spec/bellatrix","${path_to_go_eth2_client}/spec/capella","${path_to_go_eth2_client}/spec/deneb","${path_to_go_eth2_client}/spec/electra" \
  --path "${path_to_go_eth2_client}/spec/badger" \
  --objs BeaconBlockBody,BeaconBlock,BeaconState,ExecutionPayload,ExecutionPayloadHeader,SignedBeaconBlock
# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/badger

# Run goimports
goimports -w beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go