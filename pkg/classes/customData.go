package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type CustomData struct {
	Type int32 `json:"type"`
	Size int32 `json:"size"`
	Data any   `json:"data"`
}

func (customData *CustomData) ParseCustomData(reader *bitreader.Reader, tickNumber int32, packetType uint8, demo *types.Demo) {
	customData.Type = reader.TryReadSInt32()
	customData.Size = reader.TryReadSInt32()
	if customData.Type != 0 || customData.Size == 8 {
		// Not SAR data
		demo.Writer.AppendLine("[%d] %s (%d):", tickNumber, "CUSTOMDATA", packetType)
		customData.Data = string(reader.TryReadBytesToSlice(uint64(customData.Size)))
		demo.Writer.AppendLine("\t%s", customData.Data)
		return
	}
	// SAR data
	demo.Writer.AppendLine("[%d] %s (%d):", tickNumber, "SARDATA", packetType)
	sarData := SarData{}
	data := reader.TryReadBytesToSlice(uint64(customData.Size))
	sarReader := bitreader.NewReaderFromBytes(data, true)
	sarData.ParseSarData(sarReader, demo)
	customData.Data = sarData
}
