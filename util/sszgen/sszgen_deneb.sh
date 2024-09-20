#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0","${path_to_go_eth2_client}/spec/altair","${path_to_go_eth2_client}/spec/bellatrix","${path_to_go_eth2_client}/spec/capella" \
  --path "${path_to_go_eth2_client}/spec/deneb" \
  --objs BeaconBlockBody,BeaconBlock,BeaconState,BlobIdentifier,BlobSidecar,ExecutionPayload,ExecutionPayloadHeader,SignedBeaconBlock,SignedBlobSidecar

# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/deneb

# Run goimports
goimports -w beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go blobidentifier_ssz.go blobsidecar_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go