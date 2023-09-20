package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tData: %v", svcPaintmapData.Data)
	return svcPaintmapData
}
