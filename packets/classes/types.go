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
	CommandNumber int
	TickCount     int
	ViewAnglesX   float32
	ViewAnglesY   float32
	ViewAnglesZ   float32
	ForwardMove   float32
	SideMove      float32
	UpMove        float32
	Buttons       int
	Impulse       byte
	WeaponSelect  int
	WeaponSubtype int
	MouseDx       int16
	MouseDy       int16
}
