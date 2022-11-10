package classes

import (
	"bytes"
	"fmt"

	"github.com/pektezol/bitreader"
)

type DataTable struct {
	SendTable       []SendTable
	ServerClassInfo []ServerClassInfo
}

type SendTable struct {
	NeedsDecoder  bool
	NetTableName  string
	NumOfProps    uint16
	SendPropType  int8
	SendPropName  string
	SendPropFlags int16
}

type ServerClassInfo struct {
	ClassId       int16
	ClassName     string
	DataTableName string
}

func ParseDataTable(data []byte) DataTable {
	reader := bitreader.Reader(bytes.NewReader(data), true)
	sendtable := parseSendTable(reader)
	fmt.Println("AAA")
	fmt.Println(reader.TryReadBits(8))
	serverclassinfo := parseServerClassInfo(reader)
	return DataTable{
		SendTable:       sendtable,
		ServerClassInfo: serverclassinfo,
	}
}

func parseSendTable(reader *bitreader.ReaderType) []SendTable {
	var sendtables []SendTable
	for reader.TryReadBool() {
		sendtables = append(sendtables, SendTable{
			NeedsDecoder:  reader.TryReadBool(),
			NetTableName:  reader.TryReadString(),
			NumOfProps:    uint16(reader.TryReadBits(10)),
			SendPropType:  int8(reader.TryReadBits(5)),
			SendPropName:  reader.TryReadString(),
			SendPropFlags: int16(reader.TryReadInt16()),
		})
	}
	return sendtables
}
func parseServerClassInfo(reader *bitreader.ReaderType) []ServerClassInfo {
	var serverclassinfo []ServerClassInfo
	numofclasses := reader.TryReadInt16()
	fmt.Println(numofclasses)
	for i := 0; i < int(numofclasses); i++ {
		serverclassinfo = append(serverclassinfo, ServerClassInfo{
			ClassId:       int16(reader.TryReadInt16()),
			ClassName:     reader.TryReadString(),
			DataTableName: reader.TryReadString(),
		})
	}
	return serverclassinfo
}
