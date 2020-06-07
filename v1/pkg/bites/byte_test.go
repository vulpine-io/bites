package bites_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/vulpine-io/bites/v1/pkg/bites"
)

func TestByteSize_String(t *testing.T) {
	Convey("ByteSize.String", t, func() {
		So(bites.SizeByte.String(), ShouldEqual, bites.SizeByte.Base2String())
	})
}

func TestByteSize_Base2String(t *testing.T) {
	Convey("ByteSize.Base2String", t, func() {
		tests := []struct {
			in  bites.ByteSize
			out string
		}{
			{24 * bites.SizeByte, "24B"},
			{13 * bites.SizeKibibyte, "13KiB"},
			{3 * bites.SizeMebibyte, "3MiB"},
			{89 * bites.SizeGibibyte, "89GiB"},
			{101 * bites.SizeTebibyte, "101TiB"},
			{76 * bites.SizePebibyte, "76PiB"},
			{4 * bites.SizeExbibyte, "4EiB"},
		}

		for _, test := range tests {
			So(test.in.Base2String(), ShouldEqual, test.out)
		}
	})
}

func TestByteSize_Base10String(t *testing.T) {
	Convey("ByteSize.Base10String", t, func() {
		tests := []struct {
			in  bites.ByteSize
			out string
		}{
			{24 * bites.SizeByte, "24B"},
			{13 * bites.SizeKilobyte, "13KB"},
			{3 * bites.SizeMegabyte, "3MB"},
			{89 * bites.SizeGigabyte, "89GB"},
			{101 * bites.SizeTerabyte, "101TB"},
			{76 * bites.SizePetabyte, "76PB"},
			{4 * bites.SizeExabyte, "4EB"},
		}

		for _, test := range tests {
			So(test.in.Base10String(), ShouldEqual, test.out)
		}
	})
}

func TestByteSize_ToBitSize(t *testing.T) {
	Convey("ByteSize.ToBitSize", t, func() {
		So(bites.SizeTerabyte.ToBitSize(), ShouldResemble, 8*bites.SizeTerabit)
	})
}
