package messages

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
