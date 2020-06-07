package bites

import (
	"math"
	"strconv"
)

// BitSize represents a data size measurement in bits.
type BitSize uint64

const SizeBit BitSize = 1

// Base10 bit sizes.
const (
	SizeKilobit = SizeBit * 1000
	SizeMegabit = SizeKilobit * 1000
	SizeGigabit = SizeMegabit * 1000
	SizeTerabit = SizeGigabit * 1000
	SizePetabit = SizeTerabit * 1000
	SizeExabit  = SizePetabit * 1000
)

// Base2 bit sizes.
const (
	SizeKibibit = SizeBit << (10 * (iota + 1))
	SizeMebibit
	SizeGibibit
	SizeTebibit
	SizePebibit
	SizeExbibit
)

// ToByteSize converts the current size in bits to a size in bytes.
func (b BitSize) ToByteSize() ByteSize {
	return ByteSize(math.Ceil(float64(b) / float64(8)))
}

// String renders the current BitSize in base2 format.
//
// This is just an alias for Base2String.
func (b BitSize) String() string {
	return b.Base2String()
}

// Base10String renders the current BitSize in base10 format.
func (b BitSize) Base10String() string {
	var suf string
	var div BitSize

	if b < SizeKilobit {
		suf = AbbrBit
		div = SizeBit
	} else if b < SizeMegabit {
		suf = AbbrKilobit
		div = SizeKilobit
	} else if b < SizeGigabit {
		suf = AbbrMegabit
		div = SizeMegabit
	} else if b < SizeTerabit {
		suf = AbbrGigabit
		div = SizeGigabit
	} else if b < SizePetabit {
		suf = AbbrTerabit
		div = SizeTerabit
	} else if b < SizeExabit {
		suf = AbbrPetabit
		div = SizePetabit
	} else {
		suf = AbbrExabit
		div = SizeExabit
	}

	return strconv.FormatFloat(float64(b)/float64(div), 'G', 3, 64) + suf
}

// Base2String renders the current BitSize in base2 format.
func (b BitSize) Base2String() string {
	var suf string
	var div BitSize

	if b < SizeKibibit {
		suf = AbbrBit
		div = SizeBit
	} else if b < SizeMebibit {
		suf = AbbrKibibit
		div = SizeKibibit
	} else if b < SizeGibibit {
		suf = AbbrMebibit
		div = SizeMebibit
	} else if b < SizeTebibit {
		suf = AbbrGibibit
		div = SizeGibibit
	} else if b < SizePebibit {
		suf = AbbrTebibit
		div = SizeTebibit
	} else if b < SizeExbibit {
		suf = AbbrPebibit
		div = SizePebibit
	} else {
		suf = AbbrExbibit
		div = SizeExbibit
	}

	return strconv.FormatFloat(float64(b)/float64(div), 'G', 3, 64) + suf
}
