package types

import (
	"github.com/pektezol/sdp.go/pkg/writer"
)

type Demo struct {
	Headers       Headers           `json:"headers"`
	Messages      []Message         `json:"messages"`
	Writer        *writer.Writer    `json:"-"`
	GameEventList *SvcGameEventList `json:"-"`
}

type Headers struct {
	DemoFileStamp   string  `json:"demo_file_stamp"`
	DemoProtocol    int32   `json:"demo_protocol"`
	NetworkProtocol int32   `json:"network_protocol"`
	ServerName      string  `json:"server_name"`
	ClientName      string  `json:"client_name"`
	MapName         string  `json:"map_name"`
	GameDirectory   string  `json:"game_directory"`
	PlaybackTime    float32 `json:"playback_time"`
	PlaybackTicks   int32   `json:"playback_ticks"`
	PlaybackFrames  int32   `json:"playback_frames"`
	SignOnLength    int32   `json:"sign_on_length"`
}

type Message struct {
	PacketType MessageType `json:"packet_type"`
	TickNumber int32       `json:"tick_number"`
	SlotNumber uint8       `json:"slot_number"`
	Data       any         `json:"data"`
}

type MessageType uint8

const (
	SignOn MessageType = iota + 1
	Packet
	SyncTick
	ConsoleCmd
	UserCmd
	DataTables
	Stop
	CustomData
	StringTables
)

func (t MessageType) String() string {
	switch t {
	case SignOn:
		return "SIGNON"
	case Packet:
		return "PACKET"
	case SyncTick:
		return "SYNCTICK"
	case ConsoleCmd:
		return "CONSOLECMD"
	case UserCmd:
		return "USERCMD"
	case DataTables:
		return "DATATABLES"
	case Stop:
		return "STOP"
	case CustomData:
		return "CUSTOMDATA"
	case StringTables:
		return "STRINGTABLES"
	}
	return "INVALID"
}
