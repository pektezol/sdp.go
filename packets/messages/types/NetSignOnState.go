package types

import (
	"github.com/pektezol/bitreader"
)

type NetSignOnState struct {
	SignonState       int8
	SpawnCount        uint32
	NumServerPlayers  uint32
	PlayersNetworkIds []byte
	MapNameLength     uint32
	MapName           string
}

func ParseNetSignOnState(reader *bitreader.ReaderType) NetSignOnState {
	netsignonstate := NetSignOnState{
		SignonState:      int8(reader.TryReadInt8()),
		SpawnCount:       reader.TryReadInt32(),
		NumServerPlayers: reader.TryReadInt32(),
	}
	length := reader.TryReadInt32()
	netsignonstate.PlayersNetworkIds = reader.TryReadBytesToSlice(int(length))
	netsignonstate.MapNameLength = reader.TryReadInt32()
	netsignonstate.MapName = reader.TryReadStringLen(int(netsignonstate.MapNameLength))
	return netsignonstate
}
