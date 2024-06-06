package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

var GameEventList *SvcGameEventList

type SvcGameEventList struct {
	Events              int16
	Length              int32
	GameEventDescriptor []GameEventDescriptor
}

type GameEventDescriptor struct {
	EventID uint32
	Name    string
	Keys    []GameEventDescriptorKey
}

type GameEventDescriptorKey struct {
	Name string
	Type EventDescriptor
}

type EventDescriptor uint8

func ParseSvcGameEventList(reader *bitreader.Reader) SvcGameEventList {
	svcGameEventList := SvcGameEventList{
		Events: int16(reader.TryReadBits(9)),
		Length: int32(reader.TryReadBits(20)),
	}
	gameEventListReader := bitreader.NewReaderFromBytes(reader.TryReadBitsToSlice(uint64(svcGameEventList.Length)), true)
	writer.TempAppendLine("\t\t%d Events:", svcGameEventList.Events)
	svcGameEventList.parseGameEventDescriptor(gameEventListReader)
	GameEventList = &svcGameEventList
	return svcGameEventList
}

func (svcGameEventList *SvcGameEventList) parseGameEventDescriptor(reader *bitreader.Reader) {
	svcGameEventList.GameEventDescriptor = make([]GameEventDescriptor, svcGameEventList.Events)
	for event := 0; event < int(svcGameEventList.Events); event++ {
		svcGameEventList.GameEventDescriptor[event] = GameEventDescriptor{
			EventID: uint32(reader.TryReadBits(9)),
			Name:    reader.TryReadString(),
		}
		writer.TempAppendLine("\t\t\t%d: %s", svcGameEventList.GameEventDescriptor[event].EventID, svcGameEventList.GameEventDescriptor[event].Name)
		for {
			descriptorType := reader.TryReadBits(3)
			if descriptorType == 0 {
				break
			}
			KeyName := reader.TryReadString()
			svcGameEventList.GameEventDescriptor[event].Keys = append(svcGameEventList.GameEventDescriptor[event].Keys, GameEventDescriptorKey{
				Name: KeyName,
				Type: EventDescriptor(descriptorType),
			})
		}
		writer.TempAppendLine("\t\t\t\tKeys: %v", svcGameEventList.GameEventDescriptor[event].Keys)
	}
}

const (
	EventDescriptorString EventDescriptor = iota + 1
	EventDescriptorFloat
	EventDescriptorInt32
	EventDescriptorInt16
	EventDescriptorInt8
	EventDescriptorBool
	EventDescriptorUInt64
)

func (eventDescriptor EventDescriptor) String() string {
	switch eventDescriptor {
	case EventDescriptorString:
		return "String"
	case EventDescriptorFloat:
		return "Float"
	case EventDescriptorInt32:
		return "Int32"
	case EventDescriptorInt16:
		return "Int16"
	case EventDescriptorInt8:
		return "Int8"
	case EventDescriptorBool:
		return "Bool"
	case EventDescriptorUInt64:
		return "UInt64"
	default:
		return fmt.Sprintf("%d", eventDescriptor)
	}
}
