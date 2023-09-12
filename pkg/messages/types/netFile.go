package messages

import "github.com/pektezol/bitreader"

type NetFile struct {
	TransferId    int32
	FileName      string
	FileRequested bool
}

func ParseNetFile(reader *bitreader.ReaderType) NetFile {
	return NetFile{
		TransferId:    int32(reader.TryReadBits(32)),
		FileName:      reader.TryReadString(),
		FileRequested: reader.TryReadBool(),
	}
}
