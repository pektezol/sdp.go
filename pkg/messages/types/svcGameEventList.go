package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

func ParseSvcGameEventList(reader *bitreader.Reader, demo *types.Demo) types.SvcGameEventList {
	svcGameEventList := types.SvcGameEventList{
		Events: int16(reader.TryReadBits(9)),
		Length: int32(reader.TryReadBits(20)),
	}
	gameEventListReader := bitreader.NewReaderFromBytes(reader.TryReadBitsToSlice(uint64(svcGameEventList.Length)), true)
	demo.Writer.TempAppendLine("\t\t%d Events:", svcGameEventList.Events)
	svcGameEventList.ParseGameEventDescriptor(gameEventListReader, demo)
	demo.GameEventList = &svcGameEventList
	return svcGameEventList
}
