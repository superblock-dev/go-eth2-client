#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0","${path_to_go_eth2_client}/spec/altair","${path_to_go_eth2_client}/spec/bellatrix","${path_to_go_eth2_client}/spec/capella","${path_to_go_eth2_client}/spec/deneb" \
  --path "${path_to_go_eth2_client}/spec/electra" \
  --objs AggregateAndProof,Attestation,AttesterSlashing,BeaconBlockBody,BeaconBlock,BeaconState,Consolidation,DepositReceipt,ExecutionLayerWithdrawalRequest,ExecutionPayload,ExecutionPayloadHeader,PendingBalanceDeposit,PendingConsolidation,PendingPartialWithdrawal,SignedAggregateAndProof,SignedBeaconBlock,SignedConsolidation
# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/electra

# Run goimports
goimports -w aggregateandproof_ssz.go attestation_ssz.go attesterslashing_ssz.go beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go consolidation_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go pendingbalancedeposit_ssz.go pendingconsolidation_ssz.go pendingpartialwithdrawal_ssz.go signedaggregateandproof_ssz.go signedbeaconblock_ssz.go