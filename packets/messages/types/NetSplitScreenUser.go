package types

import "github.com/pektezol/bitreader"

type NetSplitScreenUser struct {
	PlayerSlot bool
}

func ParseNetSplitScreenUser(reader *bitreader.ReaderType) NetSplitScreenUser {
	return NetSplitScreenUser{
		PlayerSlot: reader.TryReadBool(),
	}
}
