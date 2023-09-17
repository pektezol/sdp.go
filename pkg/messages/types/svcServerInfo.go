package messages

import "github.com/pektezol/bitreader"

type SvcServerInfo struct {
	Protocol         uint16
	ServerCount      uint32
	IsHltv           bool
	IsDedicated      bool
	ClientCrc        int32
	StringTableCrc   uint32
	MaxServerClasses uint16
	MapCrc           uint32
	PlayerCount      uint8
	MaxClients       uint8
	TickInterval     float32
	Platform         string
	GameDir          string
	MapName          string
	SkyName          string
	HostName         string
}

func ParseSvcServerInfo(reader *bitreader.Reader) SvcServerInfo {
	return SvcServerInfo{
		Protocol:         reader.TryReadUInt16(),
		ServerCount:      reader.TryReadUInt32(),
		IsHltv:           reader.TryReadBool(),
		IsDedicated:      reader.TryReadBool(),
		ClientCrc:        reader.TryReadSInt32(),
		StringTableCrc:   reader.TryReadUInt32(),
		MaxServerClasses: reader.TryReadUInt16(),
		MapCrc:           reader.TryReadUInt32(),
		PlayerCount:      reader.TryReadUInt8(),
		MaxClients:       reader.TryReadUInt8(),
		TickInterval:     reader.TryReadFloat32(),
		Platform:         reader.TryReadStringLength(1),
		GameDir:          reader.TryReadString(),
		MapName:          reader.TryReadString(),
		SkyName:          reader.TryReadString(),
		HostName:         reader.TryReadString(),
	}
}
