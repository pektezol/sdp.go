package messages

import (
	"bytes"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/packets/messages/types"
)

func ParseMessage(data []byte) []Message {
	reader := bitreader.Reader(bytes.NewReader(data), true)
	var messages []Message
	for {
		messageType, err := reader.ReadBits(6)
		if err != nil {
			break
		}
		switch messageType {
		case 0x00:
			messages = append(messages, Message{Data: types.NetNop{}})
		case 0x01:
			messages = append(messages, Message{Data: types.ParseNetDisconnect(reader)})
		case 0x02:
			messages = append(messages, Message{Data: types.ParseNetFile(reader)})
		case 0x03:
			messages = append(messages, Message{Data: types.ParseNetSplitScreenUser(reader)})
		case 0x04:
			messages = append(messages, Message{Data: types.ParseNetTick(reader)})
		case 0x05:
			messages = append(messages, Message{Data: types.ParseNetStringCmd(reader)})
		case 0x06:
			messages = append(messages, Message{Data: types.ParseNetSetConVar(reader)})
		case 0x07:
			messages = append(messages, Message{Data: types.ParseNetSignOnState(reader)})
		case 0x08:
			messages = append(messages, Message{Data: types.ParseSvcServerInfo(reader)})
		case 0x09:
			messages = append(messages, Message{Data: types.ParseSvcSendTable(reader)})
		case 0x10:
			messages = append(messages, Message{Data: types.ParseSvcClassInfo(reader)})
		case 0x11:
			messages = append(messages, Message{Data: types.ParseSvcSetPause(reader)})
		}
	}
	return messages
}

type Message struct {
	Data any
}
