package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/classes"
	"github.com/pektezol/demoparser/pkg/messages"
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
		for count := 0; count < 2; count++ {
			reader.SkipBytes(76)
		}
		reader.SkipBytes(8)
		packetReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(reader.TryReadUInt32())), true)
		for {
			messageType, err := packetReader.ReadBits(6)
			if err != nil {
				break
			}
			messages.ParseMessages(messageType, packetReader)
		}
	case 2: // Packet
		for count := 0; count < 2; count++ {
			reader.SkipBytes(76)
		}
		reader.SkipBytes(8)
		reader.TryReadBytesToSlice(uint64(reader.TryReadUInt32()))
	case 3: // SyncTick
	case 4: // ConsoleCmd

		consoleCmd := classes.ConsoleCmd{}
		consoleCmd.ParseConsoleCmd(reader)
	case 5: // UserCmd

		userCmd := classes.UserCmd{}
		userCmd.ParseUserCmd(reader)
	case 6: // DataTables
		reader.SkipBytes(uint64(reader.TryReadUInt32()))
	case 7: // Stop
		if reader.TryReadBool() {
			reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
		}
	case 8: // CustomData TODO: not sar data
		customData := classes.CustomData{}
		customData.ParseCustomData(reader, tickNumber, packetType)
	case 9: // StringTables TODO: parsing string table data
		reader.SkipBytes(uint64(reader.TryReadUInt32()))
	default: // Invalid

		panic("invalid packet type")
	}
	return PacketMessageInfo{
		PacketType: packetType,
		TickNumber: tickNumber,
		SlotNumber: slotNumber,
	}
}
