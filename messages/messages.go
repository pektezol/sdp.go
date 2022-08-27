package messages

import (
	"fmt"
	"os"
	"parser/classes"
	"parser/utils"
)

const (
	MSSC int32 = 2
)

func MessageTypeCheck(file *os.File) (statusCode int) {
	Type := make([]byte, 1)
	Tick := make([]byte, 4)
	Slot := make([]byte, 1)
	file.Read(Type)
	file.Read(Tick)
	file.Read(Slot)
	switch Type[0] {
	case 0x01: // SignOn
		var packet Packet
		var cmdinfo classes.CmdInfo
		packet.PacketInfo = utils.ReadByteFromFile(file, 76*MSSC)
		packet.InSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.OutSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Data = utils.ReadByteFromFile(file, packet.Size)
		cmdinfo = classes.CmdInfoInit(packet.PacketInfo)
		fmt.Println(cmdinfo)
		return 1
	case 0x02: // Packet
		var packet Packet
		var cmdinfo classes.CmdInfo
		packet.PacketInfo = utils.ReadByteFromFile(file, 76*MSSC)
		packet.InSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.OutSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Data = utils.ReadByteFromFile(file, packet.Size)
		cmdinfo = classes.CmdInfoInit(packet.PacketInfo)
		fmt.Printf("[%d] %v\n", utils.IntFromBytes(Tick), cmdinfo)
		return 2
	case 0x03: // SyncTick
		return 3
	case 0x04: // Consolecmd
		var consolecmd ConsoleCmd
		consolecmd.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		consolecmd.Data = string(utils.ReadByteFromFile(file, consolecmd.Size))
		fmt.Printf("[%d] %s\n", utils.IntFromBytes(Tick), consolecmd.Data)
		return 4
	case 0x05: // Usercmd FIXME: Correct bit-packing inside classes
		var usercmd UserCmd
		var usercmdinfo classes.UserCmdInfo
		usercmd.Cmd = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		usercmd.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		usercmd.Data = utils.ReadByteFromFile(file, usercmd.Size)
		usercmdinfo = classes.UserCmdInfoInit(usercmd.Data, int(usercmd.Size))
		fmt.Printf("[%d] UserCmd: %v\n", utils.IntFromBytes(Tick), usercmdinfo)
		return 5
	case 0x06: // DataTables
		var datatables DataTables
		//var stringtable classes.StringTable
		datatables.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		datatables.Data = utils.ReadByteFromFile(file, datatables.Size)
		// stringtable = classes.StringTableInit(Data)
		// fmt.Printf("[%d] DataTables: %v\n", utils.IntFromBytes(Size), stringtable)
		return 6
	case 0x07: // Stop
		fmt.Println("Stop")
		return 7
	case 0x08: // CustomData
		Unknown := make([]byte, 4)
		file.Read(Unknown)
		Size := make([]byte, 4)
		file.Read(Size)
		Data := make([]byte, utils.IntFromBytes(Size))
		file.Read(Data)
		return 8
	case 0x09: // StringTables
		Size := make([]byte, 4)
		file.Read(Size)
		Data := make([]byte, utils.IntFromBytes(Size))
		file.Read(Data)
		return 9
	default:
		return 0
	}
	//fmt.Println(Type[0])
	//fmt.Println(utils.IntFromBytes(Tick))
	//fmt.Println(Slot[0])
}
