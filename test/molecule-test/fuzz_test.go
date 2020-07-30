package types

import (
	"bytes"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func decode_fuzz_test(data []byte) {
	var byte16, err1 = Byte16FromSlice(data, false)
	if err1 == nil {
		if bytes.Compare(byte16.AsSlice(), data) != 0 {
			panic("byte16")
		}
	}

	var word8, err2 = Word8FromSlice(data, false)
	if err2 == nil {
		if bytes.Compare(word8.AsSlice(), data) != 0 {
			panic("word8")
		}
	}

	var byte9x3, err3 = Byte9x3FromSlice(data, false)
	if err3 == nil {
		if bytes.Compare(byte9x3.AsSlice(), data) != 0 {
			panic("byte9x3")
		}
	}

	var structc, err4 = StructCFromSlice(data, false)
	if err4 == nil {
		if bytes.Compare(structc.AsSlice(), data) != 0 {
			panic("structc")
		}
	}

	var structpvec, err5 = StructPVecFromSlice(data, false)
	if err5 == nil {
		if bytes.Compare(structpvec.AsSlice(), data) != 0 {
			panic("structpvec")
		}
	}

	var table0, err6 = Table0FromSlice(data, false)
	if err6 == nil {
		if bytes.Compare(table0.AsSlice(), data) != 0 {
			panic("table0")
		}
	}

	var table6, err7 = Table6FromSlice(data, false)
	if err7 == nil {
		if bytes.Compare(table6.AsSlice(), data) != 0 {
			panic("table6")
		}
	}

	var wordsvecopt, err8 = WordsVecOptFromSlice(data, false)
	if err8 == nil {
		if bytes.Compare(wordsvecopt.AsSlice(), data) != 0 {
			panic("wordsvecopt")
		}
	}

	var bytesoptvec, err9 = BytesOptVecFromSlice(data, false)
	if err9 == nil {
		if bytes.Compare(bytesoptvec.AsSlice(), data) != 0 {
			panic("bytesoptvec")
		}
	}

	var uniona, err10 = UnionAFromSlice(data, false)
	if err10 == nil {
		if bytes.Compare(uniona.AsSlice(), data) != 0 {
			panic("uniona")
		}
	}
}

func TestDecodePanic(t *testing.T) {
	f := fuzz.New()
	var inputString string
	for i := 0; i < 1000000; i++ {
		f.Fuzz(&inputString)
		go decode_fuzz_test([]byte(inputString))
	}
}
