package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type NetDisconnect struct {
	Text string
}

func ParseNetDisconnect(reader *bitreader.Reader) NetDisconnect {
	netDisconnect := NetDisconnect{
		Text: reader.TryReadString(),
	}
	writer.TempAppendLine("\t\tText: %s", netDisconnect.Text)
	return netDisconnect
}
