package packets

import "github.com/pektezol/demoparser/pkg/classes"

type SignOn struct {
	PacketInfo  []classes.CmdInfo
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []any
}

type Packet struct {
	PacketInfo  []classes.CmdInfo
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []any
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
	Size            int32
	SendTable       []classes.SendTable
	ServerClassInfo []classes.ServerClassInfo
}

type Stop struct {
	RemainingData []byte
}

type CustomData struct {
	Type int32
	Size int32
	Data string
}

type StringTables struct {
	Size int32
	Data []classes.StringTable
}
