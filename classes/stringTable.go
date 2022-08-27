package classes

type StringTable struct {
	NumOfTables        int32
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

/*
func StringTableInit(bytes []byte) (output StringTable) {
	var class StringTable
	class.NumOfTables = int(utils.IntFromBytes(bytes[:1]))
	class.TableName = string(bytes[1:16])
	class.ViewAngles = utils.FloatArrFromBytes(bytes[16:28])
	class.LocalViewAngles = utils.FloatArrFromBytes(bytes[28:40])
	class.ViewOrigin2 = utils.FloatArrFromBytes(bytes[40:52])
	class.ViewAngles2 = utils.FloatArrFromBytes(bytes[52:64])
	class.LocalViewAngles2 = utils.FloatArrFromBytes(bytes[64:76])
	return class
}*/
