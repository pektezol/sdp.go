package types

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type SvcGameEventList struct {
	Events              int16                 `json:"events"`
	Length              int32                 `json:"length"`
	GameEventDescriptor []GameEventDescriptor `json:"game_event_descriptor"`
}

type GameEventDescriptor struct {
	EventID uint32                   `json:"event_id"`
	Name    string                   `json:"name"`
	Keys    []GameEventDescriptorKey `json:"keys"`
}

type GameEventDescriptorKey struct {
	Name string          `json:"name"`
	Type EventDescriptor `json:"type"`
}

type EventDescriptor uint8

func (svcGameEventList *SvcGameEventList) ParseGameEventDescriptor(reader *bitreader.Reader, demo *Demo) {
	svcGameEventList.GameEventDescriptor = make([]GameEventDescriptor, svcGameEventList.Events)
	for event := 0; event < int(svcGameEventList.Events); event++ {
		svcGameEventList.GameEventDescriptor[event] = GameEventDescriptor{
			EventID: uint32(reader.TryReadBits(9)),
			Name:    reader.TryReadString(),
		}
		demo.Writer.TempAppendLine("\t\t\t%d: %s", svcGameEventList.GameEventDescriptor[event].EventID, svcGameEventList.GameEventDescriptor[event].Name)
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
		demo.Writer.TempAppendLine("\t\t\t\tKeys: %v", svcGameEventList.GameEventDescriptor[event].Keys)
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
