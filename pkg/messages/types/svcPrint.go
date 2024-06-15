package messages

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcPrint struct {
	Message string `json:"message"`
}

func ParseSvcPrint(reader *bitreader.Reader, demo *types.Demo) SvcPrint {
	svcPrint := SvcPrint{
		Message: reader.TryReadString(),
	}
	// common psycopath behaviour
	demo.Writer.TempAppendLine("\t\t%s", strings.Replace(strings.ReplaceAll(strings.ReplaceAll(svcPrint.Message, "\n", "\n\t\t"), "\n\t\t\n\t\t", ""), "\n\t\t", "", 1))
	return svcPrint
}
