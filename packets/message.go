package packets

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/packets/classes"
)

const MSSC = 2

func ParseMessage(reader *bitreader.ReaderType) (status int) {
	messageType := reader.TryReadInt8()
	messageTick := reader.TryReadInt32()
	messageSlot := reader.TryReadInt8()
	switch messageType {
	case 0x01: // TODO: SignOn - Data
		signOn := SignOn{
			PacketInfo:  classes.ParseCmdInfo(reader, MSSC),
			InSequence:  int32(reader.TryReadInt32()),
			OutSequence: int32(reader.TryReadInt32()),
			Size:        int32(reader.TryReadInt32()),
		}
		reader.SkipBytes(int(signOn.Size))
		fmt.Printf("[%d] (%d) {%d} SignOn: %v\n", messageTick, messageType, messageSlot, signOn)
		return 1
	case 0x02: // TODO: Packet - Data
		packet := Packet{
			PacketInfo:  classes.ParseCmdInfo(reader, MSSC),
			InSequence:  int32(reader.TryReadInt32()),
			OutSequence: int32(reader.TryReadInt32()),
			Size:        int32(reader.TryReadInt32()),
		}
		reader.SkipBytes(int(packet.Size))
		//fmt.Printf("[%d] (%d) Packet: %v\n", messageTick, messageType, packet)
		return 2
	case 0x03:
		syncTick := SyncTick{}
		fmt.Printf("[%d] (%d) SyncTick: %v\n", messageTick, messageType, syncTick)
		return 3
	case 0x04:
		consoleCmd := ConsoleCmd{
			Size: int32(reader.TryReadInt32()),
		}
		consoleCmd.Data = reader.TryReadStringLen(int(consoleCmd.Size))
		//fmt.Printf("[%d] (%d) ConsoleCmd: %s\n", messageTick, messageType, consoleCmd.Data)
		return 4
	case 0x05: // TODO: UserCmd - Buttons
		userCmd := UserCmd{
			Cmd:  int32(reader.TryReadInt32()),
			Size: int32(reader.TryReadInt32()),
		}
		userCmd.Data = classes.ParseUserCmdInfo(reader.TryReadBytesToSlice(int(userCmd.Size)))
		// fmt.Printf("[%d] (%d) UserCmd: %v\n", messageTick, messageType, userCmd)
		return 5
	case 0x06: // TODO: DataTables
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) DataTables: \n", messageTick, messageType)
		return 6
	case 0x07:
		stop := Stop{
			RemainingData: nil,
		}
		fmt.Printf("[%d] (%d) Stop: %v\n", messageTick, messageType, stop)
		return 7
	case 0x08: // TODO: CustomData
		reader.SkipBytes(4)
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) CustomData: \n", messageTick, messageType)
		return 8
	case 0x09: // TODO: StringTables - Data
		stringTables := StringTables{
			Size: int32(reader.TryReadInt32()),
		}
		stringTables.Data = classes.ParseStringTable(reader.TryReadBytesToSlice(int(stringTables.Size)))
		// fmt.Printf("[%d] (%d) StringTables: %v\n", messageTick, messageType, stringTables)
		return 9
	default:
		return 0
	}
}
