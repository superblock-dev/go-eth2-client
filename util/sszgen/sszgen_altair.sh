#!#!/bin/bash

path_to_go_eth2_client=$(pwd)
path_to_fastssz="/Users/syjn99/projects/personal-project/fastssz"

# Remove ssz files in spec/altair
cd spec/altair
rm -rf *_ssz.go
cd ../..

# Navigate to the fastssz repository path
cd $path_to_fastssz

# Run the sszgen command
go run sszgen/*.go -suffix ssz \
  -include "${path_to_go_eth2_client}/spec/phase0" \
  --path "${path_to_go_eth2_client}/spec/altair" \
  --objs BeaconBlock,BeaconBlockBody,BeaconState,ContributionAndProof,SignedBeaconBlock,SignedContributionAndProof,SyncAggregate,SyncAggregatorSelectionData,SyncCommittee,SyncCommitteeContribution,SyncCommitteeMessage,BailOut

# Navigate to go_eth2_client repo
cd $path_to_go_eth2_client
cd spec/altair

# Run goimports
goimports -w beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go contributionandproof_ssz.go signedbeaconblock_ssz.go signedcontributionandproof_ssz.go syncaggregate_ssz.go syncaggregatorselectiondata_ssz.go synccommitteecontribution_ssz.go synccommitteemessage_ssz.go bailout_ssz.go