#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <version> <path>"
    exit 1
fi

version=$1
path=$2

export PATH="$PATH:$HOME/go/bin"

# Check if the version is one of the allowed options
case $version in
    phase0|altair|bellatrix|capella|deneb|electra|badger)
        # Remove ssz files in spec/$version
        cd spec/$version
        rm -rf *_ssz.go
        cd ../..

        # Run the corresponding version script with the provided path
        sh ./util/sszgen/sszgen_${version}.sh "$path"
        ;;
    *)
        echo "Error: Invalid version specified. Choose from 'altair', 'bellatrix', 'capella', 'deneb', 'electra', or 'badger'."
        exit 1
        ;;
esac