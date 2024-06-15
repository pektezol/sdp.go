package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/classes"
	"github.com/pektezol/sdp.go/pkg/types"
)

func ParseMessage(reader *bitreader.Reader, demo *types.Demo) types.Message {
	message := types.Message{
		PacketType: types.MessageType(reader.TryReadUInt8()),
		TickNumber: reader.TryReadSInt32(),
		SlotNumber: reader.TryReadUInt8(),
	}
	demo.Writer.AppendLine("[%d] %s (%d):", message.TickNumber, message.PacketType.String(), message.PacketType)
	switch message.PacketType {
	case types.SignOn:
		signOn := classes.SignOn{}
		signOn.ParseSignOn(reader, demo)
		message.Data = signOn
	case types.Packet:
		packet := classes.Packet{}
		packet.ParsePacket(reader, demo)
		message.Data = packet
	case types.SyncTick:
		syncTick := classes.SyncTick{}
		syncTick.ParseSyncTick()
		message.Data = syncTick
	case types.ConsoleCmd:
		consoleCmd := classes.ConsoleCmd{}
		consoleCmd.ParseConsoleCmd(reader, demo)
		message.Data = consoleCmd
	case types.UserCmd:
		userCmd := classes.UserCmd{}
		userCmd.ParseUserCmd(reader, demo)
		message.Data = userCmd
	case types.DataTables:
		dataTables := classes.DataTables{}
		dataTables.ParseDataTables(reader, demo)
		message.Data = dataTables
	case types.Stop:
		stop := classes.Stop{}
		stop.ParseStop(reader, demo)
		message.Data = stop
	case types.CustomData: // TODO: not sar data
		customData := classes.CustomData{}
		customData.ParseCustomData(reader, message.TickNumber, uint8(message.PacketType), demo)
		message.Data = customData
	case types.StringTables: // TODO: parsing string table data
		stringTables := classes.StringTables{}
		stringTables.ParseStringTables(reader, demo)
		message.Data = stringTables
	default:
		panic("invalid packet type")
	}
	return message
}
