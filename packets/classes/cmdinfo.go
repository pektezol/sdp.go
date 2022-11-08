package classes

import "github.com/pektezol/bitreader"

type CmdInfo struct {
	Flags            int32
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

func ParseCmdInfo(reader *bitreader.ReaderType, MSSC int) []CmdInfo {
	var out []CmdInfo
	for i := 0; i < MSSC; i++ {
		flags := reader.TryReadInt32()
		viewOrigin := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		viewAngles := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		localViewAngles := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		viewOrigin2 := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		viewAngles2 := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		localViewAngles2 := []float32{
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
			reader.TryReadFloat32(),
		}
		cmdInfo := CmdInfo{
			Flags:            int32(flags),
			ViewOrigin:       viewOrigin,
			ViewAngles:       viewAngles,
			LocalViewAngles:  localViewAngles,
			ViewOrigin2:      viewOrigin2,
			ViewAngles2:      viewAngles2,
			LocalViewAngles2: localViewAngles2,
		}
		out = append(out, cmdInfo)
	}
	return out
}
