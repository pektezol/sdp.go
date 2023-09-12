package classes

import (
	"bytes"

	"github.com/pektezol/bitreader"
)

type StringTable struct {
	Name         string
	TableEntries []StringTableEntry
	Classes      []StringTableClass
}

type StringTableEntry struct {
	Name      string
	EntryData StringTableEntryData
}

type StringTableEntryData struct {
	// TODO: Parse StringTableEntry
}

type StringTableClass struct {
	Name string
	Data string
}

func ParseStringTables(reader *bitreader.ReaderType) []StringTable {
	tableCount := reader.TryReadBits(8)
	stringTables := make([]StringTable, tableCount)
	for i := 0; i < int(tableCount); i++ {
		var table StringTable
		table.ParseStream(reader)
		stringTables[i] = table
	}
	return stringTables
}

func (stringTable *StringTable) ParseStream(reader *bitreader.ReaderType) {
	stringTable.Name = reader.TryReadString()
	entryCount := reader.TryReadBits(16)
	stringTable.TableEntries = make([]StringTableEntry, entryCount)

	for i := 0; i < int(entryCount); i++ {
		var entry StringTableEntry
		entry.Parse(reader)
		stringTable.TableEntries[i] = entry
	}

	if reader.TryReadBool() {
		classCount := reader.TryReadBits(16)
		stringTable.Classes = make([]StringTableClass, classCount)

		for i := 0; i < int(classCount); i++ {
			var class StringTableClass
			class.Parse(reader)
			stringTable.Classes[i] = class
		}
	}
}

func (stringTableEntry *StringTableEntry) Parse(reader *bitreader.ReaderType) {
	stringTableEntry.Name = reader.TryReadString()
	if reader.TryReadBool() {
		byteLen, err := reader.ReadBits(16)
		if err != nil {
			return
		}
		dataBsr := reader.TryReadBytesToSlice(int(byteLen))
		_ = bitreader.Reader(bytes.NewReader(dataBsr), true) // TODO: Parse StringTableEntry
		// stringTableEntry.EntryData.ParseStream(entryReader)
	}
}

func (stringTableClass *StringTableClass) Parse(reader *bitreader.ReaderType) {
	stringTableClass.Name = reader.TryReadString()
	if reader.TryReadBool() {
		dataLen := reader.TryReadBits(16)
		stringTableClass.Data = reader.TryReadStringLen(int(dataLen))
	}
}
