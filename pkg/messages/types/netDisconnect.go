package messages

import (
	"github.com/pektezol/bitreader"
)

type NetDisconnect struct {
	Text string
}

func ParseNetDisconnect(reader *bitreader.Reader) NetDisconnect {
	netDisconnect := NetDisconnect{
		Text: reader.TryReadString(),
	}

	return netDisconnect
}
