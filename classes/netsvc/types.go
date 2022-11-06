package netsvc

type NetDisconnect struct {
	Text string
}

type NetFile struct {
	TransferId    int32
	FileName      string
	FileRequested bool
}

type NetSplitScreenUser struct {
	Unknown bool
}

type NetTick struct {
	Tick                      int32
	HostFrameTime             int16
	HostFrameTimeStdDeviation int16
}

type NetStringCmd struct {
	Command string
}

type ConVar struct {
	Name  string
	Value string
}

type NetSetConVar struct {
	Length  int8
	ConVars []ConVar
}

type NetSignonStateOE struct {
	SignonState int8
	SpawnCount  int32
}

type NetSignonStateNE struct {
	NetSignonStateOE
	NumServerPlayers int32
	IdsLength        int32
	PlayerNetworkIds []byte
	MapNameLength    int32
	MapName          string
}

type SvcServerInfo struct {
	Protocol     int8
	ServerCount  int32
	IsHltv       bool
	IsDedicated  bool
	ClientCrc    int32
	MaxClasses   int16
	MapCrc       int32
	PlayerSlot   int8
	MaxClients   int8
	Unk          int32 // NE
	TickInterval float32
	COs          byte
	GameDir      string
	MapName      string
	SkyName      string
	HostName     string
}

type SvcSendTable struct {
	NeedsDecoder bool
	Length       int8
	Props        int // ?
}

type ServerClass struct {
	ClassId       int8
	ClassName     string
	DataTableName string
}

type SvcClassInfo struct {
	Length         int16
	CreateOnClient bool
	ServerClasses  []ServerClass
}

type SvcSetPause struct {
	Paused bool
}

type SvcCreateStringTable struct {
	Name              string
	MaxEntries        int16
	NumEntries        int8
	Length            int32
	UserDataFixedSize bool
	UserDataSize      int32
	UserDataSizeBits  int8
	Flags             int8
	StringData        int // ?
}

type SvcPrint struct {
	Message string
}
