package bites

import (
	"math"
	"strconv"
)

// ByteSize represents a data size measurement in bytes.
type ByteSize uint64

const SizeByte ByteSize = 1

// Base10 byte sizes.
const (
	SizeKilobyte = SizeByte * 1000
	SizeMegabyte = SizeKilobyte * 1000
	SizeGigabyte = SizeMegabyte * 1000
	SizeTerabyte = SizeGigabyte * 1000
	SizePetabyte = SizeTerabyte * 1000
	SizeExabyte  = SizePetabyte * 1000
)

// Base2 byte sizes.
const (
	SizeKibibyte = SizeByte << (10 * (iota + 1))
	SizeMebibyte
	SizeGibibyte
	SizeTebibyte
	SizePebibyte
	SizeExbibyte
)

// ToBitSize converts the current size in bytes to a size in bits.
func (b ByteSize) ToBitSize() BitSize {
	return BitSize(math.Ceil(float64(b) * float64(8)))
}

// String renders the current ByteSize in base2 format.
//
// This is just an alias for Base2String.
func (b ByteSize) String() string {
	return b.Base2String()
}

// Base10String renders the current ByteSize in base10 format.
func (b ByteSize) Base10String() string {
	var suf string
	var div ByteSize

	if b < SizeKilobyte {
		suf = AbbrByte
		div = SizeByte
	} else if b < SizeMegabyte {
		suf = AbbrKilobyte
		div = SizeKilobyte
	} else if b < SizeGigabyte {
		suf = AbbrMegabyte
		div = SizeMegabyte
	} else if b < SizeTerabyte {
		suf = AbbrGigabyte
		div = SizeGigabyte
	} else if b < SizePetabyte {
		suf = AbbrTerabyte
		div = SizeTerabyte
	} else if b < SizeExabyte {
		suf = AbbrPetabyte
		div = SizePetabyte
	} else {
		suf = AbbrExabyte
		div = SizeExabyte
	}

	return strconv.FormatFloat(float64(b)/float64(div), 'G', 3, 64) + suf
}

// Base2String renders the current ByteSize in base2 format.
func (b ByteSize) Base2String() string {
	var suf string
	var div ByteSize

	if b < SizeKibibyte {
		suf = AbbrByte
		div = 1
	} else if b < SizeMebibyte {
		suf = AbbrKibibyte
		div = SizeKibibyte
	} else if b < SizeGibibyte {
		suf = AbbrMebibyte
		div = SizeMebibyte
	} else if b < SizeTebibyte {
		suf = AbbrGibibyte
		div = SizeGibibyte
	} else if b < SizePebibyte {
		suf = AbbrTebibyte
		div = SizeTebibyte
	} else if b < SizeExbibyte {
		suf = AbbrPebibyte
		div = SizePebibyte
	} else {
		suf = AbbrExbibyte
		div = SizeExbibyte
	}

	return strconv.FormatFloat(float64(b)/float64(div), 'G', 3, 64) + suf
}
