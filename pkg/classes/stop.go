package classes

import (
	"github.com/pektezol/bitreader"
)

type Stop struct {
	RemainingData []byte
}

func (stop *Stop) ParseStop(reader *bitreader.Reader) {
	if reader.TryReadBool() {
		stop.RemainingData = reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
	}
}
