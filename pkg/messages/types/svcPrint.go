package messages

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type SvcPrint struct {
	Message string
}

func ParseSvcPrint(reader *bitreader.Reader) SvcPrint {
	svcPrint := SvcPrint{
		Message: reader.TryReadString(),
	}
	// common psycopath behaviour
	writer.TempAppendLine("\t\t%s", strings.Replace(strings.ReplaceAll(strings.ReplaceAll(svcPrint.Message, "\n", "\n\t\t"), "\n\t\t\n\t\t", ""), "\n\t\t", "", 1))
	return svcPrint
}
