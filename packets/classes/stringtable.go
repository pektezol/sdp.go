package classes

import (
	"bytes"

	"github.com/pektezol/bitreader"
)

type StringTable struct {
	TableName          string
	NumOfEntries       int16
	EntryName          string
	EntrySize          int16
	EntryData          []byte
	NumOfClientEntries int16
	ClientEntryName    string
	ClientEntrySize    int16
	ClientEntryData    []byte
}

func ParseStringTable(data []byte) []StringTable {
	reader := bitreader.Reader(bytes.NewReader(data), true)
	var stringTables []StringTable
	numOfTables := reader.TryReadInt8()
	for i := 0; i < int(numOfTables); i++ {
		var stringTable StringTable
		stringTable.TableName = reader.TryReadString()
		stringTable.NumOfEntries = int16(reader.TryReadInt16())
		stringTable.EntryName = reader.TryReadString()
		if reader.TryReadBool() {
			stringTable.EntrySize = int16(reader.TryReadInt16())
		}
		if reader.TryReadBool() {
			stringTable.EntryData = reader.TryReadBytesToSlice(int(stringTable.EntrySize))
		}
		if reader.TryReadBool() {
			stringTable.NumOfClientEntries = int16(reader.TryReadInt16())
		}
		if reader.TryReadBool() {
			stringTable.ClientEntryName = reader.TryReadString()
		}
		if reader.TryReadBool() {
			stringTable.ClientEntrySize = int16(reader.TryReadInt16())
		}
		if reader.TryReadBool() {
			stringTable.ClientEntryData = reader.TryReadBytesToSlice(int(stringTable.ClientEntrySize))
		}
		stringTables = append(stringTables, stringTable)
	}
	return stringTables
}
