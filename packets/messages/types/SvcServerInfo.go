package types

import "github.com/pektezol/bitreader"

type SvcServerInfo struct {
	Protocol     uint16
	ServerCount  uint32
	IsHltv       bool
	IsDedicated  bool
	ClientCrc    int32
	MaxClasses   uint16
	MapCrc       uint32
	PlayerSlot   uint8
	MaxClients   uint8
	Unk          uint32
	TickInterval float32
	COs          byte
	GameDir      string
	MapName      string
	SkyName      string
	HostName     string
}

func ParseSvcServerInfo(reader *bitreader.ReaderType) SvcServerInfo {
	return SvcServerInfo{
		Protocol:     reader.TryReadInt16(),
		ServerCount:  reader.TryReadInt32(),
		IsHltv:       reader.TryReadBool(),
		IsDedicated:  reader.TryReadBool(),
		ClientCrc:    int32(reader.TryReadInt32()),
		MaxClasses:   reader.TryReadInt16(),
		MapCrc:       reader.TryReadInt32(),
		PlayerSlot:   reader.TryReadInt8(),
		MaxClients:   reader.TryReadInt8(),
		Unk:          reader.TryReadInt32(),
		TickInterval: reader.TryReadFloat32(),
		COs:          reader.TryReadInt8(),
		GameDir:      reader.TryReadString(),
		MapName:      reader.TryReadString(),
		SkyName:      reader.TryReadString(),
		HostName:     reader.TryReadString(),
	}
}
