package messages

import "github.com/pektezol/bitreader"

type NetSplitScreenUser struct {
	Unknown bool
}

func ParseNetSplitScreenUser(reader *bitreader.Reader) NetSplitScreenUser {
	return NetSplitScreenUser{
		Unknown: reader.TryReadBool(),
	}
}
