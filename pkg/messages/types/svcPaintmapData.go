package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcPaintmapData struct {
	Length uint32 `json:"length"`
	Data   []byte `json:"data"`
}

func ParseSvcPaintmapData(reader *bitreader.Reader, demo *types.Demo) SvcPaintmapData {
	svcPaintmapData := SvcPaintmapData{
		Length: reader.TryReadUInt32(),
	}
	svcPaintmapData.Data = reader.TryReadBitsToSlice(uint64(svcPaintmapData.Length))
	demo.Writer.TempAppendLine("\t\tData: %v", svcPaintmapData.Data)
	return svcPaintmapData
}
