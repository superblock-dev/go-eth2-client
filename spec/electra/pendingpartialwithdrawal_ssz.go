// Code generated by fastssz. DO NOT EDIT.
// Hash: 146102f414305035f657eea197646e6993d5204cac10a6a321a0810201479c28
// Version: 0.1.4
package electra

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the PendingPartialWithdrawal object
func (p *PendingPartialWithdrawal) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PendingPartialWithdrawal object to a target array
func (p *PendingPartialWithdrawal) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, uint64(p.Index))

	// Field (1) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(p.Amount))

	// Field (2) 'WithdrawableEpoch'
	dst = ssz.MarshalUint64(dst, uint64(p.WithdrawableEpoch))

	return
}

// UnmarshalSSZ ssz unmarshals the PendingPartialWithdrawal object
func (p *PendingPartialWithdrawal) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	p.Index = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'Amount'
	p.Amount = phase0.Gwei(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'WithdrawableEpoch'
	p.WithdrawableEpoch = phase0.Epoch(ssz.UnmarshallUint64(buf[16:24]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PendingPartialWithdrawal object
func (p *PendingPartialWithdrawal) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the PendingPartialWithdrawal object
func (p *PendingPartialWithdrawal) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PendingPartialWithdrawal object with a hasher
func (p *PendingPartialWithdrawal) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(uint64(p.Index))

	// Field (1) 'Amount'
	hh.PutUint64(uint64(p.Amount))

	// Field (2) 'WithdrawableEpoch'
	hh.PutUint64(uint64(p.WithdrawableEpoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the PendingPartialWithdrawal object
func (p *PendingPartialWithdrawal) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}
