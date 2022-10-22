// Code generated by fastssz. DO NOT EDIT.
// Hash: 1d60b2f7503ae62dbe1b1eec7aef63ac15e11b4caf0d99046583503ea742c88a
package capella

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BLSToExecutionChange object
func (b *BLSToExecutionChange) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BLSToExecutionChange object to a target array
func (b *BLSToExecutionChange) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ValidatorIndex))

	// Field (1) 'FromBLSPubkey'
	dst = append(dst, b.FromBLSPubkey[:]...)

	// Field (2) 'ToExecutionAddress'
	dst = append(dst, b.ToExecutionAddress[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BLSToExecutionChange object
func (b *BLSToExecutionChange) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 76 {
		return ssz.ErrSize
	}

	// Field (0) 'ValidatorIndex'
	b.ValidatorIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'FromBLSPubkey'
	copy(b.FromBLSPubkey[:], buf[8:56])

	// Field (2) 'ToExecutionAddress'
	copy(b.ToExecutionAddress[:], buf[56:76])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BLSToExecutionChange object
func (b *BLSToExecutionChange) SizeSSZ() (size int) {
	size = 76
	return
}

// HashTreeRoot ssz hashes the BLSToExecutionChange object
func (b *BLSToExecutionChange) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BLSToExecutionChange object with a hasher
func (b *BLSToExecutionChange) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ValidatorIndex'
	hh.PutUint64(uint64(b.ValidatorIndex))

	// Field (1) 'FromBLSPubkey'
	hh.PutBytes(b.FromBLSPubkey[:])

	// Field (2) 'ToExecutionAddress'
	hh.PutBytes(b.ToExecutionAddress[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BLSToExecutionChange object
func (b *BLSToExecutionChange) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}