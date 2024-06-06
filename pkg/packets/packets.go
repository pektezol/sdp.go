package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/classes"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type MessageType uint8

const (
	SignOn MessageType = iota + 1
	Packet
	SyncTick
	ConsoleCmd
	UserCmd
	DataTables
	Stop
	CustomData
	StringTables
)

type Message struct {
	PacketType MessageType
	TickNumber int32
	SlotNumber uint8
	Data       any
}

func ParseMessage(reader *bitreader.Reader) Message {
	message := Message{
		PacketType: MessageType(reader.TryReadUInt8()),
		TickNumber: reader.TryReadSInt32(),
		SlotNumber: reader.TryReadUInt8(),
	}
	writer.AppendLine("[%d] %s (%d):", message.TickNumber, message.PacketType.String(), message.PacketType)
	switch message.PacketType {
	case SignOn:
		signOn := classes.SignOn{}
		signOn.ParseSignOn(reader)
		message.Data = signOn
	case Packet:
		packet := classes.Packet{}
		packet.ParsePacket(reader)
		message.Data = packet
	case SyncTick:
		syncTick := classes.SyncTick{}
		syncTick.ParseSyncTick()
		message.Data = syncTick
	case ConsoleCmd:
		consoleCmd := classes.ConsoleCmd{}
		consoleCmd.ParseConsoleCmd(reader)
		message.Data = consoleCmd
	case UserCmd:
		userCmd := classes.UserCmd{}
		userCmd.ParseUserCmd(reader)
		message.Data = userCmd
	case DataTables:
		dataTables := classes.DataTables{}
		dataTables.ParseDataTables(reader)
		message.Data = dataTables
	case Stop:
		stop := classes.Stop{}
		stop.ParseStop(reader)
		message.Data = stop
	case CustomData: // TODO: not sar data
		customData := classes.CustomData{}
		customData.ParseCustomData(reader, message.TickNumber, uint8(message.PacketType))
		message.Data = customData
	case StringTables: // TODO: parsing string table data
		stringTables := classes.StringTables{}
		stringTables.ParseStringTables(reader)
		message.Data = stringTables
	default:
		panic("invalid packet type")
	}
	return message
}

func (t MessageType) String() string {
	switch t {
	case SignOn:
		return "SIGNON"
	case Packet:
		return "PACKET"
	case SyncTick:
		return "SYNCTICK"
	case ConsoleCmd:
		return "CONSOLECMD"
	case UserCmd:
		return "USERCMD"
	case DataTables:
		return "DATATABLES"
	case Stop:
		return "STOP"
	case CustomData:
		return "CUSTOMDATA"
	case StringTables:
		return "STRINGTABLES"
	}
	return "INVALID"
}
