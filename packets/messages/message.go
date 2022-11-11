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
		case 0:
			messages = append(messages, Message{Data: types.NetNop{}})
		case 1:
			messages = append(messages, Message{Data: types.ParseNetDisconnect(reader)})
		case 2:
			messages = append(messages, Message{Data: types.ParseNetFile(reader)})
		case 3:
			messages = append(messages, Message{Data: types.ParseNetSplitScreenUser(reader)})
		case 4:
			messages = append(messages, Message{Data: types.ParseNetTick(reader)})
		case 5:
			messages = append(messages, Message{Data: types.ParseNetStringCmd(reader)})
		case 6:
			messages = append(messages, Message{Data: types.ParseNetSetConVar(reader)})
		case 7:
			messages = append(messages, Message{Data: types.ParseNetSignOnState(reader)})
		case 8:
			messages = append(messages, Message{Data: types.ParseSvcServerInfo(reader)})
		case 9:
			messages = append(messages, Message{Data: types.ParseSvcSendTable(reader)})
		case 10:
			messages = append(messages, Message{Data: types.ParseSvcClassInfo(reader)})
		case 11:
			messages = append(messages, Message{Data: types.ParseSvcSetPause(reader)})
		case 12:
			messages = append(messages, Message{Data: types.ParseSvcCreateStringTable(reader)})
		case 13:
			messages = append(messages, Message{Data: types.ParseSvcUpdateStringTable(reader)})
		case 14:
			messages = append(messages, Message{Data: types.ParseSvcVoiceInit(reader)})
		case 15:
			messages = append(messages, Message{Data: types.ParseSvcVoiceData(reader)})
		case 16:
			messages = append(messages, Message{Data: types.ParseSvcPrint(reader)})
		case 17:
			messages = append(messages, Message{Data: types.ParseSvcSounds(reader)})
		case 18:
			messages = append(messages, Message{Data: types.ParseSvcSetView(reader)})
		case 19:
			messages = append(messages, Message{Data: types.ParseSvcFixAngle(reader)})
		case 20:
			messages = append(messages, Message{Data: types.ParseSvcCrosshairAngle(reader)})
		case 21:
			// TODO: SvcBspDecal
		case 22:
			messages = append(messages, Message{Data: types.ParseSvcSplitScreen(reader)})
		case 23:
			messages = append(messages, Message{Data: types.ParseSvcUserMessage(reader)})
		case 24:
			messages = append(messages, Message{Data: types.ParseSvcEntityMessage(reader)})
		case 25:
			messages = append(messages, Message{Data: types.ParseSvcGameEvent(reader)})
		case 26:
			messages = append(messages, Message{Data: types.ParseSvcPacketEntities(reader)})
		case 27:
			messages = append(messages, Message{Data: types.ParseSvcTempEntities(reader)})
		case 28:
			messages = append(messages, Message{Data: types.ParseSvcPrefetch(reader)})
		case 29:
			messages = append(messages, Message{Data: types.ParseSvcMenu(reader)})
		case 30:
			messages = append(messages, Message{Data: types.ParseSvcGameEventList(reader)})
		case 31:
			messages = append(messages, Message{Data: types.ParseSvcGetCvarValue(reader)})
		case 32:
			messages = append(messages, Message{Data: types.ParseSvcCmdKeyValues(reader)})
		case 33:
			messages = append(messages, Message{Data: types.ParseSvcPaintmapData(reader)})
		}
	}
	return messages
}

type Message struct {
	Data any
}
