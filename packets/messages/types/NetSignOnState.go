package types

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type NetSignOnState struct {
	SignonState       int8
	SpawnCount        uint32
	NumServerPlayers  uint32
	IdsLength         uint32
	PlayersNetworkIds []byte
	MapNameLength     uint32
	MapName           string
}

func ParseNetSignOnState(reader *bitreader.ReaderType) NetSignOnState {
	netsignonstate := NetSignOnState{
		SignonState:      int8(reader.TryReadInt8()),
		SpawnCount:       reader.TryReadInt32(),
		NumServerPlayers: reader.TryReadInt32(),
		IdsLength:        reader.TryReadInt32(),
	}
	fmt.Println(netsignonstate.IdsLength)
	netsignonstate.PlayersNetworkIds = reader.TryReadBytesToSlice(int(netsignonstate.IdsLength))
	netsignonstate.MapNameLength = reader.TryReadInt32()
	netsignonstate.MapName = reader.TryReadStringLen(int(netsignonstate.MapNameLength))
	return netsignonstate
}
