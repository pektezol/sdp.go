package classes

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
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
	Props        []SendTableProp
}

type ServerClassInfo struct {
	DataTableID   uint16
	ClassName     string
	DataTableName string
}

type SendTableProp struct {
	SendPropType  SendPropType
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
	writer.AppendLine("\t%d Send Tables:", count)
	writer.AppendOutputFromTemp()
	numOfClasses := dataTableReader.TryReadBits(16)
	for count = 0; count < int(numOfClasses); count++ {
		dataTables.ServerClassInfo = append(dataTables.ServerClassInfo, ParseServerClassInfo(dataTableReader, count, int(numOfClasses)))
	}
	writer.AppendLine("\t%d Classes:", count)
	writer.AppendOutputFromTemp()
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
	writer.TempAppendLine("\t\t%s (%d Props):", sendTable.NetTableName, sendTable.NumOfProps)
	for count := 0; count < int(sendTable.NumOfProps); count++ {
		propType := int8(reader.TryReadBits(5))
		if propType >= int8(7) {
			return sendTable
		}
		prop := SendTableProp{
			SendPropType:  SendPropType(propType),
			SendPropName:  reader.TryReadString(),
			SendPropFlags: uint32(reader.TryReadBits(19)),
			Priority:      reader.TryReadUInt8(),
		}
		writer.TempAppend("\t\t\t%s\t", prop.SendPropType)
		if propType == int8(ESendPropTypeDataTable) || checkBit(prop.SendPropFlags, 6) {
			prop.ExcludeDtName = reader.TryReadString()
			writer.TempAppend(":\t%s\t", prop.ExcludeDtName)
		} else {
			switch propType {
			case int8(ESendPropTypeString), int8(ESendPropTypeInt), int8(ESendPropTypeFloat), int8(ESendPropTypeVector3), int8(ESendPropTypeVector2):
				prop.LowValue = reader.TryReadFloat32()
				prop.HighValue = reader.TryReadFloat32()
				prop.NumBits = int32(reader.TryReadBits(7))
				writer.TempAppend("Low: %f\tHigh: %f\t%d bits\t", prop.LowValue, prop.HighValue, prop.NumBits)
			case int8(ESendPropTypeArray):
				prop.NumElements = int32(reader.TryReadBits(10))
				writer.TempAppend("Elements: %d\t", prop.NumElements)
			default:
				writer.TempAppend("Unknown Prop Type: %v\t", propType)
				return sendTable
			}
		}
		writer.TempAppend("Flags: %v\tPriority: %d\n", prop.GetFlags(), prop.Priority)
		sendTable.Props = append(sendTable.Props, prop)
	}
	return sendTable
}

func ParseServerClassInfo(reader *bitreader.Reader, count int, numOfClasses int) ServerClassInfo {
	serverClassInfo := ServerClassInfo{
		DataTableID:   reader.TryReadUInt16(),
		ClassName:     reader.TryReadString(),
		DataTableName: reader.TryReadString(),
	}
	writer.TempAppendLine("\t\t\t[%d] %s (%s)", serverClassInfo.DataTableID, serverClassInfo.ClassName, serverClassInfo.DataTableName)
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

type SendPropType int

const (
	ESendPropTypeInt SendPropType = iota
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

func (prop SendTableProp) GetFlags() []string {
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

func (sendPropType SendPropType) String() string {
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
