package packets

type Header struct {
	DemoFileStamp   string
	DemoProtocol    uint
	NetworkProtocol uint
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float32
	PlaybackTicks   int
	PlaybackFrames  int
	SignOnLength    uint
}

type SignOn struct {
	PacketInfo  []byte
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []byte
}

type Packet struct {
	PacketInfo  []byte
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []byte
}

type SyncTick struct{}

type ConsoleCmd struct {
	Size int32
	Data string
}

type UserCmd struct {
	Cmd  int32
	Size int32
	Data []byte
}

type DataTables struct {
	Size int32
	Data []byte
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
	Data []byte
}
