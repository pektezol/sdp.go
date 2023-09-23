package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcPrint struct {
	Message string
}

func ParseSvcPrint(reader *bitreader.Reader) SvcPrint {
	svcPrint := SvcPrint{
		Message: reader.TryReadString(),
	}
	// common psycopath behaviour

	return svcPrint
}
