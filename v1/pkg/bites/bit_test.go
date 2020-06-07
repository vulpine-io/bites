package bites_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/vulpine-io/bites/v1/pkg/bites"
)

func TestBitSize_String(t *testing.T) {
	Convey("BitSize.String", t, func() {
		So(bites.SizeBit.String(), ShouldEqual, bites.SizeBit.Base2String())
	})
}

func TestBitSize_Base2String(t *testing.T) {
	Convey("BitSize.Base2String", t, func() {
		tests := []struct {
			in  bites.BitSize
			out string
		}{
			{24 * bites.SizeBit, "24b"},
			{13 * bites.SizeKibibit, "13Kib"},
			{3 * bites.SizeMebibit, "3Mib"},
			{89 * bites.SizeGibibit, "89Gib"},
			{101 * bites.SizeTebibit, "101Tib"},
			{76 * bites.SizePebibit, "76Pib"},
			{4 * bites.SizeExbibit, "4Eib"},
		}

		for _, test := range tests {
			So(test.in.Base2String(), ShouldEqual, test.out)
		}
	})
}

func TestBitSize_Base10String(t *testing.T) {
	Convey("BitSize.Base10String", t, func() {
		tests := []struct {
			in  bites.BitSize
			out string
		}{
			{24 * bites.SizeBit, "24b"},
			{13 * bites.SizeKilobit, "13Kb"},
			{3 * bites.SizeMegabit, "3Mb"},
			{89 * bites.SizeGigabit, "89Gb"},
			{101 * bites.SizeTerabit, "101Tb"},
			{76 * bites.SizePetabit, "76Pb"},
			{4 * bites.SizeExabit, "4Eb"},
		}

		for _, test := range tests {
			So(test.in.Base10String(), ShouldEqual, test.out)
		}
	})
}

func TestBitSize_ToByteSize(t *testing.T) {
	Convey("BitSize.ToByteSize", t, func() {
		So(8*bites.SizeTerabit.ToByteSize(), ShouldResemble, bites.SizeTerabyte)
	})
}
