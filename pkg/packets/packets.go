package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/classes"
	"github.com/pektezol/demoparser/pkg/writer"
)

type PacketMessageInfo struct {
	PacketType uint8
	TickNumber int32
	SlotNumber uint8
}

func ParsePackets(reader *bitreader.Reader) PacketMessageInfo {
	packetType := reader.TryReadUInt8()
	tickNumber := reader.TryReadSInt32()
	slotNumber := reader.TryReadUInt8()
	switch packetType {
	case 1: // SignOn
		writer.AppendLine("[%d] %s (%d):", tickNumber, "SIGNON", packetType)
		signOn := classes.SignOn{}
		signOn.ParseSignOn(reader)
	case 2: // Packet
		writer.AppendLine("[%d] %s (%d):", tickNumber, "PACKET", packetType)
		packet := classes.Packet{}
		packet.ParsePacket(reader)
	case 3: // SyncTick
		writer.AppendLine("[%d] %s (%d):", tickNumber, "SYNCTICK", packetType)
		syncTick := classes.SyncTick{}
		syncTick.ParseSyncTick()
	case 4: // ConsoleCmd
		writer.AppendLine("[%d] %s (%d):", tickNumber, "CONSOLECMD", packetType)
		consoleCmd := classes.ConsoleCmd{}
		consoleCmd.ParseConsoleCmd(reader)
	case 5: // UserCmd
		writer.AppendLine("[%d] %s (%d):", tickNumber, "USERCMD", packetType)
		userCmd := classes.UserCmd{}
		userCmd.ParseUserCmd(reader)
	case 6: // DataTables
		writer.AppendLine("[%d] %s (%d):", tickNumber, "DATATABLES", packetType)
		dataTables := classes.DataTables{}
		dataTables.ParseDataTables(reader)
	case 7: // Stop
		writer.AppendLine("[%d] %s (%d):", tickNumber, "STOP", packetType)
		stop := classes.Stop{}
		stop.ParseStop(reader)
	case 8: // CustomData TODO: not sar data
		customData := classes.CustomData{}
		customData.ParseCustomData(reader, tickNumber, packetType)
	case 9: // StringTables TODO: parsing string table data
		writer.AppendLine("[%d] %s (%d):", tickNumber, "STRINGTABLES", packetType)
		stringTables := classes.StringTables{}
		stringTables.ParseStringTables(reader)
	default: // Invalid
		writer.AppendLine("[%d] %s (%d):", tickNumber, "INVALID", packetType)
		panic("invalid packet type")
	}
	return PacketMessageInfo{
		PacketType: packetType,
		TickNumber: tickNumber,
		SlotNumber: slotNumber,
	}
}
