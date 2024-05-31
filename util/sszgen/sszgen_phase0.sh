#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="${1}"

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  --path "${path_to_go_eth2_client}/spec/phase0" \
  --objs AggregateAndProof,AttestationData,Attestation,AttesterSlashing,BeaconBlockBody,BeaconBlock,BeaconBlockHeader,BeaconState,Checkpoint,Deposit,DepositData,DepositMessage,ETH1Data,Fork,ForkData,IndexedAttestation,PendingAttestation,ProposerSlashing,SignedAggregateAndProof,SignedBeaconBlock,SignedBeaconBlockHeader,SignedVoluntaryExit,SigningData,Validator,VoluntaryExit

# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/phase0

# Run goimports
goimports -w aggregateandproof_ssz.go attestationdata_ssz.go attestation_ssz.go attesterslashing_ssz.go beaconblockbody_ssz.go beaconblock_ssz.go beaconblockheader_ssz.go beaconstate_ssz.go checkpoint_ssz.go depositdata_ssz.go deposit_ssz.go depositmessage_ssz.go eth1data_ssz.go forkdata_ssz.go fork_ssz.go indexedattestation_ssz.go pendingattestation_ssz.go proposerslashing_ssz.go signedaggregateandproof_ssz.go signedbeaconblock_ssz.go signedbeaconblockheader_ssz.go signedvoluntaryexit_ssz.go signingdata_ssz.go validator_ssz.go voluntaryexit_ssz.go