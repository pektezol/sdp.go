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
		case 0x12:
			messages = append(messages, Message{Data: types.ParseSvcCreateStringTable(reader)})
		case 0x13:
			messages = append(messages, Message{Data: types.ParseSvcUpdateStringTable(reader)})
		case 0x14:
			messages = append(messages, Message{Data: types.ParseSvcVoiceInit(reader)})
		case 0x15:
			messages = append(messages, Message{Data: types.ParseSvcVoiceData(reader)})
		case 0x16:
			messages = append(messages, Message{Data: types.ParseSvcPrint(reader)})
		case 0x17:
			messages = append(messages, Message{Data: types.ParseSvcSounds(reader)})
		case 0x18:
			messages = append(messages, Message{Data: types.ParseSvcSetView(reader)})
		case 0x19:
			messages = append(messages, Message{Data: types.ParseSvcFixAngle(reader)})
		case 0x20:
			messages = append(messages, Message{Data: types.ParseSvcCrosshairAngle(reader)})
		case 0x21:
			// TODO: SvcBspDecal
		case 0x22:
			messages = append(messages, Message{Data: types.ParseSvcSplitScreen(reader)})
		case 0x23:
			messages = append(messages, Message{Data: types.ParseSvcUserMessage(reader)})
		case 0x24:
			messages = append(messages, Message{Data: types.ParseSvcEntityMessage(reader)})
		case 0x25:
			// TODO: SvcGameEvent
		case 0x26:
			messages = append(messages, Message{Data: types.ParseSvcPacketEntities(reader)})
		case 0x27:
			messages = append(messages, Message{Data: types.ParseSvcTempEntities(reader)})
		case 0x28:
			messages = append(messages, Message{Data: types.ParseSvcPrefetch(reader)})
		case 0x29:
			messages = append(messages, Message{Data: types.ParseSvcMenu(reader)})
		case 0x30:
			messages = append(messages, Message{Data: types.ParseSvcGameEventList(reader)})
		case 0x31:
			messages = append(messages, Message{Data: types.ParseSvcGetCvarValue(reader)})
		case 0x32:
			messages = append(messages, Message{Data: types.ParseSvcCmdKeyValues(reader)})
		case 0x33:
			messages = append(messages, Message{Data: types.ParseSvcPaintmapData(reader)})
		}
	}
	return messages
}

type Message struct {
	Data any
}
