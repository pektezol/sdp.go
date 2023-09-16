package classes

import (
	"github.com/pektezol/bitreader"
)

type CmdInfo struct {
	Flags            int32
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

func ParseCmdInfo(reader *bitreader.Reader) CmdInfo {
	return CmdInfo{
		Flags:            int32(reader.TryReadBits(32)),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	}
}
