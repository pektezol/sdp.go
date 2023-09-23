package classes

import (
	"github.com/pektezol/bitreader"
)

type CustomData struct {
	Type int32
	Size int32
	Data string
}

func (customData *CustomData) ParseCustomData(reader *bitreader.Reader, tickNumber int32, packetType uint8) {
	customData.Type = reader.TryReadSInt32()
	customData.Size = reader.TryReadSInt32()
	if customData.Type != 0 || customData.Size == 8 {
		// Not SAR data

		customData.Data = string(reader.TryReadBytesToSlice(uint64(customData.Size)))

		return
	}
	// SAR data

	sarData := SarData{}
	data := reader.TryReadBytesToSlice(uint64(customData.Size))
	sarReader := bitreader.NewReaderFromBytes(data, true)
	sarData.ParseSarData(sarReader)
}
