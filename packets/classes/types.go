package classes

type CmdInfo struct {
	Flags            int32
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

type UserCmdInfo struct {
	CommandNumber int32
	TickCount     int32
	ViewAnglesX   float32
	ViewAnglesY   float32
	ViewAnglesZ   float32
	ForwardMove   float32
	SideMove      float32
	UpMove        float32
	Buttons       int32
	Impulse       byte
	WeaponSelect  int
	WeaponSubtype int
	MouseDx       int16
	MouseDy       int16
}

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
