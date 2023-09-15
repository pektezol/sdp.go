package messages

import "github.com/pektezol/bitreader"

type SvcPaintmapData struct {
	Length int32
	Data   []byte
}

func ParseSvcPaintmapData(reader *bitreader.Reader) SvcPaintmapData {
	svcPaintmapData := SvcPaintmapData{
		Length: int32(reader.TryReadBits(32)),
	}
	svcPaintmapData.Data = reader.TryReadBitsToSlice(uint64(svcPaintmapData.Length))
	return svcPaintmapData
}
