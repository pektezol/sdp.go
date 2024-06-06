package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type Stop struct {
	RemainingData []byte
}

func (stop *Stop) ParseStop(reader *bitreader.Reader) {
	if reader.TryReadBool() {
		stop.RemainingData = reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
		writer.AppendLine("\tRemaining Data: %v", stop.RemainingData)
	}
}
