package messages

import (
	"reflect"

	"github.com/pektezol/bitreader"
	messages "github.com/pektezol/sdp.go/pkg/messages/types"
	types "github.com/pektezol/sdp.go/pkg/types"
)

func ParseMessages(messageType uint64, reader *bitreader.Reader, demo *types.Demo) any {
	var messageData any
	switch messageType {
	case 0:
		messageData = messages.ParseNetNop(reader, demo)
	case 1:
		messageData = messages.ParseNetDisconnect(reader, demo)
	case 2:
		messageData = messages.ParseNetFile(reader, demo)
	case 3:
		messageData = messages.ParseNetSplitScreenUser(reader, demo)
	case 4:
		messageData = messages.ParseNetTick(reader, demo)
	case 5:
		messageData = messages.ParseNetStringCmd(reader, demo)
	case 6:
		messageData = messages.ParseNetSetConVar(reader, demo)
	case 7:
		messageData = messages.ParseNetSignOnState(reader, demo)
	case 8:
		messageData = messages.ParseSvcServerInfo(reader, demo)
	case 9:
		messageData = messages.ParseSvcSendTable(reader, demo)
	case 10:
		messageData = messages.ParseSvcClassInfo(reader, demo)
	case 11:
		messageData = messages.ParseSvcSetPause(reader, demo)
	case 12:
		messageData = messages.ParseSvcCreateStringTable(reader, demo) // TODO:
	case 13:
		messageData = messages.ParseSvcUpdateStringTable(reader, demo) // TODO:
	case 14:
		messageData = messages.ParseSvcVoiceInit(reader, demo)
	case 15:
		messageData = messages.ParseSvcVoiceData(reader, demo)
	case 16:
		messageData = messages.ParseSvcPrint(reader, demo)
	case 17:
		messageData = messages.ParseSvcSounds(reader, demo) // TODO:
	case 18:
		messageData = messages.ParseSvcSetView(reader, demo)
	case 19:
		messageData = messages.ParseSvcFixAngle(reader, demo)
	case 20:
		messageData = messages.ParseSvcCrosshairAngle(reader, demo)
	case 21:
		messageData = messages.ParseSvcBspDecal(reader, demo) // untested
	case 22:
		messageData = messages.ParseSvcSplitScreen(reader, demo) // skipped
	case 23:
		messageData = messages.ParseSvcUserMessage(reader, demo)
	case 24:
		messageData = messages.ParseSvcEntityMessage(reader, demo) // skipped
	case 25:
		messageData = messages.ParseSvcGameEvent(reader, demo)
	case 26:
		messageData = messages.ParseSvcPacketEntities(reader, demo) // TODO:
	case 27:
		messageData = messages.ParseSvcTempEntities(reader, demo) // skipped
	case 28:
		messageData = messages.ParseSvcPrefetch(reader, demo)
	case 29:
		messageData = messages.ParseSvcMenu(reader, demo) // skipped
	case 30:
		messageData = messages.ParseSvcGameEventList(reader, demo)
	case 31:
		messageData = messages.ParseSvcGetCvarValue(reader, demo)
	case 32:
		messageData = messages.ParseSvcCmdKeyValues(reader, demo)
	case 33:
		messageData = messages.ParseSvcPaintmapData(reader, demo) // skipped
	default:
		return nil
	}
	demo.Writer.AppendLine("\tMessage: %s (%d):", reflect.ValueOf(messageData).Type(), messageType)
	demo.Writer.AppendOutputFromTemp()
	return messageData
}
