package types

import "github.com/pektezol/bitreader"

type NetFile struct {
	TransferId    uint32
	FileName      string
	FileRequested bool
}

func ParseNetFile(reader *bitreader.ReaderType) NetFile {
	return NetFile{
		TransferId:    reader.TryReadInt32(),
		FileName:      reader.TryReadString(),
		FileRequested: reader.TryReadBool(),
	}
}
