package types

import "github.com/pektezol/bitreader"

type NetSplitScreenUser struct {
	Unknown bool
}

func ParseNetSplitScreenUser(reader *bitreader.ReaderType) NetSplitScreenUser {
	return NetSplitScreenUser{
		Unknown: reader.TryReadBool(),
	}
}
