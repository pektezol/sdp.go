package packets

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseMessage(reader *bitreader.ReaderType) (status int) {
	messageType := reader.TryReadInt8()
	messageTick := reader.TryReadInt32()
	messageSlot := reader.TryReadInt8()
	//fmt.Println(messageType, messageTick, messageSlot)
	switch messageType {
	case 0x01:
		//signOn := SignOn{}
		reader.SkipBytes(76*2 + 8)
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		fmt.Printf("[%d] (%d) {%d} SignOn: \n", messageTick, messageType, messageSlot)
		return 1
	case 0x02:
		reader.SkipBytes(76*2 + 8)
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) Packet: \n", messageTick, messageType)
		return 2
	case 0x03:
		fmt.Printf("[%d] (%d) SyncTick: \n", messageTick, messageType)
		return 3
	case 0x04:
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) ConsoleCmd: \n", messageTick, messageType)
		return 4
	case 0x05:
		reader.SkipBytes(4)
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) UserCmd: \n", messageTick, messageType)
		return 5
	case 0x06:
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) DataTables: \n", messageTick, messageType)
		return 6
	case 0x07:
		fmt.Printf("[%d] (%d) Stop: \n", messageTick, messageType)
		return 7
	case 0x08:
		reader.SkipBytes(4)
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) CustomData: \n", messageTick, messageType)
		return 8
	case 0x09:
		val := reader.TryReadInt32()
		reader.SkipBytes(int(val))
		// fmt.Printf("[%d] (%d) StringTables: \n", messageTick, messageType)
		return 9
	default:
		return 0
	}
}
