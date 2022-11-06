package utils

import (
	"bytes"
	"os"
	"unsafe"

	"github.com/pektezol/bitreader"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadByteFromFile(file *os.File, size int32) []byte {
	tmp := make([]byte, size)
	file.Read(tmp)
	return tmp
}

func ReadStringFromFile(file *os.File) string {
	var output string
	reader := bitreader.Reader(file, true)
	for {
		value, err := reader.ReadBytes(1)
		CheckError(err)
		if value == 0 {
			break
		}
		output += string(rune(value))
	}
	return output
}

func ReadStringFromSlice(file []byte) string {
	var output string
	reader := bitreader.Reader(bytes.NewReader(file), true)
	for {
		value, err := reader.ReadBytes(1)
		CheckError(err)
		if value == 0 {
			break
		}
		output += string(rune(value))
	}
	return output
}

func FloatArrFromBytes(byteArr []byte) []float32 {
	if len(byteArr) == 0 {
		return nil
	}
	l := len(byteArr) / 4
	ptr := unsafe.Pointer(&byteArr[0])
	// It is important to keep in mind that the Go garbage collector
	// will not interact with this data, and that if src if freed,
	// the behavior of any Go code using the slice is nondeterministic.
	// Reference: https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	return (*[1 << 26]float32)((*[1 << 26]float32)(ptr))[:l:l]
}
