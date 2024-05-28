#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0","${path_to_go_eth2_client}/spec/altair" \
  --path "${path_to_go_eth2_client}/spec/bellatrix" \
  --objs BeaconBlock,BeaconBlockBody,BeaconState,ExecutionPayload,ExecutionPaylodHeader,SignedBeaconBlock

# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/bellatrix

# Run goimports
goimports -w beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go