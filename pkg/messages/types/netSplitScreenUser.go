package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetSplitScreenUser struct {
	Slot bool `json:"slot"`
}

func ParseNetSplitScreenUser(reader *bitreader.Reader, demo *types.Demo) NetSplitScreenUser {
	netSplitScreenUser := NetSplitScreenUser{
		Slot: reader.TryReadBool(),
	}
	demo.Writer.TempAppendLine("\t\tSlot: %t", netSplitScreenUser.Slot)
	return netSplitScreenUser
}
