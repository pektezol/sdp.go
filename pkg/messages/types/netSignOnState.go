package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetSignOnState struct {
	SignOnState        string `json:"sign_on_state"`
	SpawnCount         int32  `json:"spawn_count"`
	NumServerPlayers   uint32 `json:"num_server_players"`
	IdsLength          uint32 `json:"ids_length"`
	PlayersNetworksIds []byte `json:"players_networks_ids"`
	MapNameLength      uint32 `json:"map_name_length"`
	MapName            string `json:"map_name"`
}

type SignOnState int

const (
	ESignOnStateNone        SignOnState = iota // no state yet, about to connect
	ESignOnStateChallenge                      // client challenging server, all OOB packets
	ESignOnStateConnected                      // client is connected to server, netchans ready
	ESignOnStateNew                            // just got server info and string tables
	ESignOnStatePreSpawn                       // received signon buggers
	ESignOnStateSpawn                          // ready to receive entity packets
	ESignOnStateFull                           // we are fully connected, first non-delta packet received
	ESignOnStateChangeLevel                    // server is changing level, please wait
)

func (signOnState SignOnState) String() string {
	switch signOnState {
	case ESignOnStateNone:
		return "None"
	case ESignOnStateChallenge:
		return "Challenge"
	case ESignOnStateConnected:
		return "Connected"
	case ESignOnStateNew:
		return "New"
	case ESignOnStatePreSpawn:
		return "PreSpawn"
	case ESignOnStateSpawn:
		return "Spawn"
	case ESignOnStateFull:
		return "Full"
	case ESignOnStateChangeLevel:
		return "ChangeLevel"
	default:
		return fmt.Sprintf("%d", int(signOnState))
	}
}

func ParseNetSignOnState(reader *bitreader.Reader, demo *types.Demo) NetSignOnState {
	netSignOnState := NetSignOnState{
		SignOnState:      SignOnState(reader.TryReadUInt8()).String(),
		SpawnCount:       reader.TryReadSInt32(),
		NumServerPlayers: reader.TryReadUInt32(),
		IdsLength:        reader.TryReadUInt32(),
	}
	demo.Writer.TempAppendLine("\t\tSign On State: %s", netSignOnState.SignOnState)
	demo.Writer.TempAppendLine("\t\tSpawn Count: %d", netSignOnState.SpawnCount)
	demo.Writer.TempAppendLine("\t\tNumber Of Server Players: %d", netSignOnState.NumServerPlayers)
	if netSignOnState.IdsLength > 0 {
		netSignOnState.PlayersNetworksIds = reader.TryReadBytesToSlice(uint64(netSignOnState.IdsLength))
		demo.Writer.TempAppendLine("\t\tPlayer Network IDs: %v", netSignOnState.PlayersNetworksIds)
	}
	netSignOnState.MapNameLength = reader.TryReadUInt32()
	if netSignOnState.MapNameLength > 0 {
		netSignOnState.MapName = reader.TryReadStringLength(uint64(netSignOnState.MapNameLength))
		demo.Writer.TempAppendLine("\t\tMap Name: %s", netSignOnState.MapName)
	}
	return netSignOnState
}
