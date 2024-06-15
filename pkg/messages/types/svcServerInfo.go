package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcServerInfo struct {
	Protocol         uint16  `json:"protocol"`
	ServerCount      uint32  `json:"server_count"`
	IsHltv           bool    `json:"is_hltv"`
	IsDedicated      bool    `json:"is_dedicated"`
	ClientCrc        int32   `json:"client_crc"`
	StringTableCrc   uint32  `json:"string_table_crc"`
	MaxServerClasses uint16  `json:"max_server_classes"`
	MapCrc           uint32  `json:"map_crc"`
	PlayerCount      uint8   `json:"player_count"`
	MaxClients       uint8   `json:"max_clients"`
	TickInterval     float32 `json:"tick_interval"`
	Platform         string  `json:"platform"`
	GameDir          string  `json:"game_dir"`
	MapName          string  `json:"map_name"`
	SkyName          string  `json:"sky_name"`
	HostName         string  `json:"host_name"`
}

func ParseSvcServerInfo(reader *bitreader.Reader, demo *types.Demo) SvcServerInfo {
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
	demo.Writer.TempAppendLine("\t\tNetwork Protocol: %d", svcServerInfo.Protocol)
	demo.Writer.TempAppendLine("\t\tServer Count: %d", svcServerInfo.ServerCount)
	demo.Writer.TempAppendLine("\t\tIs Hltv: %t", svcServerInfo.IsHltv)
	demo.Writer.TempAppendLine("\t\tIs Dedicated: %t", svcServerInfo.IsDedicated)
	demo.Writer.TempAppendLine("\t\tServer Client CRC: %d", svcServerInfo.ClientCrc)
	demo.Writer.TempAppendLine("\t\tString Table CRC: %d", svcServerInfo.StringTableCrc)
	demo.Writer.TempAppendLine("\t\tMax Server Classes: %d", svcServerInfo.MaxServerClasses)
	demo.Writer.TempAppendLine("\t\tServer Map CRC: %d", svcServerInfo.MapCrc)
	demo.Writer.TempAppendLine("\t\tCurrent Player Count: %d", svcServerInfo.PlayerCount)
	demo.Writer.TempAppendLine("\t\tMax Player Count: %d", svcServerInfo.MaxClients)
	demo.Writer.TempAppendLine("\t\tInterval Per Tick: %f", svcServerInfo.TickInterval)
	demo.Writer.TempAppendLine("\t\tPlatform: %s", svcServerInfo.Platform)
	demo.Writer.TempAppendLine("\t\tGame Directory: %s", svcServerInfo.GameDir)
	demo.Writer.TempAppendLine("\t\tMap Name: %s", svcServerInfo.MapName)
	demo.Writer.TempAppendLine("\t\tSky Name: %s", svcServerInfo.SkyName)
	demo.Writer.TempAppendLine("\t\tHost Name: %s", svcServerInfo.HostName)
	return svcServerInfo
}
