package types

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestByte7Simple1(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	x := NewByte7Builder().Nth0(*ByteFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("01000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", y.AsSlice(), expected)
	}
}

func TestByte7Simple2(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	item3, _ := hex.DecodeString("03")
	x := NewByte7Builder().Nth0(*ByteFromSliceUnchecked(item0)).Nth3(*ByteFromSliceUnchecked(item3)).Build()

	expected, _ := hex.DecodeString("01000003000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", y.AsSlice(), expected)
	}
}

func TestByte7Simple3(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	item3, _ := hex.DecodeString("03")
	item6, _ := hex.DecodeString("06")
	x := NewByte7Builder().Nth0(*ByteFromSliceUnchecked(item0)).Nth3(*ByteFromSliceUnchecked(item3)).Nth6(*ByteFromSliceUnchecked(item6)).Build()

	expected, _ := hex.DecodeString("01000003000006")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7 error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple4(t *testing.T) {
	f1, _ := hex.DecodeString("001122")
	x := NewStructHBuilder().F1(*Byte3FromSliceUnchecked(f1)).Build()

	expected, _ := hex.DecodeString("00112200000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple5(t *testing.T) {
	f2, _ := hex.DecodeString("33")
	x := NewStructHBuilder().F2(*ByteFromSliceUnchecked(f2)).Build()

	expected, _ := hex.DecodeString("00000033000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple6(t *testing.T) {
	f3, _ := hex.DecodeString("4455")
	x := NewStructHBuilder().F3(*Byte2FromSliceUnchecked(f3)).Build()

	expected, _ := hex.DecodeString("00000000445500000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple7(t *testing.T) {
	f4, _ := hex.DecodeString("66778899")
	x := NewStructHBuilder().F4(*Byte4FromSliceUnchecked(f4)).Build()

	expected, _ := hex.DecodeString("00000000000066778899")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple8(t *testing.T) {
	f1, _ := hex.DecodeString("001122")
	f2, _ := hex.DecodeString("33")
	x := NewStructHBuilder().F1(*Byte3FromSliceUnchecked(f1)).F2(*ByteFromSliceUnchecked(f2)).Build()

	expected, _ := hex.DecodeString("00112233000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple9(t *testing.T) {
	f1, _ := hex.DecodeString("001122")
	f2, _ := hex.DecodeString("33")
	f3, _ := hex.DecodeString("4455")
	x := NewStructHBuilder().F1(*Byte3FromSliceUnchecked(f1)).F2(*ByteFromSliceUnchecked(f2)).F3(*Byte2FromSliceUnchecked(f3)).Build()

	expected, _ := hex.DecodeString("00112233445500000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestStructHSimple10(t *testing.T) {
	f1, _ := hex.DecodeString("001122")
	f2, _ := hex.DecodeString("33")
	f3, _ := hex.DecodeString("4455")
	f4, _ := hex.DecodeString("66778899")
	x := NewStructHBuilder().F1(*Byte3FromSliceUnchecked(f1)).F2(*ByteFromSliceUnchecked(f2)).F3(*Byte2FromSliceUnchecked(f3)).F4(*Byte4FromSliceUnchecked(f4)).Build()

	expected, _ := hex.DecodeString("00112233445566778899")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH error: ", y.AsSlice(), expected)
	}
}

func TestBytesSimple11(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	x := NewBytesBuilder().Push(*ByteFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0100000001")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", x.AsSlice(), expected)
	}
	y, _ := BytesFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", y.AsSlice(), expected)
	}
}

func TestBytesSimple12(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	item1, _ := hex.DecodeString("02")
	x := NewBytesBuilder().Push(*ByteFromSliceUnchecked(item0)).Push(*ByteFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("020000000102")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", x.AsSlice(), expected)
	}
	y, _ := BytesFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", y.AsSlice(), expected)
	}
}

func TestBytesSimple13(t *testing.T) {
	item0, _ := hex.DecodeString("01")
	item1, _ := hex.DecodeString("02")
	item2, _ := hex.DecodeString("03")
	x := NewBytesBuilder().Push(*ByteFromSliceUnchecked(item0)).Push(*ByteFromSliceUnchecked(item1)).Push(*ByteFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000010203")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", x.AsSlice(), expected)
	}
	y, _ := BytesFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Bytes error: ", y.AsSlice(), expected)
	}
}

func TestWordsSimple14(t *testing.T) {
	item0, _ := hex.DecodeString("0123")
	x := NewWordsBuilder().Push(*WordFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("010000000123")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", x.AsSlice(), expected)
	}
	y, _ := WordsFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", y.AsSlice(), expected)
	}
}

func TestWordsSimple15(t *testing.T) {
	item0, _ := hex.DecodeString("0123")
	item1, _ := hex.DecodeString("4567")
	x := NewWordsBuilder().Push(*WordFromSliceUnchecked(item0)).Push(*WordFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("0200000001234567")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", x.AsSlice(), expected)
	}
	y, _ := WordsFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", y.AsSlice(), expected)
	}
}

func TestWordsSimple16(t *testing.T) {
	item0, _ := hex.DecodeString("0123")
	item1, _ := hex.DecodeString("4567")
	item2, _ := hex.DecodeString("89ab")
	x := NewWordsBuilder().Push(*WordFromSliceUnchecked(item0)).Push(*WordFromSliceUnchecked(item1)).Push(*WordFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("030000000123456789ab")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", x.AsSlice(), expected)
	}
	y, _ := WordsFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Words error: ", y.AsSlice(), expected)
	}
}

func TestByte3VecSimple17(t *testing.T) {
	item0, _ := hex.DecodeString("000000")
	x := NewByte3VecBuilder().Push(*Byte3FromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("01000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", y.AsSlice(), expected)
	}
}

func TestByte3VecSimple18(t *testing.T) {
	item0, _ := hex.DecodeString("000000")
	item1, _ := hex.DecodeString("111111")
	x := NewByte3VecBuilder().Push(*Byte3FromSliceUnchecked(item0)).Push(*Byte3FromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("02000000000000111111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", y.AsSlice(), expected)
	}
}

func TestByte3VecSimple19(t *testing.T) {
	item0, _ := hex.DecodeString("000000")
	item1, _ := hex.DecodeString("111111")
	item2, _ := hex.DecodeString("222222")
	x := NewByte3VecBuilder().Push(*Byte3FromSliceUnchecked(item0)).Push(*Byte3FromSliceUnchecked(item1)).Push(*Byte3FromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000000000111111222222")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec error: ", y.AsSlice(), expected)
	}
}

func TestByte7VecSimple20(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	x := NewByte7VecBuilder().Push(*Byte7FromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0100000000000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", y.AsSlice(), expected)
	}
}

func TestByte7VecSimple21(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	item1, _ := hex.DecodeString("11111111111111")
	x := NewByte7VecBuilder().Push(*Byte7FromSliceUnchecked(item0)).Push(*Byte7FromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("020000000000000000000011111111111111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", y.AsSlice(), expected)
	}
}

func TestByte7VecSimple22(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	item1, _ := hex.DecodeString("11111111111111")
	item2, _ := hex.DecodeString("22222222222222")
	x := NewByte7VecBuilder().Push(*Byte7FromSliceUnchecked(item0)).Push(*Byte7FromSliceUnchecked(item1)).Push(*Byte7FromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000000000000000001111111111111122222222222222")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec error: ", y.AsSlice(), expected)
	}
}

func TestStructIVecSimple23(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	x := NewStructIVecBuilder().Push(*StructIFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0100000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructIVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", y.AsSlice(), expected)
	}
}

func TestStructIVecSimple24(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("11111111")
	x := NewStructIVecBuilder().Push(*StructIFromSliceUnchecked(item0)).Push(*StructIFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("020000000000000011111111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructIVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", y.AsSlice(), expected)
	}
}

func TestStructIVecSimple25(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("11111111")
	item2, _ := hex.DecodeString("22222222")
	x := NewStructIVecBuilder().Push(*StructIFromSliceUnchecked(item0)).Push(*StructIFromSliceUnchecked(item1)).Push(*StructIFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000000000001111111122222222")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructIVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructIVec error: ", y.AsSlice(), expected)
	}
}

func TestStructJVecSimple26(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	x := NewStructJVecBuilder().Push(*StructJFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0100000000000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructJVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", y.AsSlice(), expected)
	}
}

func TestStructJVecSimple27(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	item1, _ := hex.DecodeString("11111111111111")
	x := NewStructJVecBuilder().Push(*StructJFromSliceUnchecked(item0)).Push(*StructJFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("020000000000000000000011111111111111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructJVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", y.AsSlice(), expected)
	}
}

func TestStructJVecSimple28(t *testing.T) {
	item0, _ := hex.DecodeString("00000000000000")
	item1, _ := hex.DecodeString("11111111111111")
	item2, _ := hex.DecodeString("22222222222222")
	x := NewStructJVecBuilder().Push(*StructJFromSliceUnchecked(item0)).Push(*StructJFromSliceUnchecked(item1)).Push(*StructJFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000000000000000001111111111111122222222222222")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructJVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructJVec error: ", y.AsSlice(), expected)
	}
}

func TestStructPVecSimple29(t *testing.T) {
	item0, _ := hex.DecodeString("0000000000000000")
	x := NewStructPVecBuilder().Push(*StructPFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("010000000000000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructPVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", y.AsSlice(), expected)
	}
}

func TestStructPVecSimple30(t *testing.T) {
	item0, _ := hex.DecodeString("0000000000000000")
	item1, _ := hex.DecodeString("1111111111111111")
	x := NewStructPVecBuilder().Push(*StructPFromSliceUnchecked(item0)).Push(*StructPFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("0200000000000000000000001111111111111111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructPVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", y.AsSlice(), expected)
	}
}

func TestStructPVecSimple31(t *testing.T) {
	item0, _ := hex.DecodeString("0000000000000000")
	item1, _ := hex.DecodeString("1111111111111111")
	item2, _ := hex.DecodeString("2222222222222222")
	x := NewStructPVecBuilder().Push(*StructPFromSliceUnchecked(item0)).Push(*StructPFromSliceUnchecked(item1)).Push(*StructPFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("03000000000000000000000011111111111111112222222222222222")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", x.AsSlice(), expected)
	}
	y, _ := StructPVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPVec error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecSimple32(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	x := NewBytesVecBuilder().Push(*BytesFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0c0000000800000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecSimple33(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("0100000000")
	x := NewBytesVecBuilder().Push(*BytesFromSliceUnchecked(item0)).Push(*BytesFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("150000000c00000010000000000000000100000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecSimple34(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("0100000000")
	item2, _ := hex.DecodeString("020000000011")
	x := NewBytesVecBuilder().Push(*BytesFromSliceUnchecked(item0)).Push(*BytesFromSliceUnchecked(item1)).Push(*BytesFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("1f000000100000001400000019000000000000000100000000020000000011")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVec error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecSimple35(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	x := NewWordsVecBuilder().Push(*WordsFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0c0000000800000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecSimple36(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("010000000000")
	x := NewWordsVecBuilder().Push(*WordsFromSliceUnchecked(item0)).Push(*WordsFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("160000000c0000001000000000000000010000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecSimple37(t *testing.T) {
	item0, _ := hex.DecodeString("00000000")
	item1, _ := hex.DecodeString("010000000000")
	item2, _ := hex.DecodeString("0200000000001111")
	x := NewWordsVecBuilder().Push(*WordsFromSliceUnchecked(item0)).Push(*WordsFromSliceUnchecked(item1)).Push(*WordsFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("2200000010000000140000001a000000000000000100000000000200000000001111")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVec error: ", y.AsSlice(), expected)
	}
}

func TestByteOptVecSimple38(t *testing.T) {
	item0, _ := hex.DecodeString("")
	x := NewByteOptVecBuilder().Push(*ByteOptFromSliceUnchecked(item0)).Build()

	expected, _ := hex.DecodeString("0800000008000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", y.AsSlice(), expected)
	}
}

func TestByteOptVecSimple39(t *testing.T) {
	item0, _ := hex.DecodeString("")
	item1, _ := hex.DecodeString("")
	x := NewByteOptVecBuilder().Push(*ByteOptFromSliceUnchecked(item0)).Push(*ByteOptFromSliceUnchecked(item1)).Build()

	expected, _ := hex.DecodeString("0c0000000c0000000c000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", y.AsSlice(), expected)
	}
}

func TestByteOptVecSimple40(t *testing.T) {
	item0, _ := hex.DecodeString("")
	item1, _ := hex.DecodeString("")
	item2, _ := hex.DecodeString("")
	x := NewByteOptVecBuilder().Push(*ByteOptFromSliceUnchecked(item0)).Push(*ByteOptFromSliceUnchecked(item1)).Push(*ByteOptFromSliceUnchecked(item2)).Build()

	expected, _ := hex.DecodeString("10000000100000001000000010000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec error: ", y.AsSlice(), expected)
	}
}

func TestTable5Simple41(t *testing.T) {
	f4, _ := hex.DecodeString("020000001234")
	x := NewTable5Builder().F4(*BytesFromSliceUnchecked(f4)).Build()

	expected, _ := hex.DecodeString("2d00000018000000190000001d0000002300000029000000000000000000000000000002000000123404000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table5 error: ", x.AsSlice(), expected)
	}
	y, _ := Table5FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table5 error: ", y.AsSlice(), expected)
	}
}

func TestTable5Simple42(t *testing.T) {
	f4, _ := hex.DecodeString("020000001234")
	f5, _ := hex.DecodeString("1f000000100000001400000019000000000000000100000000020000000011")
	x := NewTable5Builder().F4(*BytesFromSliceUnchecked(f4)).F5(*BytesVecFromSliceUnchecked(f5)).Build()

	expected, _ := hex.DecodeString("4800000018000000190000001d000000230000002900000000000000000000000000000200000012341f000000100000001400000019000000000000000100000000020000000011")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table5 error: ", x.AsSlice(), expected)
	}
	y, _ := Table5FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table5 error: ", y.AsSlice(), expected)
	}
}

func TestTableASimple43(t *testing.T) {
	f1, _ := hex.DecodeString("12345678")
	f2, _ := hex.DecodeString("aabbccddeeff")
	f4, _ := hex.DecodeString("1f000000100000001400000019000000000000000100000000020000000011")
	f6, _ := hex.DecodeString("00000000")
	f7, _ := hex.DecodeString("0300000000000000")
	f8, _ := hex.DecodeString("99")
	x := NewTableABuilder().F1(*Word2FromSliceUnchecked(f1)).F2(*StructAFromSliceUnchecked(f2)).F4(*BytesVecFromSliceUnchecked(f4)).F6(*BytesOptFromSliceUnchecked(f6)).F7(*UnionAFromSliceUnchecked(f7)).F8(*ByteFromSliceUnchecked(f8)).Build()

	expected, _ := hex.DecodeString("6700000024000000280000002e00000032000000510000005a0000005e0000006600000012345678aabbccddeeff000000001f00000010000000140000001900000000000000010000000002000000001109000000080000000000000000030000000000000099")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type TableA error: ", x.AsSlice(), expected)
	}
	y, _ := TableAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type TableA error: ", y.AsSlice(), expected)
	}
}

func TestByteOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("00")
	x := NewByteOptBuilder().Set(*ByteFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("00")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOpt error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOpt error: ", y.AsSlice(), expected)
	}
}

func TestWordOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("0000")
	x := NewWordOptBuilder().Set(*WordFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("0000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordOpt error: ", x.AsSlice(), expected)
	}
	y, _ := WordOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordOpt error: ", y.AsSlice(), expected)
	}
}

func TestStructAOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("000000000000")
	x := NewStructAOptBuilder().Set(*StructAFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructAOpt error: ", x.AsSlice(), expected)
	}
	y, _ := StructAOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructAOpt error: ", y.AsSlice(), expected)
	}
}

func TestStructPOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("0000000000000000")
	x := NewStructPOptBuilder().Set(*StructPFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("0000000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPOpt error: ", x.AsSlice(), expected)
	}
	y, _ := StructPOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPOpt error: ", y.AsSlice(), expected)
	}
}

func TestBytesOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("00000000")
	x := NewBytesOptBuilder().Set(*BytesFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("00000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesOpt error: ", x.AsSlice(), expected)
	}
	y, _ := BytesOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesOpt error: ", y.AsSlice(), expected)
	}
}

func TestWordsOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("00000000")
	x := NewWordsOptBuilder().Set(*WordsFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("00000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsOpt error: ", x.AsSlice(), expected)
	}
	y, _ := WordsOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsOpt error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("04000000")
	x := NewBytesVecOptBuilder().Set(*BytesVecFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("04000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVecOpt error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVecOpt error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("04000000")
	x := NewWordsVecOptBuilder().Set(*WordsVecFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("04000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVecOpt error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVecOpt error: ", y.AsSlice(), expected)
	}
}

func TestTable0OptSimple(t *testing.T) {
	item, _ := hex.DecodeString("04000000")
	x := NewTable0OptBuilder().Set(*Table0FromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("04000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table0Opt error: ", x.AsSlice(), expected)
	}
	y, _ := Table0OptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table0Opt error: ", y.AsSlice(), expected)
	}
}

func TestTable6OptSimple(t *testing.T) {
	item, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	x := NewTable6OptBuilder().Set(*Table6FromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table6Opt error: ", x.AsSlice(), expected)
	}
	y, _ := Table6OptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table6Opt error: ", y.AsSlice(), expected)
	}
}

func TestTable6OptOptSimple(t *testing.T) {
	item, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	x := NewTable6OptOptBuilder().Set(*Table6OptFromSliceUnchecked(item)).Build()

	expected, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table6OptOpt error: ", x.AsSlice(), expected)
	}
	y, _ := Table6OptOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table6OptOpt error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple55(t *testing.T) {
	item, _ := hex.DecodeString("00")
	x := NewUnionABuilder().Set(UnionAUnionFromByte(*ByteFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("0000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple56(t *testing.T) {
	item, _ := hex.DecodeString("0000")
	x := NewUnionABuilder().Set(UnionAUnionFromWord(*WordFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("010000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple57(t *testing.T) {
	item, _ := hex.DecodeString("000000000000")
	x := NewUnionABuilder().Set(UnionAUnionFromStructA(*StructAFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("02000000000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple58(t *testing.T) {
	item, _ := hex.DecodeString("00000000")
	x := NewUnionABuilder().Set(UnionAUnionFromBytes(*BytesFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("0300000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple59(t *testing.T) {
	item, _ := hex.DecodeString("00000000")
	x := NewUnionABuilder().Set(UnionAUnionFromWords(*WordsFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("0400000000000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple60(t *testing.T) {
	item, _ := hex.DecodeString("04000000")
	x := NewUnionABuilder().Set(UnionAUnionFromTable0(*Table0FromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("0500000004000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple61(t *testing.T) {
	item, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	x := NewUnionABuilder().Set(UnionAUnionFromTable6(*Table6FromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("060000005a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple62(t *testing.T) {
	item, _ := hex.DecodeString("")
	x := NewUnionABuilder().Set(UnionAUnionFromTable6Opt(*Table6OptFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("07000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}

func TestUnionASimple63(t *testing.T) {
	item, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	x := NewUnionABuilder().Set(UnionAUnionFromTable6Opt(*Table6OptFromSliceUnchecked(item))).Build()

	expected, _ := hex.DecodeString("070000005a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")

	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA error: ", y.AsSlice(), expected)
	}
}
