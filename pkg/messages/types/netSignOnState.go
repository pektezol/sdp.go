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
	None        SignOnState = iota // no state yet, about to connect
	Challenge                      // client challenging server, all OOB packets
	Connected                      // client is connected to server, netchans ready
	New                            // just got server info and string tables
	PreSpawn                       // received signon buggers
	Spawn                          // ready to receive entity packets
	Full                           // we are fully connected, first non-delta packet received
	ChangeLevel                    // server is changing level, please wait
)

func (signOnState SignOnState) String() string {
	switch signOnState {
	case None:
		return "None"
	case Challenge:
		return "Challenge"
	case Connected:
		return "Connected"
	case New:
		return "New"
	case PreSpawn:
		return "PreSpawn"
	case Spawn:
		return "Spawn"
	case Full:
		return "Full"
	case ChangeLevel:
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
