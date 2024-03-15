// Code generated by fastssz. DO NOT EDIT.
// Hash: 8d95c04360cb520b903fd948d060f5d35ca079f12b8e96804a39bf9ce1ef9c5d
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBlockContents object
func (s *SignedBlockContents) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBlockContents object to a target array
func (s *SignedBlockContents) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'SignedBlock'
	dst = ssz.WriteOffset(dst, offset)
	if s.SignedBlock == nil {
		s.SignedBlock = new(deneb.SignedBeaconBlock)
	}
	offset += s.SignedBlock.SizeSSZ()

	// Offset (1) 'KZGProofs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.KZGProofs) * 48

	// Offset (2) 'Blobs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Blobs) * 131072

	// Field (0) 'SignedBlock'
	if dst, err = s.SignedBlock.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'KZGProofs'
	if size := len(s.KZGProofs); size > 4096 {
		err = ssz.ErrListTooBigFn("SignedBlockContents.KZGProofs", size, 4096)
		return
	}
	for ii := 0; ii < len(s.KZGProofs); ii++ {
		dst = append(dst, s.KZGProofs[ii][:]...)
	}

	// Field (2) 'Blobs'
	if size := len(s.Blobs); size > 4096 {
		err = ssz.ErrListTooBigFn("SignedBlockContents.Blobs", size, 4096)
		return
	}
	for ii := 0; ii < len(s.Blobs); ii++ {
		dst = append(dst, s.Blobs[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBlockContents object
func (s *SignedBlockContents) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'SignedBlock'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'KZGProofs'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'Blobs'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'SignedBlock'
	{
		buf = tail[o0:o1]
		if s.SignedBlock == nil {
			s.SignedBlock = new(deneb.SignedBeaconBlock)
		}
		if err = s.SignedBlock.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'KZGProofs'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 48, 4096)
		if err != nil {
			return err
		}
		s.KZGProofs = make([]deneb.KZGProof, num)
		for ii := 0; ii < num; ii++ {
			copy(s.KZGProofs[ii][:], buf[ii*48:(ii+1)*48])
		}
	}

	// Field (2) 'Blobs'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 131072, 4096)
		if err != nil {
			return err
		}
		s.Blobs = make([]deneb.Blob, num)
		for ii := 0; ii < num; ii++ {
			copy(s.Blobs[ii][:], buf[ii*131072:(ii+1)*131072])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBlockContents object
func (s *SignedBlockContents) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'SignedBlock'
	if s.SignedBlock == nil {
		s.SignedBlock = new(deneb.SignedBeaconBlock)
	}
	size += s.SignedBlock.SizeSSZ()

	// Field (1) 'KZGProofs'
	size += len(s.KZGProofs) * 48

	// Field (2) 'Blobs'
	size += len(s.Blobs) * 131072

	return
}

// HashTreeRoot ssz hashes the SignedBlockContents object
func (s *SignedBlockContents) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBlockContents object with a hasher
func (s *SignedBlockContents) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SignedBlock'
	if err = s.SignedBlock.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'KZGProofs'
	{
		if size := len(s.KZGProofs); size > 4096 {
			err = ssz.ErrListTooBigFn("SignedBlockContents.KZGProofs", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.KZGProofs {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(s.KZGProofs))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
	}

	// Field (2) 'Blobs'
	{
		if size := len(s.Blobs); size > 4096 {
			err = ssz.ErrListTooBigFn("SignedBlockContents.Blobs", size, 4096)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Blobs {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(s.Blobs))
		hh.MerkleizeWithMixin(subIndx, numItems, 4096)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedBlockContents object
func (s *SignedBlockContents) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}