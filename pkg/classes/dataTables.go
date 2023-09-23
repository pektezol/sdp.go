package classes

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type DataTables struct {
	Size            int32
	SendTable       []SendTable
	ServerClassInfo []ServerClassInfo
}

type SendTable struct {
	NeedsDecoder bool
	NetTableName string
	NumOfProps   int16
	Props        []prop
}

type ServerClassInfo struct {
	ClassId       uint16
	ClassName     string
	DataTableName string
}

type prop struct {
	SendPropType  sendPropType
	SendPropName  string
	SendPropFlags uint32
	Priority      uint8
	ExcludeDtName string
	LowValue      float32
	HighValue     float32
	NumBits       int32
	NumElements   int32
}

func (dataTables *DataTables) ParseDataTables(reader *bitreader.Reader) {
	dataTables.Size = int32(reader.TryReadSInt32())
	dataTableReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(dataTables.Size)), true)
	count := 0
	for dataTableReader.TryReadBool() {
		count++
		dataTables.SendTable = append(dataTables.SendTable, ParseSendTable(dataTableReader))
	}
	numOfClasses := dataTableReader.TryReadBits(16)
	for count = 0; count < int(numOfClasses); count++ {
		dataTables.ServerClassInfo = append(dataTables.ServerClassInfo, ParseServerClassInfo(dataTableReader, count, int(numOfClasses)))
	}
}

func ParseSendTable(reader *bitreader.Reader) SendTable {
	sendTable := SendTable{
		NeedsDecoder: reader.TryReadBool(),
		NetTableName: reader.TryReadString(),
		NumOfProps:   int16(reader.TryReadBits(10)),
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
			SendPropType:  sendPropType(propType),
			SendPropName:  reader.TryReadString(),
			SendPropFlags: uint32(reader.TryReadBits(19)),
			Priority:      reader.TryReadUInt8(),
		}
		if propType == int8(ESendPropTypeDataTable) || checkBit(prop.SendPropFlags, 6) {
			prop.ExcludeDtName = reader.TryReadString()
		} else {
			switch propType {
			case int8(ESendPropTypeString), int8(ESendPropTypeInt), int8(ESendPropTypeFloat), int8(ESendPropTypeVector3), int8(ESendPropTypeVector2):
				prop.LowValue = reader.TryReadFloat32()
				prop.HighValue = reader.TryReadFloat32()
				prop.NumBits = int32(reader.TryReadBits(7))
			case int8(ESendPropTypeArray):
				prop.NumElements = int32(reader.TryReadBits(10))
			default:
				return sendTable
			}
		}
		sendTable.Props = append(sendTable.Props, prop)
	}
	return sendTable
}

func ParseServerClassInfo(reader *bitreader.Reader, count int, numOfClasses int) ServerClassInfo {
	serverClassInfo := ServerClassInfo{
		ClassId:       reader.TryReadUInt16(),
		ClassName:     reader.TryReadString(),
		DataTableName: reader.TryReadString(),
	}
	return serverClassInfo
}

// func serverClassBits(numOfClasses int) int {
// 	return highestBitIndex(uint(numOfClasses)) + 1
// }

// func highestBitIndex(i uint) int {
// 	var j int
// 	for j = 31; j >= 0 && (i&(1<<j)) == 0; j-- {
// 	}
// 	return j
// }

func checkBit(val uint32, bit int) bool {
	return (val & (uint32(1) << bit)) != 0
}

type sendPropType int

const (
	ESendPropTypeInt sendPropType = iota
	ESendPropTypeFloat
	ESendPropTypeVector3
	ESendPropTypeVector2
	ESendPropTypeString
	ESendPropTypeArray
	ESendPropTypeDataTable
)

const (
	ESendPropFlagUnsigned       string = "Unsigned"
	ESendPropFlagCoord          string = "Coord"
	ESendPropFlagNoScale        string = "NoScale"
	ESendPropFlagRoundDown      string = "RoundDown"
	ESendPropFlagRoundUp        string = "RoundUp"
	ESendPropFlagNormal         string = "Normal"
	ESendPropFlagExclude        string = "Exclude"
	ESendPropFlagXyze           string = "Xyze"
	ESendPropFlagInsideArray    string = "InsideArray"
	ESendPropFlagProxyAlwaysYes string = "ProxyAlwaysYes"
	ESendPropFlagIsVectorElem   string = "IsVectorElem"
	ESendPropFlagCollapsible    string = "Collapsible"
	ESendPropFlagCoordMp        string = "CoordMp"
	ESendPropFlagCoordMpLp      string = "CoordMpLp"
	ESendPropFlagCoordMpInt     string = "CoordMpInt"
	ESendPropFlagCellCoord      string = "CellCoord"
	ESendPropFlagCellCoordLp    string = "CellCoordLp"
	ESendPropFlagCellCoordInt   string = "CellCoordInt"
	ESendPropFlagChangesOften   string = "ChangesOften"
)

func (prop prop) GetFlags() []string {
	flags := []string{}
	if checkBit(prop.SendPropFlags, 0) {
		flags = append(flags, ESendPropFlagUnsigned)
	}
	if checkBit(prop.SendPropFlags, 1) {
		flags = append(flags, ESendPropFlagCoord)
	}
	if checkBit(prop.SendPropFlags, 2) {
		flags = append(flags, ESendPropFlagNoScale)
	}
	if checkBit(prop.SendPropFlags, 3) {
		flags = append(flags, ESendPropFlagRoundDown)
	}
	if checkBit(prop.SendPropFlags, 4) {
		flags = append(flags, ESendPropFlagRoundUp)
	}
	if checkBit(prop.SendPropFlags, 5) {
		flags = append(flags, ESendPropFlagNormal)
	}
	if checkBit(prop.SendPropFlags, 6) {
		flags = append(flags, ESendPropFlagExclude)
	}
	if checkBit(prop.SendPropFlags, 7) {
		flags = append(flags, ESendPropFlagXyze)
	}
	if checkBit(prop.SendPropFlags, 8) {
		flags = append(flags, ESendPropFlagInsideArray)
	}
	if checkBit(prop.SendPropFlags, 9) {
		flags = append(flags, ESendPropFlagProxyAlwaysYes)
	}
	if checkBit(prop.SendPropFlags, 10) {
		flags = append(flags, ESendPropFlagIsVectorElem)
	}
	if checkBit(prop.SendPropFlags, 11) {
		flags = append(flags, ESendPropFlagCollapsible)
	}
	if checkBit(prop.SendPropFlags, 12) {
		flags = append(flags, ESendPropFlagCoordMp)
	}
	if checkBit(prop.SendPropFlags, 13) {
		flags = append(flags, ESendPropFlagCoordMpLp)
	}
	if checkBit(prop.SendPropFlags, 14) {
		flags = append(flags, ESendPropFlagCoordMpInt)
	}
	if checkBit(prop.SendPropFlags, 15) {
		flags = append(flags, ESendPropFlagCellCoord)
	}
	if checkBit(prop.SendPropFlags, 16) {
		flags = append(flags, ESendPropFlagCellCoordLp)
	}
	if checkBit(prop.SendPropFlags, 17) {
		flags = append(flags, ESendPropFlagCellCoordInt)
	}
	if checkBit(prop.SendPropFlags, 18) {
		flags = append(flags, ESendPropFlagChangesOften)
	}
	return flags
}

func (sendPropType sendPropType) String() string {
	switch sendPropType {
	case ESendPropTypeInt:
		return "Int"
	case ESendPropTypeFloat:
		return "Float"
	case ESendPropTypeVector3:
		return "Vector3"
	case ESendPropTypeVector2:
		return "Vector2"
	case ESendPropTypeString:
		return "String"
	case ESendPropTypeArray:
		return "Array"
	case ESendPropTypeDataTable:
		return "DataTable"
	default:
		return fmt.Sprintf("%d", int(sendPropType))
	}
}
