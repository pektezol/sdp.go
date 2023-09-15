package messages

import "github.com/pektezol/bitreader"

type NetDisconnect struct {
	Text string
}

func ParseNetDisconnect(reader *bitreader.Reader) NetDisconnect {
	return NetDisconnect{
		Text: reader.TryReadString(),
	}
}
