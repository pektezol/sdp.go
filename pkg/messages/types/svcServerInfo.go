package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

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
	svcServerInfo := SvcServerInfo{
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
	writer.TempAppendLine("\t\tNetwork Protocol: %d", svcServerInfo.Protocol)
	writer.TempAppendLine("\t\tServer Count: %d", svcServerInfo.ServerCount)
	writer.TempAppendLine("\t\tIs Hltv: %t", svcServerInfo.IsHltv)
	writer.TempAppendLine("\t\tIs Dedicated: %t", svcServerInfo.IsDedicated)
	writer.TempAppendLine("\t\tServer Client CRC: %d", svcServerInfo.ClientCrc)
	writer.TempAppendLine("\t\tString Table CRC: %d", svcServerInfo.StringTableCrc)
	writer.TempAppendLine("\t\tMax Server Classes: %d", svcServerInfo.MaxServerClasses)
	writer.TempAppendLine("\t\tServer Map CRC: %d", svcServerInfo.MapCrc)
	writer.TempAppendLine("\t\tCurrent Player Count: %d", svcServerInfo.PlayerCount)
	writer.TempAppendLine("\t\tMax Player Count: %d", svcServerInfo.MaxClients)
	writer.TempAppendLine("\t\tInterval Per Tick: %f", svcServerInfo.TickInterval)
	writer.TempAppendLine("\t\tPlatform: %s", svcServerInfo.Platform)
	writer.TempAppendLine("\t\tGame Directory: %s", svcServerInfo.GameDir)
	writer.TempAppendLine("\t\tMap Name: %s", svcServerInfo.MapName)
	writer.TempAppendLine("\t\tSky Name: %s", svcServerInfo.SkyName)
	writer.TempAppendLine("\t\tHost Name: %s", svcServerInfo.HostName)
	return svcServerInfo
}
