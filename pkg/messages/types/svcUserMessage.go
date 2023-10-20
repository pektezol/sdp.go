package messages

import (
	"fmt"
	"math"

	"github.com/pektezol/bitreader"
)

type SvcUserMessage struct {
	Type   int8
	Length int16
	Data   any
}

func ParseSvcUserMessage(reader *bitreader.Reader) SvcUserMessage {
	svcUserMessage := SvcUserMessage{
		Type:   int8(reader.TryReadBits(8)),
		Length: int16(reader.TryReadBits(12)),
	}
	svcUserMessage.Data = reader.TryReadBitsToSlice(uint64(svcUserMessage.Length))
	userMessageReader := bitreader.NewReaderFromBytes(svcUserMessage.Data.([]byte), true)
	switch svcUserMessage.Type {
	case 60:
		svcUserMessage.parseScoreboardTempUpdate(userMessageReader)
	}
	return svcUserMessage
}

func (svcUserMessage *SvcUserMessage) parseScoreboardTempUpdate(reader *bitreader.Reader) {
	scoreboardTempUpdate := struct {
		NumPortals int32
		TimeTaken  int32
	}{
		NumPortals: reader.TryReadSInt32(),
		TimeTaken:  reader.TryReadSInt32(),
	}
	svcUserMessage.Data = scoreboardTempUpdate
	fmt.Printf("Portal Count: %d\n", scoreboardTempUpdate.NumPortals)
	fmt.Printf("CM Ticks: %d\n", int(math.Round(float64((float32(scoreboardTempUpdate.TimeTaken)/100.0)/float32(1.0/60.0)))))
}
