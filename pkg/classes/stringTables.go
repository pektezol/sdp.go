package classes

import (
	"fmt"
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type StringTables struct {
	Size int32         `json:"size"`
	Data []StringTable `json:"data"`
}

type StringTable struct {
	Name         string             `json:"name"`
	TableEntries []StringTableEntry `json:"table_entries"`
	Classes      []StringTableClass `json:"classes"`
}

type StringTableClass struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
type StringTableEntry struct {
	Name      string `json:"name"`
	EntryData any    `json:"entry_data"`
}

func (stringTables *StringTables) ParseStringTables(reader *bitreader.Reader, demo *types.Demo) {
	stringTables.Size = reader.TryReadSInt32()
	stringTableReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(stringTables.Size)), true)
	tableCount := stringTableReader.TryReadBits(8)
	tables := make([]StringTable, tableCount)
	for i := 0; i < int(tableCount); i++ {
		var table StringTable
		table.ParseStream(stringTableReader, demo)
		tables[i] = table
	}
	stringTables.Data = tables
}

func (stringTable *StringTable) ParseStream(reader *bitreader.Reader, demo *types.Demo) {
	stringTable.Name = reader.TryReadString()
	entryCount := reader.TryReadBits(16)
	demo.Writer.AppendLine("\tTable Name: %s", stringTable.Name)
	stringTable.TableEntries = make([]StringTableEntry, entryCount)

	for i := 0; i < int(entryCount); i++ {
		var entry StringTableEntry
		entry.Parse(stringTable.Name, reader, demo)
		stringTable.TableEntries[i] = entry
	}
	if entryCount != 0 {
		demo.Writer.AppendLine("\t\t%d Table Entries:", entryCount)
		demo.Writer.AppendOutputFromTemp()
	} else {
		demo.Writer.AppendLine("\t\tNo Table Entries")
	}
	if reader.TryReadBool() {
		classCount := reader.TryReadBits(16)
		stringTable.Classes = make([]StringTableClass, classCount)

		for i := 0; i < int(classCount); i++ {
			var class StringTableClass
			class.Parse(reader, demo)
			stringTable.Classes[i] = class
		}
		demo.Writer.AppendLine("\t\t%d Classes:", classCount)
		demo.Writer.AppendOutputFromTemp()
	} else {
		demo.Writer.AppendLine("\t\tNo Class Entries")
	}
}

func (stringTableClass *StringTableClass) Parse(reader *bitreader.Reader, demo *types.Demo) {
	stringTableClass.Name = reader.TryReadString()
	demo.Writer.TempAppendLine("\t\t\tName: %s", stringTableClass.Name)
	if reader.TryReadBool() {
		stringTableClass.Data = reader.TryReadStringLength(uint64(reader.TryReadUInt16()))
		demo.Writer.TempAppendLine("\t\t\tData: %s", stringTableClass.Data)
	}
}

func (stringTableEntry *StringTableEntry) Parse(tableName string, reader *bitreader.Reader, demo *types.Demo) {
	stringTableEntry.Name = reader.TryReadString()
	demo.Writer.TempAppendLine("\t\t\tName: %s", stringTableEntry.Name)
	if reader.TryReadBool() {
		byteLen, err := reader.ReadBits(16)
		if err != nil {
			return
		}
		stringTableEntryReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(byteLen), true)
		switch tableName {
		case StringTableUserInfo:
			stringTableEntry.ParseUserInfo(stringTableEntryReader, demo)
		case StringTableServerQueryInfo:
			stringTableEntry.ParseServerQueryInfo(stringTableEntryReader, demo)
		case StringTableGameRulesCreation:
			stringTableEntry.ParseGamesRulesCreation(stringTableEntryReader, demo)
		case StringTableInfoPanel:
			stringTableEntry.ParseInfoPanel(stringTableEntryReader, demo)
		case StringTableLightStyles:
			stringTableEntry.ParseLightStyles(stringTableEntryReader, demo)
		case StringTableModelPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader, demo)
		case StringTableGenericPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader, demo)
		case StringTableSoundPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader, demo)
		case StringTableDecalPreCache:
			stringTableEntry.ParsePrecacheData(stringTableEntryReader, demo)
		default:
			stringTableEntry.ParseUnknown(stringTableEntryReader, demo)
		}
	}
}

func (stringTableEntry *StringTableEntry) ParseUserInfo(reader *bitreader.Reader, demo *types.Demo) {
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
	demo.Writer.TempAppendLine("\t\t\t\tSteam Account ID: %d", uint32((userInfo.SteamID&0xFFFFFFFF00000000)|userInfo.SteamID))
	demo.Writer.TempAppendLine("\t\t\t\tSteam Account Instance: %d", uint32(userInfo.SteamID>>32)&0x000FFFFF)
	demo.Writer.TempAppendLine("\t\t\t\tSteam Account Type: %d", uint32(userInfo.SteamID>>52)&0xF)
	demo.Writer.TempAppendLine("\t\t\t\tSteam Account Universe: %d", uint32(userInfo.SteamID>>56)&0xFF)
	demo.Writer.TempAppendLine("\t\t\t\tName: %s", userInfo.Name)
	demo.Writer.TempAppendLine("\t\t\t\tUser ID: %d", userInfo.UserID)
	demo.Writer.TempAppendLine("\t\t\t\tGUID: %s", userInfo.GUID)
	demo.Writer.TempAppendLine("\t\t\t\tFriends ID: %d", userInfo.FriendsID)
	demo.Writer.TempAppendLine("\t\t\t\tFriends Name: %s", userInfo.FriendsName)
	demo.Writer.TempAppendLine("\t\t\t\tFake Player: %t", userInfo.FakePlayer)
	demo.Writer.TempAppendLine("\t\t\t\tIs Htlv: %t", userInfo.IsHltv)
	if userInfo.CustomFiles != nil {
		demo.Writer.TempAppendLine("\t\t\t\tCustom File CRCs: [logo: 0x%d, sounds: 0x%d, models: 0x%d, txt: 0x%d]", userInfo.CustomFiles[0], userInfo.CustomFiles[1], userInfo.CustomFiles[2], userInfo.CustomFiles[3])
	}
	demo.Writer.TempAppendLine("\t\t\t\tFiles Downloaded: %d", userInfo.FilesDownloaded)
}

func (stringTableEntry *StringTableEntry) ParseServerQueryInfo(reader *bitreader.Reader, demo *types.Demo) {
	serverQueryInfo := struct{ Port uint32 }{
		Port: reader.TryReadUInt32(),
	}
	stringTableEntry.EntryData = serverQueryInfo
	demo.Writer.TempAppendLine("\t\t\t\tPort: %d", serverQueryInfo.Port)
}

func (stringTableEntry *StringTableEntry) ParseGamesRulesCreation(reader *bitreader.Reader, demo *types.Demo) {
	gamesRulesCreation := struct{ Message string }{
		Message: reader.TryReadString(),
	}
	stringTableEntry.EntryData = gamesRulesCreation
	demo.Writer.TempAppendLine("\t\t\t\tMessage: %s", gamesRulesCreation.Message)
}

func (stringTableEntry *StringTableEntry) ParseInfoPanel(reader *bitreader.Reader, demo *types.Demo) {
	infoPanel := struct{ Message string }{
		Message: reader.TryReadString(),
	}
	stringTableEntry.EntryData = infoPanel
	demo.Writer.TempAppendLine("\t\t\t\tMessage: %s", infoPanel.Message)
}

func (stringTableEntry *StringTableEntry) ParseLightStyles(reader *bitreader.Reader, demo *types.Demo) {
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
		demo.Writer.TempAppendLine("\t\t\t\t0 Frames (256)")
	} else {
		demo.Writer.TempAppendLine("\t\t\t\t%d frames: %v", len(lightStyles.Values), lightStyles.Values)
	}
}

func (stringTableEntry *StringTableEntry) ParsePrecacheData(reader *bitreader.Reader, demo *types.Demo) {
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
	demo.Writer.TempAppendLine("\t\t\t\tFlags: %v", getFlags(PrecacheFlag(precacheData.Flags)))
}

func (stringTableEntry *StringTableEntry) ParseUnknown(reader *bitreader.Reader, demo *types.Demo) {
	unknown := reader.TryReadBitsToSlice(reader.TryReadRemainingBits())
	binaryString := ""
	for _, byteValue := range unknown {
		binaryString += fmt.Sprintf("%08b ", byteValue)
	}
	demo.Writer.TempAppendLine("\t\t\t\tUnknown: (%s)", strings.TrimSpace(binaryString))
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
