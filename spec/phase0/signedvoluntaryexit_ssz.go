// Code generated by fastssz. DO NOT EDIT.
// Hash: 9d533094809c8c7d7e0d2dde661f10704af2045ea959402d40d49f3c41c90bb1
// Version: 0.1.4
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedVoluntaryExit object to a target array
func (s *SignedVoluntaryExit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(VoluntaryExit)
	}
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(VoluntaryExit)
	}
	if err = s.Message.UnmarshalSSZ(buf[0:16]); err != nil {
		return err
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[16:112])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedVoluntaryExit object with a hasher
func (s *SignedVoluntaryExit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(VoluntaryExit)
	}
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
