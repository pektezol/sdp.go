package types

import "github.com/pektezol/bitreader"

type SvcPaintmapData struct {
	Data []byte
}

func ParseSvcPaintmapData(reader *bitreader.ReaderType) SvcPaintmapData {
	length := reader.TryReadInt32()
	return SvcPaintmapData{
		Data: reader.TryReadBytesToSlice(int(length / 8)),
	}
}
