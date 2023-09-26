package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcGameEventList struct {
	Events              int16
	Length              int32
	GameEventDescriptor []GameEventDescriptor
}

type GameEventDescriptor struct {
	EventID uint32
	Name    string
	Keys    []struct {
		Name string
		Type EventDescriptor
	}
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
	return svcGameEventList
}

func (svcGameEventList *SvcGameEventList) parseGameEventDescriptor(reader *bitreader.Reader) {
	svcGameEventList.GameEventDescriptor = make([]GameEventDescriptor, svcGameEventList.Events)
	for event := 0; event < int(svcGameEventList.Events); event++ {
		gameEventDescriptor := GameEventDescriptor{
			EventID: uint32(reader.TryReadBits(9)),
			Name:    reader.TryReadString(),
		}
		writer.TempAppendLine("\t\t\t%d: %s", gameEventDescriptor.EventID, gameEventDescriptor.Name)
		for {
			descriptorType := reader.TryReadBits(3)
			if descriptorType == 0 {
				break
			}
			KeyName := reader.TryReadString()
			gameEventDescriptor.Keys = append(gameEventDescriptor.Keys, struct {
				Name string
				Type EventDescriptor
			}{
				Name: KeyName,
				Type: EventDescriptor(descriptorType),
			})
		}
		writer.TempAppendLine("\t\t\t\tKeys: %v", gameEventDescriptor.Keys)
	}
}

const (
	String EventDescriptor = iota + 1
	Float
	Int32
	Int16
	Int8
	Bool
	UInt64
)

func (eventDescriptor EventDescriptor) String() string {
	switch eventDescriptor {
	case String:
		return "String"
	case Float:
		return "Float"
	case Int32:
		return "Int32"
	case Int16:
		return "Int16"
	case Int8:
		return "Int8"
	case Bool:
		return "Bool"
	case UInt64:
		return "UInt64"
	default:
		return fmt.Sprintf("%d", eventDescriptor)
	}
}
