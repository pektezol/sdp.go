package messages

import "github.com/pektezol/bitreader"

type SvcServerInfo struct {
	Protocol     int16
	ServerCount  int32
	IsHltv       bool
	IsDedicated  bool
	ClientCrc    int32
	MaxClasses   int16
	MapCrc       int32
	PlayerSlot   int8
	MaxClients   int8
	Unk          int32
	TickInterval int32
	COs          int8
	GameDir      string
	MapName      string
	SkyName      string
	HostName     string
}

func ParseSvcServerInfo(reader *bitreader.Reader) SvcServerInfo {
	return SvcServerInfo{
		Protocol:     int16(reader.TryReadBits(16)),
		ServerCount:  int32(reader.TryReadBits(32)),
		IsHltv:       reader.TryReadBool(),
		IsDedicated:  reader.TryReadBool(),
		ClientCrc:    int32(reader.TryReadBits(32)),
		MaxClasses:   int16(reader.TryReadBits(16)),
		MapCrc:       int32(reader.TryReadBits(32)),
		PlayerSlot:   int8(reader.TryReadBits(8)),
		MaxClients:   int8(reader.TryReadBits(8)),
		Unk:          int32(reader.TryReadBits(32)),
		TickInterval: int32(reader.TryReadBits(32)),
		COs:          int8(reader.TryReadBits(8)),
		GameDir:      reader.TryReadString(),
		MapName:      reader.TryReadString(),
		SkyName:      reader.TryReadString(),
		HostName:     reader.TryReadString(),
	}
}
