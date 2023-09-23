package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/classes"
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
		signOn := classes.SignOn{}
		signOn.ParseSignOn(reader)
	case 2: // Packet
		packet := classes.Packet{}
		packet.ParsePacket(reader)
	case 3: // SyncTick
		syncTick := classes.SyncTick{}
		syncTick.ParseSyncTick()
	case 4: // ConsoleCmd
		consoleCmd := classes.ConsoleCmd{}
		consoleCmd.ParseConsoleCmd(reader)
	case 5: // UserCmd
		userCmd := classes.UserCmd{}
		userCmd.ParseUserCmd(reader)
	case 6: // DataTables
		dataTables := classes.DataTables{}
		dataTables.ParseDataTables(reader)
	case 7: // Stop
		stop := classes.Stop{}
		stop.ParseStop(reader)
	case 8: // CustomData TODO: not sar data
		customData := classes.CustomData{}
		customData.ParseCustomData(reader, tickNumber, packetType)
	case 9: // StringTables TODO: parsing string table data
		stringTables := classes.StringTables{}
		stringTables.ParseStringTables(reader)
	default: // Invalid
		panic("invalid packet type")
	}
	return PacketMessageInfo{
		PacketType: packetType,
		TickNumber: tickNumber,
		SlotNumber: slotNumber,
	}
}
