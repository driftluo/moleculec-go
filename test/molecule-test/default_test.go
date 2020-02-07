package types

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestByte2Default(t *testing.T) {
	x := Byte2Default()
	expected, _ := hex.DecodeString("0000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte2 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte2FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte2 default error: ", y.AsSlice(), expected)
	}
}

func TestByte3Default(t *testing.T) {
	x := Byte3Default()
	expected, _ := hex.DecodeString("000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3 default error: ", y.AsSlice(), expected)
	}
}

func TestByte4Default(t *testing.T) {
	x := Byte4Default()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte4 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte4FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte4 default error: ", y.AsSlice(), expected)
	}
}

func TestByte5Default(t *testing.T) {
	x := Byte5Default()
	expected, _ := hex.DecodeString("0000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte5 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte5FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte5 default error: ", y.AsSlice(), expected)
	}
}

func TestByte6Default(t *testing.T) {
	x := Byte6Default()
	expected, _ := hex.DecodeString("000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte6 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte6FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte6 default error: ", y.AsSlice(), expected)
	}
}

func TestByte7Default(t *testing.T) {
	x := Byte7Default()
	expected, _ := hex.DecodeString("00000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7 default error: ", y.AsSlice(), expected)
	}
}

func TestByte8Default(t *testing.T) {
	x := Byte8Default()
	expected, _ := hex.DecodeString("0000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte8 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte8FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte8 default error: ", y.AsSlice(), expected)
	}
}

func TestByte9Default(t *testing.T) {
	x := Byte9Default()
	expected, _ := hex.DecodeString("000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte9 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte9FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte9 default error: ", y.AsSlice(), expected)
	}
}

func TestByte10Default(t *testing.T) {
	x := Byte10Default()
	expected, _ := hex.DecodeString("00000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte10 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte10FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte10 default error: ", y.AsSlice(), expected)
	}
}

func TestByte11Default(t *testing.T) {
	x := Byte11Default()
	expected, _ := hex.DecodeString("0000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte11 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte11FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte11 default error: ", y.AsSlice(), expected)
	}
}

func TestByte12Default(t *testing.T) {
	x := Byte12Default()
	expected, _ := hex.DecodeString("000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte12 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte12FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte12 default error: ", y.AsSlice(), expected)
	}
}

func TestByte13Default(t *testing.T) {
	x := Byte13Default()
	expected, _ := hex.DecodeString("00000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte13 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte13FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte13 default error: ", y.AsSlice(), expected)
	}
}

func TestByte14Default(t *testing.T) {
	x := Byte14Default()
	expected, _ := hex.DecodeString("0000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte14 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte14FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte14 default error: ", y.AsSlice(), expected)
	}
}

func TestByte15Default(t *testing.T) {
	x := Byte15Default()
	expected, _ := hex.DecodeString("000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte15 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte15FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte15 default error: ", y.AsSlice(), expected)
	}
}

func TestByte16Default(t *testing.T) {
	x := Byte16Default()
	expected, _ := hex.DecodeString("00000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte16 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte16FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte16 default error: ", y.AsSlice(), expected)
	}
}

func TestWordDefault(t *testing.T) {
	x := WordDefault()
	expected, _ := hex.DecodeString("0000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word default error: ", x.AsSlice(), expected)
	}
	y, _ := WordFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word default error: ", y.AsSlice(), expected)
	}
}

func TestWord2Default(t *testing.T) {
	x := Word2Default()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word2 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word2FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word2 default error: ", y.AsSlice(), expected)
	}
}

func TestWord3Default(t *testing.T) {
	x := Word3Default()
	expected, _ := hex.DecodeString("000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word3 default error: ", y.AsSlice(), expected)
	}
}

func TestWord4Default(t *testing.T) {
	x := Word4Default()
	expected, _ := hex.DecodeString("0000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word4 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word4FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word4 default error: ", y.AsSlice(), expected)
	}
}

func TestWord5Default(t *testing.T) {
	x := Word5Default()
	expected, _ := hex.DecodeString("00000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word5 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word5FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word5 default error: ", y.AsSlice(), expected)
	}
}

func TestWord6Default(t *testing.T) {
	x := Word6Default()
	expected, _ := hex.DecodeString("000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word6 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word6FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word6 default error: ", y.AsSlice(), expected)
	}
}

func TestWord7Default(t *testing.T) {
	x := Word7Default()
	expected, _ := hex.DecodeString("0000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word7 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word7FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word7 default error: ", y.AsSlice(), expected)
	}
}

func TestWord8Default(t *testing.T) {
	x := Word8Default()
	expected, _ := hex.DecodeString("00000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Word8 default error: ", x.AsSlice(), expected)
	}
	y, _ := Word8FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Word8 default error: ", y.AsSlice(), expected)
	}
}

func TestByte3x3Default(t *testing.T) {
	x := Byte3x3Default()
	expected, _ := hex.DecodeString("000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3x3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3x3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3x3 default error: ", y.AsSlice(), expected)
	}
}

func TestByte5x3Default(t *testing.T) {
	x := Byte5x3Default()
	expected, _ := hex.DecodeString("000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte5x3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte5x3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte5x3 default error: ", y.AsSlice(), expected)
	}
}

func TestByte7x3Default(t *testing.T) {
	x := Byte7x3Default()
	expected, _ := hex.DecodeString("000000000000000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7x3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7x3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7x3 default error: ", y.AsSlice(), expected)
	}
}

func TestByte9x3Default(t *testing.T) {
	x := Byte9x3Default()
	expected, _ := hex.DecodeString("000000000000000000000000000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte9x3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte9x3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte9x3 default error: ", y.AsSlice(), expected)
	}
}

func TestStructIx3Default(t *testing.T) {
	x := StructIx3Default()
	expected, _ := hex.DecodeString("000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructIx3 default error: ", x.AsSlice(), expected)
	}
	y, _ := StructIx3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructIx3 default error: ", y.AsSlice(), expected)
	}
}

func TestStructADefault(t *testing.T) {
	x := StructADefault()
	expected, _ := hex.DecodeString("000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructA default error: ", x.AsSlice(), expected)
	}
	y, _ := StructAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructA default error: ", y.AsSlice(), expected)
	}
}

func TestStructBDefault(t *testing.T) {
	x := StructBDefault()
	expected, _ := hex.DecodeString("00000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructB default error: ", x.AsSlice(), expected)
	}
	y, _ := StructBFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructB default error: ", y.AsSlice(), expected)
	}
}

func TestStructCDefault(t *testing.T) {
	x := StructCDefault()
	expected, _ := hex.DecodeString("0000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructC default error: ", x.AsSlice(), expected)
	}
	y, _ := StructCFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructC default error: ", y.AsSlice(), expected)
	}
}

func TestStructDDefault(t *testing.T) {
	x := StructDDefault()
	expected, _ := hex.DecodeString("000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructD default error: ", x.AsSlice(), expected)
	}
	y, _ := StructDFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructD default error: ", y.AsSlice(), expected)
	}
}

func TestStructEDefault(t *testing.T) {
	x := StructEDefault()
	expected, _ := hex.DecodeString("000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructE default error: ", x.AsSlice(), expected)
	}
	y, _ := StructEFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructE default error: ", y.AsSlice(), expected)
	}
}

func TestStructFDefault(t *testing.T) {
	x := StructFDefault()
	expected, _ := hex.DecodeString("0000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructF default error: ", x.AsSlice(), expected)
	}
	y, _ := StructFFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructF default error: ", y.AsSlice(), expected)
	}
}

func TestStructGDefault(t *testing.T) {
	x := StructGDefault()
	expected, _ := hex.DecodeString("00000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructG default error: ", x.AsSlice(), expected)
	}
	y, _ := StructGFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructG default error: ", y.AsSlice(), expected)
	}
}

func TestStructHDefault(t *testing.T) {
	x := StructHDefault()
	expected, _ := hex.DecodeString("00000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructH default error: ", x.AsSlice(), expected)
	}
	y, _ := StructHFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructH default error: ", y.AsSlice(), expected)
	}
}

func TestStructIDefault(t *testing.T) {
	x := StructIDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructI default error: ", x.AsSlice(), expected)
	}
	y, _ := StructIFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructI default error: ", y.AsSlice(), expected)
	}
}

func TestStructJDefault(t *testing.T) {
	x := StructJDefault()
	expected, _ := hex.DecodeString("00000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructJ default error: ", x.AsSlice(), expected)
	}
	y, _ := StructJFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructJ default error: ", y.AsSlice(), expected)
	}
}

func TestStructODefault(t *testing.T) {
	x := StructODefault()
	expected, _ := hex.DecodeString("00000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructO default error: ", x.AsSlice(), expected)
	}
	y, _ := StructOFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructO default error: ", y.AsSlice(), expected)
	}
}

func TestStructPDefault(t *testing.T) {
	x := StructPDefault()
	expected, _ := hex.DecodeString("0000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructP default error: ", x.AsSlice(), expected)
	}
	y, _ := StructPFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructP default error: ", y.AsSlice(), expected)
	}
}

func TestBytesDefault(t *testing.T) {
	x := BytesDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Bytes default error: ", x.AsSlice(), expected)
	}
	y, _ := BytesFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Bytes default error: ", y.AsSlice(), expected)
	}
}

func TestWordsDefault(t *testing.T) {
	x := WordsDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Words default error: ", x.AsSlice(), expected)
	}
	y, _ := WordsFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Words default error: ", y.AsSlice(), expected)
	}
}

func TestByte3VecDefault(t *testing.T) {
	x := Byte3VecDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte3VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte3Vec default error: ", y.AsSlice(), expected)
	}
}

func TestByte7VecDefault(t *testing.T) {
	x := Byte7VecDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec default error: ", x.AsSlice(), expected)
	}
	y, _ := Byte7VecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Byte7Vec default error: ", y.AsSlice(), expected)
	}
}

func TestStructJVecDefault(t *testing.T) {
	x := StructJVecDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructJVec default error: ", x.AsSlice(), expected)
	}
	y, _ := StructJVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructJVec default error: ", y.AsSlice(), expected)
	}
}

func TestStructPVecDefault(t *testing.T) {
	x := StructPVecDefault()
	expected, _ := hex.DecodeString("00000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPVec default error: ", x.AsSlice(), expected)
	}
	y, _ := StructPVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPVec default error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecDefault(t *testing.T) {
	x := BytesVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVec default error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVec default error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecDefault(t *testing.T) {
	x := WordsVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVec default error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVec default error: ", y.AsSlice(), expected)
	}
}

func TestByteOptVecDefault(t *testing.T) {
	x := ByteOptVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec default error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOptVec default error: ", y.AsSlice(), expected)
	}
}

func TestWordOptVecDefault(t *testing.T) {
	x := WordOptVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordOptVec default error: ", x.AsSlice(), expected)
	}
	y, _ := WordOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordOptVec default error: ", y.AsSlice(), expected)
	}
}

func TestWordsOptVecDefault(t *testing.T) {
	x := WordsOptVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsOptVec default error: ", x.AsSlice(), expected)
	}
	y, _ := WordsOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsOptVec default error: ", y.AsSlice(), expected)
	}
}

func TestBytesOptVecDefault(t *testing.T) {
	x := BytesOptVecDefault()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesOptVec default error: ", x.AsSlice(), expected)
	}
	y, _ := BytesOptVecFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesOptVec default error: ", y.AsSlice(), expected)
	}
}

func TestTable0Default(t *testing.T) {
	x := Table0Default()
	expected, _ := hex.DecodeString("04000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table0 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table0FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table0 default error: ", y.AsSlice(), expected)
	}
}

func TestTable1Default(t *testing.T) {
	x := Table1Default()
	expected, _ := hex.DecodeString("090000000800000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table1 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table1FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table1 default error: ", y.AsSlice(), expected)
	}
}

func TestTable2Default(t *testing.T) {
	x := Table2Default()
	expected, _ := hex.DecodeString("110000000c0000000d0000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table2 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table2FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table2 default error: ", y.AsSlice(), expected)
	}
}

func TestTable3Default(t *testing.T) {
	x := Table3Default()
	expected, _ := hex.DecodeString("1b0000001000000011000000150000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table3 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table3FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table3 default error: ", y.AsSlice(), expected)
	}
}

func TestTable4Default(t *testing.T) {
	x := Table4Default()
	expected, _ := hex.DecodeString("230000001400000015000000190000001f000000000000000000000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table4 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table4FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table4 default error: ", y.AsSlice(), expected)
	}
}

func TestTable5Default(t *testing.T) {
	x := Table5Default()
	expected, _ := hex.DecodeString("2b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table5 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table5FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table5 default error: ", y.AsSlice(), expected)
	}
}

func TestTable6Default(t *testing.T) {
	x := Table6Default()
	expected, _ := hex.DecodeString("5a0000001c0000001d00000021000000270000002b0000002f000000000000000000000000000000000000040000002b00000018000000190000001d000000230000002700000000000000000000000000000000000004000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table6 default error: ", x.AsSlice(), expected)
	}
	y, _ := Table6FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table6 default error: ", y.AsSlice(), expected)
	}
}

func TestTableADefault(t *testing.T) {
	x := TableADefault()
	expected, _ := hex.DecodeString("4500000024000000280000002e00000032000000360000003f0000003f00000044000000000000000000000000000000000004000000090000000800000000000000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type TableA default error: ", x.AsSlice(), expected)
	}
	y, _ := TableAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type TableA default error: ", y.AsSlice(), expected)
	}
}

func TestByteOptDefault(t *testing.T) {
	x := ByteOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type ByteOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := ByteOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type ByteOpt default error: ", y.AsSlice(), expected)
	}
}

func TestWordOptDefault(t *testing.T) {
	x := WordOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := WordOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordOpt default error: ", y.AsSlice(), expected)
	}
}

func TestStructAOptDefault(t *testing.T) {
	x := StructAOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructAOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := StructAOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructAOpt default error: ", y.AsSlice(), expected)
	}
}

func TestStructPOptDefault(t *testing.T) {
	x := StructPOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type StructPOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := StructPOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type StructPOpt default error: ", y.AsSlice(), expected)
	}
}

func TestBytesOptDefault(t *testing.T) {
	x := BytesOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := BytesOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesOpt default error: ", y.AsSlice(), expected)
	}
}

func TestWordsOptDefault(t *testing.T) {
	x := WordsOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := WordsOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsOpt default error: ", y.AsSlice(), expected)
	}
}

func TestBytesVecOptDefault(t *testing.T) {
	x := BytesVecOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type BytesVecOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := BytesVecOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type BytesVecOpt default error: ", y.AsSlice(), expected)
	}
}

func TestWordsVecOptDefault(t *testing.T) {
	x := WordsVecOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type WordsVecOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := WordsVecOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type WordsVecOpt default error: ", y.AsSlice(), expected)
	}
}

func TestTable0OptDefault(t *testing.T) {
	x := Table0OptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table0Opt default error: ", x.AsSlice(), expected)
	}
	y, _ := Table0OptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table0Opt default error: ", y.AsSlice(), expected)
	}
}

func TestTable6OptDefault(t *testing.T) {
	x := Table6OptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table6Opt default error: ", x.AsSlice(), expected)
	}
	y, _ := Table6OptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table6Opt default error: ", y.AsSlice(), expected)
	}
}

func TestTable6OptOptDefault(t *testing.T) {
	x := Table6OptOptDefault()
	expected, _ := hex.DecodeString("")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type Table6OptOpt default error: ", x.AsSlice(), expected)
	}
	y, _ := Table6OptOptFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type Table6OptOpt default error: ", y.AsSlice(), expected)
	}
}

func TestUnionADefault(t *testing.T) {
	x := UnionADefault()
	expected, _ := hex.DecodeString("0000000000")
	if bytes.Compare(x.AsSlice(), expected) != 0 {
		t.Error("type UnionA default error: ", x.AsSlice(), expected)
	}
	y, _ := UnionAFromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {
		t.Error("type UnionA default error: ", y.AsSlice(), expected)
	}
}
