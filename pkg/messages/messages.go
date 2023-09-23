package messages

import (
	"reflect"

	"github.com/pektezol/bitreader"
	messages "github.com/pektezol/demoparser/pkg/messages/types"
	"github.com/pektezol/demoparser/pkg/writer"
)

func ParseMessages(messageType uint64, reader *bitreader.Reader) any {
	var messageData any
	switch messageType {
	case 0:
		messageData = messages.ParseNetNop(reader)
	case 1:
		messageData = messages.ParseNetDisconnect(reader)
	case 2:
		messageData = messages.ParseNetFile(reader)
	case 3:
		messageData = messages.ParseNetSplitScreenUser(reader)
	case 4:
		messageData = messages.ParseNetTick(reader)
	case 5:
		messageData = messages.ParseNetStringCmd(reader)
	case 6:
		messageData = messages.ParseNetSetConVar(reader)
	case 7:
		messageData = messages.ParseNetSignOnState(reader)
	case 8:
		messageData = messages.ParseSvcServerInfo(reader)
	case 9:
		messageData = messages.ParseSvcSendTable(reader)
	case 10:
		messageData = messages.ParseSvcClassInfo(reader)
	case 11:
		messageData = messages.ParseSvcSetPause(reader)
	case 12:
		messageData = messages.ParseSvcCreateStringTable(reader) // TODO:
	case 13:
		messageData = messages.ParseSvcUpdateStringTable(reader) // TODO:
	case 14:
		messageData = messages.ParseSvcVoiceInit(reader)
	case 15:
		messageData = messages.ParseSvcVoiceData(reader)
	case 16:
		messageData = messages.ParseSvcPrint(reader)
	case 17:
		messageData = messages.ParseSvcSounds(reader) // TODO:
	case 18:
		messageData = messages.ParseSvcSetView(reader)
	case 19:
		messageData = messages.ParseSvcFixAngle(reader)
	case 20:
		messageData = messages.ParseSvcCrosshairAngle(reader)
	case 21:
		messageData = messages.ParseSvcBspDecal(reader) // untested
	case 22:
		messageData = messages.ParseSvcSplitScreen(reader) // skipped
	case 23:
		messageData = messages.ParseSvcUserMessage(reader)
	case 24:
		messageData = messages.ParseSvcEntityMessage(reader) // skipped
	case 25:
		messageData = messages.ParseSvcGameEvent(reader) // TODO:
	case 26:
		messageData = messages.ParseSvcPacketEntities(reader) // TODO:
	case 27:
		messageData = messages.ParseSvcTempEntities(reader) // skipped
	case 28:
		messageData = messages.ParseSvcPrefetch(reader)
	case 29:
		messageData = messages.ParseSvcMenu(reader) // skipped
	case 30:
		messageData = messages.ParseSvcGameEventList(reader) // TODO:
	case 31:
		messageData = messages.ParseSvcGetCvarValue(reader)
	case 32:
		messageData = messages.ParseSvcCmdKeyValues(reader)
	case 33:
		messageData = messages.ParseSvcPaintmapData(reader)
	default:
		return nil
	}
	writer.AppendLine("\tMessage: %s (%d):", reflect.ValueOf(messageData).Type(), messageType)
	writer.AppendOutputFromTemp()
	return messageData
}
