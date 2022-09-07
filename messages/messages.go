package messages

import (
	"fmt"
	"os"

	"github.com/bisaxa/bitreader"
	"github.com/bisaxa/demoparser/classes"
	"github.com/bisaxa/demoparser/utils"
)

func ParseMessage(file *os.File) (statusCode int) {
	reader := bitreader.Reader(file, true)
	messageType := reader.TryReadInt8()
	messageTick := reader.TryReadInt32()
	messageSlot := reader.TryReadInt8()
	//fmt.Println(messageType, messageTick, messageSlot)
	switch messageType {
	case 0x01: // SignOn
		var packet Packet
		packet.PacketInfo = classes.ParseCmdInfo(file, 2)
		packet.InSequence = int32(reader.TryReadInt32())
		packet.OutSequence = int32(reader.TryReadInt32())
		packet.Size = int32(reader.TryReadInt32())
		reader.SkipBytes(int(packet.Size)) // TODO: NET/SVC Message Parsing
		fmt.Printf("[%d] (%d) SignOn: %v\n", messageTick, messageSlot, packet)
		return 1
	case 0x02: // Packet
		var packet Packet
		packet.PacketInfo = classes.ParseCmdInfo(file, 2)
		packet.InSequence = int32(reader.TryReadInt32())
		packet.OutSequence = int32(reader.TryReadInt32())
		packet.Size = int32(reader.TryReadInt32())
		reader.SkipBytes(int(packet.Size)) // TODO: NET/SVC Message Parsing
		fmt.Printf("[%d] Packet: %v\n", messageTick, packet)
		return 2
	case 0x03: // SyncTick
		return 3
	case 0x04: // ConsoleCmd
		var consolecmd ConsoleCmd
		consolecmd.Size = int32(reader.TryReadInt32())
		consolecmd.Data = string(utils.ReadByteFromFile(file, consolecmd.Size))
		fmt.Printf("[%d] ConsoleCmd: %s\n", messageTick, consolecmd.Data)
		return 4
	case 0x05: // UserCmd
		var usercmd UserCmd
		usercmd.Cmd = int32(reader.TryReadInt32())
		usercmd.Size = int32(reader.TryReadInt32())
		usercmd.Data = classes.ParseUserCmdInfo(file, int(usercmd.Size))
		fmt.Printf("[%d] UserCmd: %v\n", messageTick, usercmd.Data)
		return 5
	case 0x06: // DataTables
		var datatables DataTables
		datatables.Size = int32(reader.TryReadInt32())
		reader.SkipBytes(int(datatables.Size)) // TODO: DataTables Data
		return 6
	case 0x07: // Stop
		fmt.Printf("[%d] Stop\n", messageTick)
		return 7
	case 0x08: // CustomData
		var customdata CustomData
		customdata.Unknown = int32(reader.TryReadInt32())
		customdata.Size = int32(reader.TryReadInt32())
		reader.SkipBytes(int(customdata.Size)) // TODO: CustomData Data
		return 8
	case 0x09: // StringTables
		var stringtables StringTables
		stringtables.Size = int32(reader.TryReadInt32())
		reader.SkipBytes(int(stringtables.Size)) // TODO: StringTables Data
		return 9
	default:
		return 0
	}
}
