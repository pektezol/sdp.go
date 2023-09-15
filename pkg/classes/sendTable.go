package classes

import (
	"github.com/pektezol/bitreader"
)

type SendTable struct {
	NeedsDecoder bool
	NetTableName string
	NumOfProps   int16
	Props        []prop
}

type prop struct {
	SendPropType  int8
	SendPropName  string
	SendPropFlags int32
	Priority      int8
	ExcludeDtName string
	LowValue      float32
	HighValue     float32
	NumBits       int32
	NumElements   int32
}

type sendPropFlag int

const (
	Unsigned sendPropFlag = iota
	Coord
	NoScale
	RoundDown
	RoundUp
	Normal
	Exclude
	Xyze
	InsideArray
	ProxyAlwaysYes
	IsVectorElem
	Collapsible
	CoordMp
	CoordMpLp // low precision
	CoordMpInt
	CellCoord
	CellCoordLp
	CellCoordInt
	ChangesOften
)

type sendPropType int

const (
	Int sendPropType = iota
	Float
	Vector3
	Vector2
	String
	Array
	DataTable
)

func ParseSendTable(reader *bitreader.Reader) SendTable {
	sendTable := SendTable{
		NeedsDecoder: reader.TryReadBool(),
		NetTableName: reader.TryReadString(),
		NumOfProps:   int16(reader.TryReadBits(10)),
		// SendPropType:  int8(reader.TryReadBits(5)),
		// SendPropName:  reader.TryReadString(),
		// SendPropFlags: int16(reader.TryReadBits(16)),
	}
	if sendTable.NumOfProps < 0 {
		return sendTable
	}
	for count := 0; count < int(sendTable.NumOfProps); count++ {
		propType := int8(reader.TryReadBits(5))
		if propType >= int8(7) {
			return sendTable
		}
		prop := prop{
			SendPropType:  propType,
			SendPropName:  reader.TryReadString(),
			SendPropFlags: int32(reader.TryReadBits(19)),
			Priority:      int8(reader.TryReadBits(8)),
		}
		if propType == int8(DataTable) || CheckBit(int64(prop.SendPropFlags), int(Exclude)) {
			prop.ExcludeDtName = reader.TryReadString()
		} else {
			switch propType {
			case int8(String), int8(Int), int8(Float), int8(Vector3), int8(Vector2):
				prop.LowValue = reader.TryReadFloat32()
				prop.HighValue = reader.TryReadFloat32()
				prop.NumBits = int32(reader.TryReadBits(7))
			case int8(Array):
				prop.NumElements = int32(reader.TryReadBits(10))
			default:
				return sendTable
			}
		}
		sendTable.Props = append(sendTable.Props, prop)
	}
	return sendTable
}

func CheckBit(val int64, bit int) bool {
	return (val & (int64(1) << bit)) != 0
}
