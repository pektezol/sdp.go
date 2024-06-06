package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type NetSplitScreenUser struct {
	Slot bool
}

func ParseNetSplitScreenUser(reader *bitreader.Reader) NetSplitScreenUser {
	netSplitScreenUser := NetSplitScreenUser{
		Slot: reader.TryReadBool(),
	}
	writer.TempAppendLine("\t\tSlot: %t", netSplitScreenUser.Slot)
	return netSplitScreenUser
}
