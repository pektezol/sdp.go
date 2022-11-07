package classes

import (
	"github.com/pektezol/bitreader"
)

func ParseUserCmdInfo(reader *bitreader.ReaderType, size int) UserCmdInfo {
	var bitCount int
	var userCmdInfo UserCmdInfo
	if reader.TryReadBool() {
		userCmdInfo.CommandNumber = int(reader.TryReadInt32())
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.TickCount = int(reader.TryReadInt32())
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesX = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesY = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesZ = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.ForwardMove = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.SideMove = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.UpMove = reader.TryReadFloat32()
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.Buttons = int(reader.TryReadInt32())
		bitCount += 32
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.Impulse = reader.TryReadInt8()
		bitCount += 8
	}
	bitCount++
	if reader.TryReadBool() {
		value, err := reader.ReadBits(11)
		if err != nil {
			panic(err)
		}
		userCmdInfo.WeaponSelect = int(value)
		bitCount += 11
		if reader.TryReadBool() {
			value, err := reader.ReadBits(6)
			if err != nil {
				panic(err)
			}
			userCmdInfo.WeaponSubtype = int(value)
			bitCount += 6
		}
		bitCount++
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.MouseDx = int16(reader.TryReadInt16())
		bitCount += 16
	}
	bitCount++
	if reader.TryReadBool() {
		userCmdInfo.MouseDy = int16(reader.TryReadInt16())
		bitCount += 16
	}
	bitCount++
	/*if bitCount > size*8 {
		//reader.SkipBits(size * 8)
		return userCmdInfo
	}*/
	reader.SkipBits(size*8 - bitCount)
	return userCmdInfo
}
