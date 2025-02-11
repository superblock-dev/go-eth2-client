// Code generated by fastssz. DO NOT EDIT.
// Hash: 9d533094809c8c7d7e0d2dde661f10704af2045ea959402d40d49f3c41c90bb1
// Version: 0.1.4
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(2687401)

	// Field (0) 'GenesisTime'
	dst = ssz.MarshalUint64(dst, b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if dst, err = b.Fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	if size := len(b.BlockRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = append(dst, b.BlockRoots[ii][:]...)
	}

	// Field (6) 'StateRoots'
	if size := len(b.StateRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = append(dst, b.StateRoots[ii][:]...)
	}

	// Offset (7) 'HistoricalRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.HistoricalRoots) * 32

	// Field (8) 'RewardAdjustmentFactor'
	dst = ssz.MarshalUint64(dst, b.RewardAdjustmentFactor)

	// Field (9) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(ETH1Data)
	}
	if dst, err = b.ETH1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (10) 'ETH1DataVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ETH1DataVotes) * 72

	// Field (11) 'ETH1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.ETH1DepositIndex)

	// Offset (12) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (13) 'Balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Balances) * 8

	// Field (14) 'PreviousEpochReserve'
	dst = ssz.MarshalUint64(dst, b.PreviousEpochReserve)

	// Field (15) 'CurrentEpochReserve'
	dst = ssz.MarshalUint64(dst, b.CurrentEpochReserve)

	// Field (16) 'RANDAOMixes'
	if size := len(b.RANDAOMixes); size != 65536 {
		err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
		return
	}
	for ii := 0; ii < 65536; ii++ {
		dst = append(dst, b.RANDAOMixes[ii][:]...)
	}

	// Field (17) 'Slashings'
	if size := len(b.Slashings); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Slashings[ii]))
	}

	// Offset (18) 'PreviousEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		offset += 4
		offset += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Offset (19) 'CurrentEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)

	// Field (20) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	dst = append(dst, b.JustificationBits...)

	// Field (21) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.PreviousJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (22) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.CurrentJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (23) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.FinalizedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'HistoricalRoots'
	if size := len(b.HistoricalRoots); size > 16777216 {
		err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
		return
	}
	for ii := 0; ii < len(b.HistoricalRoots); ii++ {
		dst = append(dst, b.HistoricalRoots[ii][:]...)
	}

	// Field (10) 'ETH1DataVotes'
	if size := len(b.ETH1DataVotes); size > 2048 {
		err = ssz.ErrListTooBigFn("BeaconState.ETH1DataVotes", size, 2048)
		return
	}
	for ii := 0; ii < len(b.ETH1DataVotes); ii++ {
		if dst, err = b.ETH1DataVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (12) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (13) 'Balances'
	if size := len(b.Balances); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Balances); ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Balances[ii]))
	}

	// Field (18) 'PreviousEpochAttestations'
	if size := len(b.PreviousEpochAttestations); size > 4096 {
		err = ssz.ErrListTooBigFn("BeaconState.PreviousEpochAttestations", size, 4096)
		return
	}
	{
		offset = 4 * len(b.PreviousEpochAttestations)
		for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.PreviousEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		if dst, err = b.PreviousEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (19) 'CurrentEpochAttestations'
	if size := len(b.CurrentEpochAttestations); size > 4096 {
		err = ssz.ErrListTooBigFn("BeaconState.CurrentEpochAttestations", size, 4096)
		return
	}
	{
		offset = 4 * len(b.CurrentEpochAttestations)
		for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.CurrentEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		if dst, err = b.CurrentEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 2687401 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o10, o12, o13, o18, o19 uint64

	// Field (0) 'GenesisTime'
	b.GenesisTime = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[8:40])

	// Field (2) 'Slot'
	b.Slot = Slot(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if err = b.Fork.UnmarshalSSZ(buf[48:64]); err != nil {
		return err
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[64:176]); err != nil {
		return err
	}

	// Field (5) 'BlockRoots'
	b.BlockRoots = make([]Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.BlockRoots[ii][:], buf[176:262320][ii*32:(ii+1)*32])
	}

	// Field (6) 'StateRoots'
	b.StateRoots = make([]Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.StateRoots[ii][:], buf[262320:524464][ii*32:(ii+1)*32])
	}

	// Offset (7) 'HistoricalRoots'
	if o7 = ssz.ReadOffset(buf[524464:524468]); o7 > size {
		return ssz.ErrOffset
	}

	if o7 != 2687401 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (8) 'RewardAdjustmentFactor'
	b.RewardAdjustmentFactor = ssz.UnmarshallUint64(buf[524468:524476])

	// Field (9) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(ETH1Data)
	}
	if err = b.ETH1Data.UnmarshalSSZ(buf[524476:524548]); err != nil {
		return err
	}

	// Offset (10) 'ETH1DataVotes'
	if o10 = ssz.ReadOffset(buf[524548:524552]); o10 > size || o7 > o10 {
		return ssz.ErrOffset
	}

	// Field (11) 'ETH1DepositIndex'
	b.ETH1DepositIndex = ssz.UnmarshallUint64(buf[524552:524560])

	// Offset (12) 'Validators'
	if o12 = ssz.ReadOffset(buf[524560:524564]); o12 > size || o10 > o12 {
		return ssz.ErrOffset
	}

	// Offset (13) 'Balances'
	if o13 = ssz.ReadOffset(buf[524564:524568]); o13 > size || o12 > o13 {
		return ssz.ErrOffset
	}

	// Field (14) 'PreviousEpochReserve'
	b.PreviousEpochReserve = ssz.UnmarshallUint64(buf[524568:524576])

	// Field (15) 'CurrentEpochReserve'
	b.CurrentEpochReserve = ssz.UnmarshallUint64(buf[524576:524584])

	// Field (16) 'RANDAOMixes'
	b.RANDAOMixes = make([]Root, 65536)
	for ii := 0; ii < 65536; ii++ {
		copy(b.RANDAOMixes[ii][:], buf[524584:2621736][ii*32:(ii+1)*32])
	}

	// Field (17) 'Slashings'
	b.Slashings = make([]Gwei, 8192)
	for ii := 0; ii < 8192; ii++ {
		b.Slashings[ii] = Gwei(ssz.UnmarshallUint64(buf[2621736:2687272][ii*8 : (ii+1)*8]))
	}

	// Offset (18) 'PreviousEpochAttestations'
	if o18 = ssz.ReadOffset(buf[2687272:2687276]); o18 > size || o13 > o18 {
		return ssz.ErrOffset
	}

	// Offset (19) 'CurrentEpochAttestations'
	if o19 = ssz.ReadOffset(buf[2687276:2687280]); o19 > size || o18 > o19 {
		return ssz.ErrOffset
	}

	// Field (20) 'JustificationBits'
	if cap(b.JustificationBits) == 0 {
		b.JustificationBits = make([]byte, 0, len(buf[2687280:2687281]))
	}
	b.JustificationBits = append(b.JustificationBits, buf[2687280:2687281]...)

	// Field (21) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.UnmarshalSSZ(buf[2687281:2687321]); err != nil {
		return err
	}

	// Field (22) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.UnmarshalSSZ(buf[2687321:2687361]); err != nil {
		return err
	}

	// Field (23) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(Checkpoint)
	}
	if err = b.FinalizedCheckpoint.UnmarshalSSZ(buf[2687361:2687401]); err != nil {
		return err
	}

	// Field (7) 'HistoricalRoots'
	{
		buf = tail[o7:o10]
		num, err := ssz.DivideInt2(len(buf), 32, 16777216)
		if err != nil {
			return err
		}
		b.HistoricalRoots = make([]Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.HistoricalRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (10) 'ETH1DataVotes'
	{
		buf = tail[o10:o12]
		num, err := ssz.DivideInt2(len(buf), 72, 2048)
		if err != nil {
			return err
		}
		b.ETH1DataVotes = make([]*ETH1Data, num)
		for ii := 0; ii < num; ii++ {
			if b.ETH1DataVotes[ii] == nil {
				b.ETH1DataVotes[ii] = new(ETH1Data)
			}
			if err = b.ETH1DataVotes[ii].UnmarshalSSZ(buf[ii*72 : (ii+1)*72]); err != nil {
				return err
			}
		}
	}

	// Field (12) 'Validators'
	{
		buf = tail[o12:o13]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (13) 'Balances'
	{
		buf = tail[o13:o18]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Balances = make([]Gwei, num)
		for ii := 0; ii < num; ii++ {
			b.Balances[ii] = Gwei(ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8]))
		}
	}

	// Field (18) 'PreviousEpochAttestations'
	{
		buf = tail[o18:o19]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.PreviousEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.PreviousEpochAttestations[indx] == nil {
				b.PreviousEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.PreviousEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (19) 'CurrentEpochAttestations'
	{
		buf = tail[o19:]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.CurrentEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.CurrentEpochAttestations[indx] == nil {
				b.CurrentEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.CurrentEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 2687401

	// Field (7) 'HistoricalRoots'
	size += len(b.HistoricalRoots) * 32

	// Field (10) 'ETH1DataVotes'
	size += len(b.ETH1DataVotes) * 72

	// Field (12) 'Validators'
	size += len(b.Validators) * 121

	// Field (13) 'Balances'
	size += len(b.Balances) * 8

	// Field (18) 'PreviousEpochAttestations'
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		size += 4
		size += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Field (19) 'CurrentEpochAttestations'
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		size += 4
		size += b.CurrentEpochAttestations[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the BeaconState object
func (b *BeaconState) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher
func (b *BeaconState) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisTime'
	hh.PutUint64(b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (2) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if err = b.Fork.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	{
		if size := len(b.BlockRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (6) 'StateRoots'
	{
		if size := len(b.StateRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (7) 'HistoricalRoots'
	{
		if size := len(b.HistoricalRoots); size > 16777216 {
			err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.HistoricalRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.HistoricalRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 16777216)
	}

	// Field (8) 'RewardAdjustmentFactor'
	hh.PutUint64(b.RewardAdjustmentFactor)

	// Field (9) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(ETH1Data)
	}
	if err = b.ETH1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (10) 'ETH1DataVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.ETH1DataVotes))
		if num > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.ETH1DataVotes {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2048)
	}

	// Field (11) 'ETH1DepositIndex'
	hh.PutUint64(b.ETH1DepositIndex)

	// Field (12) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (13) 'Balances'
	{
		if size := len(b.Balances); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Balances {
			hh.AppendUint64(uint64(i))
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Balances))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (14) 'PreviousEpochReserve'
	hh.PutUint64(b.PreviousEpochReserve)

	// Field (15) 'CurrentEpochReserve'
	hh.PutUint64(b.CurrentEpochReserve)

	// Field (16) 'RANDAOMixes'
	{
		if size := len(b.RANDAOMixes); size != 65536 {
			err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RANDAOMixes {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (17) 'Slashings'
	{
		if size := len(b.Slashings); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Slashings {
			hh.AppendUint64(uint64(i))
		}
		hh.Merkleize(subIndx)
	}

	// Field (18) 'PreviousEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.PreviousEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.PreviousEpochAttestations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (19) 'CurrentEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.CurrentEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.CurrentEpochAttestations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (20) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	hh.PutBytes(b.JustificationBits)

	// Field (21) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (22) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (23) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(Checkpoint)
	}
	if err = b.FinalizedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconState object
func (b *BeaconState) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
