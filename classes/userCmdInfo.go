package classes

import (
	"parser/utils"

	"github.com/potterxu/bitreader"
)

type UserCmdInfo struct {
	CommandNumber int32
	TickCount     int32
	ViewAnglesX   float32
	ViewAnglesY   float32
	ViewAnglesZ   float32
	ForwardMove   float32
	SideMove      float32
	UpMove        float32
}

func UserCmdInfoInit(byteArr []byte, size int) (output UserCmdInfo) {
	var class UserCmdInfo
	reversedByteArr := utils.ReverseByteArrayValues(byteArr, size)
	reader := bitreader.BitReader(reversedByteArr)
	if size-1 >= 4 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.CommandNumber = int32(value)
		} else {
			return class
		}
	}
	if size-1 >= 8 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.TickCount = int32(value)
		} else {
			return class
		}
	}
	if size-1 >= 12 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.ViewAnglesX = float32(value)
		} else {
			return class
		}
	}
	if size-1 >= 16 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.ViewAnglesY = float32(value)
		} else {
			return class
		}
	}
	if size-1 >= 20 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.ViewAnglesZ = float32(value)
		} else {
			return class
		}
	}
	if size-1 >= 24 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.ForwardMove = float32(value)
		} else {
			return class
		}
	}
	if size-1 >= 28 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.SideMove = float32(value)
		} else {
			return class
		}
	}
	if size-1 >= 32 {
		bit, err := reader.ReadBit()
		utils.CheckError(err)
		if bit {
			value, err := reader.ReadBits(32)
			utils.CheckError(err)
			class.UpMove = float32(value)
		} else {
			return class
		}
	}
	return class
}
