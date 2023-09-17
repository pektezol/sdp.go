package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type NetSignOnState struct {
	SignOnState        string
	SpawnCount         int32
	NumServerPlayers   int32
	IdsLength          int32
	PlayersNetworksIds []byte
	MapNameLength      int32
	MapName            string
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

func ParseNetSignOnState(reader *bitreader.Reader) NetSignOnState {
	netSignOnState := NetSignOnState{
		SignOnState:      SignOnState(reader.TryReadBits(8)).String(),
		SpawnCount:       int32(reader.TryReadBits(32)),
		NumServerPlayers: int32(reader.TryReadBits(32)),
		IdsLength:        int32(reader.TryReadBits(32)),
	}
	netSignOnState.PlayersNetworksIds = reader.TryReadBytesToSlice(uint64(netSignOnState.IdsLength))
	netSignOnState.MapNameLength = int32(reader.TryReadBits(32))
	netSignOnState.MapName = reader.TryReadStringLength(uint64(netSignOnState.MapNameLength))
	return netSignOnState
}
