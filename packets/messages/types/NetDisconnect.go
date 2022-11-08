package types

import "github.com/pektezol/bitreader"

type NetDisconnect struct {
	Text string
}

func ParseNetDisconnect(reader *bitreader.ReaderType) NetDisconnect {
	return NetDisconnect{Text: reader.TryReadString()}
}
