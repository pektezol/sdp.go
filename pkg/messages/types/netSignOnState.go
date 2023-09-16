package messages

import "github.com/pektezol/bitreader"

type NetSignOnState struct {
	SignOnState        int8
	SpawnCount         int32
	NumServerPlayers   int32
	IdsLength          int32
	PlayersNetworksIds []byte
	MapNameLength      int32
	MapName            string
}

func ParseNetSignOnState(reader *bitreader.Reader) NetSignOnState {
	netSignOnState := NetSignOnState{
		SignOnState:      int8(reader.TryReadBits(8)),
		SpawnCount:       int32(reader.TryReadBits(32)),
		NumServerPlayers: int32(reader.TryReadBits(32)),
		IdsLength:        int32(reader.TryReadBits(32)),
	}
	netSignOnState.PlayersNetworksIds = reader.TryReadBytesToSlice(uint64(netSignOnState.IdsLength))
	netSignOnState.MapNameLength = int32(reader.TryReadBits(32))
	netSignOnState.MapName = reader.TryReadStringLength(uint64(netSignOnState.MapNameLength))
	return netSignOnState
}
