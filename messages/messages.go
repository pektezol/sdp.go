package messages

import (
	"fmt"
	"os"
	"parser/utils"
)

const MSSC int32 = 2

func ParseMessage(file *os.File) (statusCode int) {
	var message Message
	message.Type = utils.ReadByteFromFile(file, 1)[0]
	message.Tick = int(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	message.Slot = utils.ReadByteFromFile(file, 1)[0]
	switch message.Type {
	case 0x01: // SignOn
		var packet Packet
		// var cmdinfo classes.CmdInfo
		packet.PacketInfo = utils.ReadByteFromFile(file, 76*MSSC)
		packet.InSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.OutSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Data = utils.ReadByteFromFile(file, packet.Size)
		// cmdinfo = classes.CmdInfoInit(packet.PacketInfo)
		// fmt.Println(cmdinfo)
		return 1
	case 0x02: // Packet
		var packet Packet
		// var cmdinfo classes.CmdInfo
		packet.PacketInfo = utils.ReadByteFromFile(file, 76*MSSC)
		packet.InSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.OutSequence = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		packet.Data = utils.ReadByteFromFile(file, packet.Size)
		// cmdinfo = classes.CmdInfoInit(packet.PacketInfo)
		// fmt.Printf("[%d] %v\n", utils.IntFromBytes(Tick), cmdinfo)
		return 2
	case 0x03: // SyncTick
		return 3
	case 0x04: // Consolecmd
		var consolecmd ConsoleCmd
		consolecmd.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		consolecmd.Data = string(utils.ReadByteFromFile(file, consolecmd.Size))
		fmt.Printf("[%d] %s\n", message.Tick, consolecmd.Data)
		return 4
	case 0x05: // Usercmd FIXME: Correct bit-packing inside classes
		var usercmd UserCmd
		// var usercmdinfo classes.UserCmdInfo
		usercmd.Cmd = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		usercmd.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		usercmd.Data = utils.ReadByteFromFile(file, usercmd.Size)
		// usercmdinfo = classes.UserCmdInfoInit(usercmd.Data, int(usercmd.Size))
		// fmt.Printf("[%d] UserCmd: %v\n", utils.IntFromBytes(Tick), usercmdinfo)
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
		fmt.Println("Stop - End of Demo")
		return 7
	case 0x08: // CustomData
		var customdata CustomData
		customdata.Unknown = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		customdata.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		customdata.Data = utils.ReadByteFromFile(file, customdata.Size)
		return 8
	case 0x09: // StringTables
		var stringtables StringTables
		stringtables.Size = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
		stringtables.Data = utils.ReadByteFromFile(file, stringtables.Size)
		return 9
	default:
		return 0
	}

}

func ParseHeader(file *os.File) {
	var header Header
	header.DemoFileStamp = string(utils.ReadByteFromFile(file, 8))
	header.DemoProtocol = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	header.NetworkProtocol = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	header.ServerName = string(utils.ReadByteFromFile(file, 260))
	header.ClientName = string(utils.ReadByteFromFile(file, 260))
	header.MapName = string(utils.ReadByteFromFile(file, 260))
	header.GameDirectory = string(utils.ReadByteFromFile(file, 260))
	header.PlaybackTime = float32(utils.FloatFromBytes(utils.ReadByteFromFile(file, 4)))
	header.PlaybackTicks = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	header.PlaybackFrames = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	header.SignOnLength = int32(utils.IntFromBytes(utils.ReadByteFromFile(file, 4)))
	fmt.Println(header)
}
