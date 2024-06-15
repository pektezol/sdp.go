package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcGameEvent struct {
	Length           uint16                    `json:"length"`
	EventID          uint32                    `json:"event_id"`
	EventDescription types.GameEventDescriptor `json:"event_description"`
	EventDescriptors []EventDescriptorKey      `json:"event_descriptors"`
}

type EventDescriptorKey struct {
	Name       string `json:"name"`
	Descriptor any    `json:"descriptor"`
}

func ParseSvcGameEvent(reader *bitreader.Reader, demo *types.Demo) SvcGameEvent {
	svcGameEvent := SvcGameEvent{
		Length: uint16(reader.TryReadBits(11)),
	}
	gameEventReader := bitreader.NewReaderFromBytes(reader.TryReadBitsToSlice(uint64(svcGameEvent.Length)), true)
	svcGameEvent.parseGameEvent(gameEventReader, demo)
	return svcGameEvent
}

func (svcGameEvent *SvcGameEvent) parseGameEvent(reader *bitreader.Reader, demo *types.Demo) {
	svcGameEvent.EventID = uint32(reader.TryReadBits(9))
	svcGameEvent.EventDescription = demo.GameEventList.GameEventDescriptor[svcGameEvent.EventID]
	demo.Writer.TempAppendLine("\t\t%s (%d):", svcGameEvent.EventDescription.Name, svcGameEvent.EventID)
	for _, descriptor := range svcGameEvent.EventDescription.Keys {
		var value any
		switch descriptor.Type {
		case types.EventDescriptorString:
			value = reader.TryReadString()
		case types.EventDescriptorFloat:
			value = reader.TryReadFloat32()
		case types.EventDescriptorInt32:
			value = reader.TryReadSInt32()
		case types.EventDescriptorInt16:
			value = reader.TryReadSInt16()
		case types.EventDescriptorInt8:
			value = reader.TryReadUInt8()
		case types.EventDescriptorBool:
			value = reader.TryReadBool()
		case types.EventDescriptorUInt64:
			value = reader.TryReadUInt64()
		}
		svcGameEvent.EventDescriptors = append(svcGameEvent.EventDescriptors, EventDescriptorKey{
			Name:       descriptor.Name,
			Descriptor: value,
		})
		demo.Writer.TempAppendLine("\t\t\t%s: %v", descriptor.Name, value)
	}
}
