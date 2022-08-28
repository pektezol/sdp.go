package messages

type Header struct {
	DemoFileStamp   string
	DemoProtocol    int32
	NetworkProtocol int32
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float32
	PlaybackTicks   int32
	PlaybackFrames  int32
	SignOnLength    int32
}

type Message struct {
	Type byte
	Tick int
	Slot byte
}

type Packet struct {
	PacketInfo  []byte
	InSequence  int32
	OutSequence int32
	Size        int32
	Data        []byte
}

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

type CustomData struct {
	Unknown int32
	Size    int32
	Data    []byte
}

type StringTables struct {
	Size int32
	Data []byte
}