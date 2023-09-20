package classes

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type CmdInfo struct {
	Flags            string
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

type CmdInfoFlags int

const (
	ECmdInfoFlagsNone        = 0
	ECmdInfoFlagsUseOrigin2  = 1
	ECmdInfoFlagsUserAngles2 = 1 << 1
	ECmdInfoFlagsNoInterp    = 1 << 2
)

func (cmdInfoFlags CmdInfoFlags) String() string {
	switch cmdInfoFlags {
	case ECmdInfoFlagsNone:
		return "None"
	case ECmdInfoFlagsUseOrigin2:
		return "UseOrigin2"
	case ECmdInfoFlagsUserAngles2:
		return "UserAngles2"
	case ECmdInfoFlagsNoInterp:
		return "NoInterp"
	default:
		return fmt.Sprintf("%d", int(cmdInfoFlags))
	}
}

func ParseCmdInfo(reader *bitreader.Reader) CmdInfo {
	cmdInfo := CmdInfo{
		Flags:            CmdInfoFlags(reader.TryReadUInt32()).String(),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	}
	writer.AppendLine("\tFlags: %s", cmdInfo.Flags)
	writer.AppendLine("\tView Origin: %v", cmdInfo.ViewOrigin)
	writer.AppendLine("\tView Angles: %v", cmdInfo.ViewAngles)
	writer.AppendLine("\tLocal View Angles: %v", cmdInfo.LocalViewAngles)
	writer.AppendLine("\tView Origin 2: %v", cmdInfo.ViewOrigin2)
	writer.AppendLine("\tView Angles 2: %v", cmdInfo.ViewAngles2)
	writer.AppendLine("\tLocal View Angles 2: %v", cmdInfo.LocalViewAngles2)
	writer.AppendLine("")
	return cmdInfo
}
