package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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

func ParseStringTables(reader *bitreader.Reader) []StringTable {
	tableCount := reader.TryReadBits(8)
	stringTables := make([]StringTable, tableCount)
	for i := 0; i < int(tableCount); i++ {
		var table StringTable
		table.ParseStream(reader)
		stringTables[i] = table
	}
	return stringTables
}

func (stringTable *StringTable) ParseStream(reader *bitreader.Reader) {
	stringTable.Name = reader.TryReadString()
	entryCount := reader.TryReadBits(16)
	writer.AppendLine("\tTable Name: %s", stringTable.Name)
	stringTable.TableEntries = make([]StringTableEntry, entryCount)

	for i := 0; i < int(entryCount); i++ {
		var entry StringTableEntry
		entry.Parse(reader)
		stringTable.TableEntries[i] = entry
	}
	if entryCount != 0 {
		writer.AppendLine("\t\t%d Table Entries:", entryCount)
		writer.AppendOutputFromTemp()
	} else {
		writer.AppendLine("\t\tNo Table Entries")
	}
	if reader.TryReadBool() {
		classCount := reader.TryReadBits(16)
		stringTable.Classes = make([]StringTableClass, classCount)

		for i := 0; i < int(classCount); i++ {
			var class StringTableClass
			class.Parse(reader)
			stringTable.Classes[i] = class
		}
		writer.AppendLine("\t\t%d Classes:", classCount)
		writer.AppendOutputFromTemp()
	} else {
		writer.AppendLine("\t\tNo Class Entries")
	}
}

func (stringTableEntry *StringTableEntry) Parse(reader *bitreader.Reader) {
	stringTableEntry.Name = reader.TryReadString()
	if reader.TryReadBool() {
		byteLen, err := reader.ReadBits(16)
		if err != nil {
			return
		}
		dataBsr := reader.TryReadBytesToSlice(byteLen)
		_ = bitreader.NewReaderFromBytes(dataBsr, true) // TODO: Parse StringTableEntry
		// stringTableEntry.EntryData.ParseStream(entryReader)
	}
}

func (stringTableClass *StringTableClass) Parse(reader *bitreader.Reader) {
	stringTableClass.Name = reader.TryReadString()
	writer.TempAppendLine("\t\t\tName: %s", stringTableClass.Name)
	if reader.TryReadBool() {
		dataLen := reader.TryReadBits(16)
		stringTableClass.Data = reader.TryReadStringLength(dataLen)
		writer.TempAppendLine("\t\t\tData: %s", stringTableClass.Data)
	}
}
