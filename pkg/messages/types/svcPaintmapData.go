package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcPaintmapData struct {
	Length uint32
	Data   []byte
}

func ParseSvcPaintmapData(reader *bitreader.Reader) SvcPaintmapData {
	svcPaintmapData := SvcPaintmapData{
		Length: reader.TryReadUInt32(),
	}
	svcPaintmapData.Data = reader.TryReadBitsToSlice(uint64(svcPaintmapData.Length))

	return svcPaintmapData
}
