package classes

import (
	"fmt"
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type StringTables struct {
	Size int32
	Data []StringTable
}

type StringTable struct {
	Name         string
	TableEntries []StringTableEntry
	Classes      []StringTableClass
}

type StringTableClass struct {
	Name string
	Data string
}
type StringTableEntry struct {
	Name      string
	EntryData any
}

func (stringTables *StringTables) ParseStringTables(reader *bitreader.Reader) {
	stringTables.Size = reader.TryReadSInt32()
	stringTableReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(stringTables.Size)), true)
	tableCount := stringTableReader.TryReadBits(8)
	tables := make([]StringTable, tableCount)
	for i := 0; i < int(tableCount); i++ {
		var table StringTable
		table.ParseStream(stringTableReader)
		tables[i] = table
	}
	stringTables.Data = tables
}

func (stringTable *StringTable) ParseStream(reader *bitreader.Reader) {
	stringTable.Name = reader.TryReadString()
	entryCount := reader.TryReadBits(16)
	writer.AppendLine("\tTable Name: %s", stringTable.Name)
	stringTable.TableEntries = make([]StringTableEntry, entryCount)

	for i := 0; i < int(entryCount); i++ {
		var entry StringTableEntry
		entry.Parse(stringTable.Name, reader)
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

func (stringTableClass *StringTableClass) Parse(reader *bitreader.Reader) {
	stringTableClass.Name = reader.TryReadString()
	writer.TempAppendLine("\t\t\tName: %s", stringTableClass.Name)
	if reader.TryReadBool() {
		stringTableClass.Data = reader.TryReadStringLength(uint64(reader.TryReadUInt16()))
		writer.TempAppendLine("\t\t\tData: %s", stringTableClass.Data)
	}
}

func (stringTableEntry *StringTableEntry) Parse(tableName string, reader *bitreader.Reader) {
	stringTableEntry.Name = reader.TryReadString()
	writer.TempAppendLine("\t\t\tName: %s", stringTableEntry.Name)
	if reader.TryReadBool() {
		byteLen, err := reader.ReadBits(16)
		if err != nil {
			return
		}
		stringTableEntryReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(byteLen), true)
		switch tableName {
		case StringTableUserInfo:
			stringTableEntry.ParseUserInfo(stringTableEntryReader)
		case StringTableServerQueryInfo:
			stringTableEntry.ParseServerQueryInfo(stringTableEntryReader)
		case StringTableGameRulesCreation:
			stringTableEntry.ParseGamesRulesCreation(stringTableEntryReader)
		case StringTableInfoPanel:
			stringTableEntry.ParseInfoPanel(stringTableEntryReader)
		case StringTableLightStyles:
			stringTableEntry.ParseLightStyles(stringTableEntryReader)
		case StringTableModelPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader)
		case StringTableGenericPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader)
		case StringTableSoundPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader)
		case StringTableDecalPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader)
		default:
			stringTableEntry.ParseUnknown(stringTableEntryReader)
		}
	}
}

func (stringTableEntry *StringTableEntry) ParseUserInfo(reader *bitreader.Reader) {
	const SignedGuidLen int32 = 32
	const MaxPlayerNameLength int32 = 32
	userInfo := struct {
		SteamID         uint64
		Name            string
		UserID          int32
		GUID            string
		FriendsID       uint32
		FriendsName     string
		FakePlayer      bool
		IsHltv          bool
		CustomFiles     []uint32
		FilesDownloaded uint8
	}{
		SteamID: reader.TryReadUInt64(),
		Name:    reader.TryReadStringLength(uint64(MaxPlayerNameLength)),
		UserID:  reader.TryReadSInt32(),
		GUID:    reader.TryReadStringLength(uint64(SignedGuidLen) + 1),
	}
	reader.SkipBytes(3)
	userInfo.FriendsID = reader.TryReadUInt32()
	userInfo.FriendsName = reader.TryReadStringLength(uint64(MaxPlayerNameLength))
	userInfo.FakePlayer = reader.TryReadUInt8() != 0
	userInfo.IsHltv = reader.TryReadUInt8() != 0
	reader.SkipBytes(2)
	userInfo.CustomFiles = []uint32{reader.TryReadUInt32(), reader.TryReadUInt32(), reader.TryReadUInt32(), reader.TryReadUInt32()}
	userInfo.FilesDownloaded = reader.TryReadUInt8()
	reader.SkipBytes(3)
	stringTableEntry.EntryData = userInfo
	writer.TempAppendLine("\t\t\t\tSteam Account ID: %d", uint32((userInfo.SteamID&0xFFFFFFFF00000000)|userInfo.SteamID))
	writer.TempAppendLine("\t\t\t\tSteam Account Instance: %d", uint32(userInfo.SteamID>>32)&0x000FFFFF)
	writer.TempAppendLine("\t\t\t\tSteam Account Type: %d", uint32(userInfo.SteamID>>52)&0xF)
	writer.TempAppendLine("\t\t\t\tSteam Account Universe: %d", uint32(userInfo.SteamID>>56)&0xFF)
	writer.TempAppendLine("\t\t\t\tName: %s", userInfo.Name)
	writer.TempAppendLine("\t\t\t\tUser ID: %d", userInfo.UserID)
	writer.TempAppendLine("\t\t\t\tGUID: %s", userInfo.GUID)
	writer.TempAppendLine("\t\t\t\tFriends ID: %d", userInfo.FriendsID)
	writer.TempAppendLine("\t\t\t\tFriends Name: %s", userInfo.FriendsName)
	writer.TempAppendLine("\t\t\t\tFake Player: %t", userInfo.FakePlayer)
	writer.TempAppendLine("\t\t\t\tIs Htlv: %t", userInfo.IsHltv)
	if userInfo.CustomFiles != nil {
		writer.TempAppendLine("\t\t\t\tCustom File CRCs: [logo: 0x%d, sounds: 0x%d, models: 0x%d, txt: 0x%d]", userInfo.CustomFiles[0], userInfo.CustomFiles[1], userInfo.CustomFiles[2], userInfo.CustomFiles[3])
	}
	writer.TempAppendLine("\t\t\t\tFiles Downloaded: %d", userInfo.FilesDownloaded)
}

func (stringTableEntry *StringTableEntry) ParseServerQueryInfo(reader *bitreader.Reader) {
	serverQueryInfo := struct{ Port uint32 }{
		Port: reader.TryReadUInt32(),
	}
	stringTableEntry.EntryData = serverQueryInfo
	writer.TempAppendLine("\t\t\t\tPort: %d", serverQueryInfo.Port)
}

func (stringTableEntry *StringTableEntry) ParseGamesRulesCreation(reader *bitreader.Reader) {
	gamesRulesCreation := struct{ Message string }{
		Message: reader.TryReadString(),
	}
	stringTableEntry.EntryData = gamesRulesCreation
	writer.TempAppendLine("\t\t\t\tMessage: %s", gamesRulesCreation.Message)
}

func (stringTableEntry *StringTableEntry) ParseInfoPanel(reader *bitreader.Reader) {
	infoPanel := struct{ Message string }{
		Message: reader.TryReadString(),
	}
	stringTableEntry.EntryData = infoPanel
	writer.TempAppendLine("\t\t\t\tMessage: %s", infoPanel.Message)
}

func (stringTableEntry *StringTableEntry) ParseLightStyles(reader *bitreader.Reader) {
	lightStyles := struct{ Values []byte }{}
	str := reader.TryReadString()
	if len(str) != 0 {
		for _, c := range str {
			value := byte((c - 'a') * 22)
			lightStyles.Values = append(lightStyles.Values, value)
		}
	}
	stringTableEntry.EntryData = lightStyles
	if lightStyles.Values == nil {
		writer.TempAppendLine("\t\t\t\t0 Frames (256)")
	} else {
		writer.TempAppendLine("\t\t\t\t%d frames: %v", len(lightStyles.Values), lightStyles.Values)
	}
}

func (stringTableEntry *StringTableEntry) ParsePrecacheData(reader *bitreader.Reader) {
	type PrecacheFlag uint16
	const (
		None           PrecacheFlag = 0
		FatalIfMissing PrecacheFlag = 1
		Preload        PrecacheFlag = 1 << 1
	)
	precacheData := struct{ Flags uint8 }{
		Flags: uint8(reader.TryReadBits(2)),
	}
	getFlags := func(flags PrecacheFlag) []string {
		var flagStrings []string
		if flags&FatalIfMissing != 0 {
			flagStrings = append(flagStrings, "FatalIfMissing")
		}
		if flags&Preload != 0 {
			flagStrings = append(flagStrings, "Preload")
		}
		return flagStrings
	}
	writer.TempAppendLine("\t\t\t\tFlags: %v", getFlags(PrecacheFlag(precacheData.Flags)))
}

func (stringTableEntry *StringTableEntry) ParseUnknown(reader *bitreader.Reader) {
	unknown := reader.TryReadBitsToSlice(reader.TryReadRemainingBits())
	binaryString := ""
	for _, byteValue := range unknown {
		binaryString += fmt.Sprintf("%08b ", byteValue)
	}
	writer.TempAppendLine("\t\t\t\tUnknown: (%s)", strings.TrimSpace(binaryString))
}

const (
	StringTableDownloadables       string = "downloadables"
	StringTableModelPreCache       string = "modelprecache"
	StringTableGenericPreCache     string = "genericprecache"
	StringTableSoundPreCache       string = "soundprecache"
	StringTableDecalPreCache       string = "decalprecache"
	StringTableInstanceBaseLine    string = "instancebaseline"
	StringTableLightStyles         string = "lightstyles"
	StringTableUserInfo            string = "userinfo"
	StringTableServerQueryInfo     string = "server_query_info"
	StringTableParticleEffectNames string = "ParticleEffectNames"
	StringTableEffectDispatch      string = "EffectDispatch"
	StringTableVguiScreen          string = "VguiScreen"
	StringTableMaterials           string = "Materials"
	StringTableInfoPanel           string = "InfoPanel"
	StringTableScenes              string = "Scenes"
	StringTableMeleeWeapons        string = "MeleeWeapons"
	StringTableGameRulesCreation   string = "GameRulesCreation"
	StringTableBlackMarket         string = "BlackMarketTable"
	// custom?
	StringTableDynamicModels  string = "DynamicModels"
	StringTableServerMapCycle string = "ServerMapCycle"
)
