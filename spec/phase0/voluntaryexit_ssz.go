// Code generated by fastssz. DO NOT EDIT.
// Hash: 9d533094809c8c7d7e0d2dde661f10704af2045ea959402d40d49f3c41c90bb1
// Version: 0.1.4
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the VoluntaryExit object
func (v *VoluntaryExit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the VoluntaryExit object to a target array
func (v *VoluntaryExit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, uint64(v.Epoch))

	// Field (1) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(v.ValidatorIndex))

	return
}

// UnmarshalSSZ ssz unmarshals the VoluntaryExit object
func (v *VoluntaryExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Epoch'
	v.Epoch = Epoch(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ValidatorIndex'
	v.ValidatorIndex = ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the VoluntaryExit object
func (v *VoluntaryExit) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the VoluntaryExit object
func (v *VoluntaryExit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the VoluntaryExit object with a hasher
func (v *VoluntaryExit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Epoch'
	hh.PutUint64(uint64(v.Epoch))

	// Field (1) 'ValidatorIndex'
	hh.PutUint64(uint64(v.ValidatorIndex))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the VoluntaryExit object
func (v *VoluntaryExit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}
