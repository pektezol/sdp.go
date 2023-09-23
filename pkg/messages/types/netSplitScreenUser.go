package messages

import (
	"github.com/pektezol/bitreader"
)

type NetSplitScreenUser struct {
	Slot bool
}

func ParseNetSplitScreenUser(reader *bitreader.Reader) NetSplitScreenUser {
	netSplitScreenUser := NetSplitScreenUser{
		Slot: reader.TryReadBool(),
	}

	return netSplitScreenUser
}
