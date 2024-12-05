#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0","${path_to_go_eth2_client}/spec/altair","${path_to_go_eth2_client}/spec/bellatrix" \
  --path "${path_to_go_eth2_client}/spec/capella" \
  --objs BeaconBlockBody,BeaconBlock,BeaconState,ExecutionPayload,ExecutionPayloadHeader,HistoricalSummary,SignedBeaconBlock,Withdrawal

# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/capella

# Run goimports
goimports -w beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go executionpayloadheader_ssz.go executionpayload_ssz.go historicalsummary_ssz.go signedbeaconblock_ssz.go withdrawal_ssz.go