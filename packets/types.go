package packets

import (
	"github.com/pektezol/demoparser/packets/classes"
	"github.com/pektezol/demoparser/packets/messages"
)

type Header struct {
	DemoFileStamp   string
	DemoProtocol    uint32
	NetworkProtocol uint32
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float32
	PlaybackTicks   int32
	PlaybackFrames  int32
	SignOnLength    uint32
}

type SignOn struct {
	PacketInfo  []classes.CmdInfo
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []messages.Message
}

type Packet struct {
	PacketInfo  []classes.CmdInfo
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []messages.Message
}

type SyncTick struct{}

type ConsoleCmd struct {
	Size int32
	Data string
}

type UserCmd struct {
	Cmd  int32
	Size int32
	Data classes.UserCmdInfo
}

type DataTables struct {
	Size int32
	Data classes.DataTable
}

type Stop struct {
	RemainingData []byte
}

type CustomData struct {
	Unknown int32
	Size    int32
	Data    []byte
}

type StringTables struct {
	Size int32
	Data []classes.StringTable
}
