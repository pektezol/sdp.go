package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type Stop struct {
	RemainingData []byte `json:"remaining_data"`
}

func (stop *Stop) ParseStop(reader *bitreader.Reader, demo *types.Demo) {
	if reader.TryReadBool() {
		stop.RemainingData = reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
		demo.Writer.AppendLine("\tRemaining Data: %v", stop.RemainingData)
	}
}
