package types

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestVectorReplace(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	item1, _ := hex.DecodeString("02")
	item2, _ := hex.DecodeString("03")
	item3, _ := hex.DecodeString("04")
	item4, _ := hex.DecodeString("05")
	xBuild := NewBytesBuilder().Push(*ByteFromSliceUnchecked(item0)).Push(*ByteFromSliceUnchecked(item3)).Push(*ByteFromSliceUnchecked(item4))
	x3 := xBuild.Replace(1, *ByteFromSliceUnchecked(item1))
	x4 := xBuild.Replace(2, *ByteFromSliceUnchecked(item2))
	x := xBuild.Build()
	if bytes.Compare(item3, x3.AsSlice()) != 0 {
		t.Error("type x3 error: ", x3.AsSlice(), item3)
	}
	if bytes.Compare(item4, x4.AsSlice()) != 0 {
		t.Error("type x4 error: ", x4.AsSlice(), item4)
	}

	expected, _ := hex.DecodeString("03000000010203")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", x.AsSlice(), expected)
	}
	y, _ := BytesFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", y.AsSlice(), expected)
	}

	if x.TotalSize() != uint(len(x.AsSlice())) {
		t.Error("struct: Bytes:\n data: ", x.AsSlice(), "\n partial read total_size: ", x.TotalSize(), ", actual: ", len(x.AsSlice()))
	}
}
