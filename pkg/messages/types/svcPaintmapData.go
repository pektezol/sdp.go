package messages

import "github.com/pektezol/bitreader"

type SvcPaintmapData struct {
	Length int32
	Data   []byte
}

func ParseSvcPaintmapData(reader *bitreader.ReaderType) SvcPaintmapData {
	svcPaintmapData := SvcPaintmapData{
		Length: int32(reader.TryReadBits(32)),
	}
	svcPaintmapData.Data = reader.TryReadBitsToSlice(int(svcPaintmapData.Length))
	return svcPaintmapData
}
