package messages

import "github.com/pektezol/bitreader"

type SvcPrint struct {
	Message string
}

func ParseSvcPrint(reader *bitreader.ReaderType) SvcPrint {
	return SvcPrint{
		Message: reader.TryReadString(),
	}
}
