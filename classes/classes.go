package classes

import (
	"os"

	"github.com/bisaxa/bitreader"
	"github.com/bisaxa/demoparser/utils"
)

func ParseCmdInfo(file *os.File, MSSC int) []CmdInfo {
	reader := bitreader.Reader(file, true)
	var cmdinfo CmdInfo
	var cmdinfoarray []CmdInfo
	for count := 0; count < MSSC; count++ {
		cmdinfo.Flags = int32(reader.TryReadInt32())
		var floatArray [3]float32
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.ViewOrigin = floatArray[:]
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.ViewAngles = floatArray[:]
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.LocalViewAngles = floatArray[:]
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.ViewOrigin2 = floatArray[:]
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.ViewAngles2 = floatArray[:]
		for i := 0; i < 3; i++ {
			floatArray[i] = reader.TryReadFloat32()
		}
		cmdinfo.LocalViewAngles2 = floatArray[:]
		cmdinfoarray = append(cmdinfoarray, cmdinfo)
	}
	return cmdinfoarray
}

func ParseUserCmdInfo(file *os.File, size int) UserCmdInfo {
	count := 0
	reader := bitreader.Reader(file, true)
	var usercmd UserCmdInfo
	flag, err := reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.CommandNumber = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.TickCount = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesX = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesY = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesZ = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ForwardMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.SideMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.UpMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.Buttons = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		//reader.SkipBits(8)
		usercmd.Impulse = int8(reader.TryReadInt8())
		count += 8
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		value, err := reader.ReadBits(11)
		utils.CheckError(err)
		usercmd.WeaponSelect = int(value)
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		count += 11
		if flag {
			value, err := reader.ReadBits(6)
			utils.CheckError(err)
			usercmd.WeaponSubtype = int(value)
			count += 6
		}
		count++
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.MouseDx = int16(reader.TryReadInt16())
		count += 16
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.MouseDy = int16(reader.TryReadInt16())
		count += 16
	}
	count++
	reader.SkipBits(size*8 - count)
	return usercmd
}
