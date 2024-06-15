package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetDisconnect struct {
	Text string `json:"text"`
}

func ParseNetDisconnect(reader *bitreader.Reader, demo *types.Demo) NetDisconnect {
	netDisconnect := NetDisconnect{
		Text: reader.TryReadString(),
	}
	demo.Writer.TempAppendLine("\t\tText: %s", netDisconnect.Text)
	return netDisconnect
}
