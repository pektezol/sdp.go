package messages

import (
	"log"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcGameEvent struct {
	Length           uint16
	EventID          uint32
	EventDescription GameEventDescriptor
	EventDescriptors []EventDescriptorKey
}

type EventDescriptorKey struct {
	Name       string
	Descriptor any
}

func ParseSvcGameEvent(reader *bitreader.Reader) SvcGameEvent {
	svcGameEvent := SvcGameEvent{
		Length: uint16(reader.TryReadBits(11)),
	}
	gameEventReader := bitreader.NewReaderFromBytes(reader.TryReadBitsToSlice(uint64(svcGameEvent.Length)), true)
	svcGameEvent.parseGameEvent(gameEventReader)
	return svcGameEvent
}

func (svcGameEvent *SvcGameEvent) parseGameEvent(reader *bitreader.Reader) {
	svcGameEvent.EventID = uint32(reader.TryReadBits(9))
	log.Println(GameEventList.GameEventDescriptor)
	svcGameEvent.EventDescription = GameEventList.GameEventDescriptor[svcGameEvent.EventID]
	writer.TempAppendLine("\t\t%s (%d):", svcGameEvent.EventDescription.Name, svcGameEvent.EventID)
	for _, descriptor := range svcGameEvent.EventDescription.Keys {
		var value any
		switch descriptor.Type {
		case EventDescriptorString:
			value = reader.TryReadString()
		case EventDescriptorFloat:
			value = reader.TryReadFloat32()
		case EventDescriptorInt32:
			value = reader.TryReadSInt32()
		case EventDescriptorInt16:
			value = reader.TryReadSInt16()
		case EventDescriptorInt8:
			value = reader.TryReadUInt8()
		case EventDescriptorBool:
			value = reader.TryReadBool()
		case EventDescriptorUInt64:
			value = reader.TryReadUInt64()
		}
		svcGameEvent.EventDescriptors = append(svcGameEvent.EventDescriptors, EventDescriptorKey{
			Name:       descriptor.Name,
			Descriptor: value,
		})
		writer.TempAppendLine("\t\t\t%s: %v", descriptor.Name, value)
	}
}
