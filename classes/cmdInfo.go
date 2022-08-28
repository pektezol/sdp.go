package classes

import "github.com/bisaxa/demoparser/utils"

type CmdInfo struct {
	Flags            int32
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

func CmdInfoInit(bytes []byte) (output CmdInfo) {
	var class CmdInfo
	class.Flags = int32(utils.IntFromBytes(bytes[:4]))
	class.ViewOrigin = utils.FloatArrFromBytes(bytes[4:16])
	class.ViewAngles = utils.FloatArrFromBytes(bytes[16:28])
	class.LocalViewAngles = utils.FloatArrFromBytes(bytes[28:40])
	class.ViewOrigin2 = utils.FloatArrFromBytes(bytes[40:52])
	class.ViewAngles2 = utils.FloatArrFromBytes(bytes[52:64])
	class.LocalViewAngles2 = utils.FloatArrFromBytes(bytes[64:76])
	return class
}
