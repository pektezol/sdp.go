package utils

import (
	"encoding/binary"
	"log"
	"math"
	"os"
	"unsafe"

	"github.com/potterxu/bitreader"
)

func ReadByteFromFile(file *os.File, size int32) []byte {
	tmp := make([]byte, size)
	file.Read(tmp)
	return tmp
}

func CheckError(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func CheckFirstBit(byteArr []byte) bool {
	reader := bitreader.BitReader(byteArr)
	state, err := reader.ReadBit()
	if err != nil {
		state = false
	}
	return state
}

func IntFromBytes(byteArr []byte) uint32 {
	int := binary.LittleEndian.Uint32(byteArr)
	return int
}

func FloatFromBytes(byteArr []byte) float32 {
	bits := binary.LittleEndian.Uint32(byteArr)
	float := math.Float32frombits(bits)
	return float
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