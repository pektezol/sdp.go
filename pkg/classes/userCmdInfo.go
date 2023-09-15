package classes

import (
	"bytes"

	"github.com/pektezol/bitreader"
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
	Buttons       int32
	Impulse       int8
	WeaponSelect  int16
	WeaponSubType int8
	MouseDx       int16
	MouseDy       int16
}

func ParseUserCmdInfo(data []byte) UserCmdInfo {
	reader := bitreader.NewReader(bytes.NewReader(data), true)
	userCmdInfo := UserCmdInfo{}
	if reader.TryReadBool() {
		userCmdInfo.CommandNumber = int32(reader.TryReadBits(32))
	}
	if reader.TryReadBool() {
		userCmdInfo.TickCount = int32(reader.TryReadBits(32))
	}
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesX = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesY = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.ViewAnglesZ = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.ForwardMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.SideMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.UpMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmdInfo.Buttons = int32(reader.TryReadBits(32))
	}
	if reader.TryReadBool() {
		userCmdInfo.Impulse = int8(reader.TryReadBits(8))
	}
	if reader.TryReadBool() {
		userCmdInfo.WeaponSelect = int16(reader.TryReadBits(11))
		if reader.TryReadBool() {
			userCmdInfo.WeaponSubType = int8(reader.TryReadBits(6))
		}
	}
	if reader.TryReadBool() {
		userCmdInfo.MouseDx = int16(reader.TryReadBits(16))
	}
	if reader.TryReadBool() {
		userCmdInfo.MouseDy = int16(reader.TryReadBits(16))
	}
	return userCmdInfo
}
