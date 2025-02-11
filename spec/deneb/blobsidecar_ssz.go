// Code generated by fastssz. DO NOT EDIT.
// Hash: ace578ab0fc89807a23f06c73a2f166b813eb59bb3cd6fc5fce03a349b97c8fa
// Version: 0.1.4
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlobSidecar object
func (b *BlobSidecar) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobSidecar object to a target array
func (b *BlobSidecar) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, uint64(b.Index))

	// Field (1) 'Blob'
	dst = append(dst, b.Blob[:]...)

	// Field (2) 'KZGCommitment'
	dst = append(dst, b.KZGCommitment[:]...)

	// Field (3) 'KZGProof'
	dst = append(dst, b.KZGProof[:]...)

	// Field (4) 'SignedBlockHeader'
	if b.SignedBlockHeader == nil {
		b.SignedBlockHeader = new(phase0.SignedBeaconBlockHeader)
	}
	if dst, err = b.SignedBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'KZGCommitmentInclusionProof'
	for ii := 0; ii < 544; ii++ {
		dst = append(dst, b.KZGCommitmentInclusionProof[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobSidecar object
func (b *BlobSidecar) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 148792 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	b.Index = BlobIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'Blob'
	copy(b.Blob[:], buf[8:131080])

	// Field (2) 'KZGCommitment'
	copy(b.KZGCommitment[:], buf[131080:131128])

	// Field (3) 'KZGProof'
	copy(b.KZGProof[:], buf[131128:131176])

	// Field (4) 'SignedBlockHeader'
	if b.SignedBlockHeader == nil {
		b.SignedBlockHeader = new(phase0.SignedBeaconBlockHeader)
	}
	if err = b.SignedBlockHeader.UnmarshalSSZ(buf[131176:131384]); err != nil {
		return err
	}

	// Field (5) 'KZGCommitmentInclusionProof'

	for ii := 0; ii < 544; ii++ {
		copy(b.KZGCommitmentInclusionProof[ii][:], buf[131384:148792][ii*32:(ii+1)*32])
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobSidecar object
func (b *BlobSidecar) SizeSSZ() (size int) {
	size = 148792
	return
}

// HashTreeRoot ssz hashes the BlobSidecar object
func (b *BlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecar object with a hasher
func (b *BlobSidecar) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(uint64(b.Index))

	// Field (1) 'Blob'
	hh.PutBytes(b.Blob[:])

	// Field (2) 'KZGCommitment'
	hh.PutBytes(b.KZGCommitment[:])

	// Field (3) 'KZGProof'
	hh.PutBytes(b.KZGProof[:])

	// Field (4) 'SignedBlockHeader'
	if b.SignedBlockHeader == nil {
		b.SignedBlockHeader = new(phase0.SignedBeaconBlockHeader)
	}
	if err = b.SignedBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'KZGCommitmentInclusionProof'
	{
		subIndx := hh.Index()
		for _, i := range b.KZGCommitmentInclusionProof {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlobSidecar object
func (b *BlobSidecar) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
