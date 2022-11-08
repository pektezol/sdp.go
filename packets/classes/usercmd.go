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
	Impulse       byte
	WeaponSelect  int
	WeaponSubtype int
	MouseDx       int16
	MouseDy       int16
}

func ParseUserCmdInfo(data []byte) UserCmdInfo {
	reader := bitreader.Reader(bytes.NewReader(data), true)
	var userCmdInfo UserCmdInfo
	if reader.TryReadBool() {
		userCmdInfo.CommandNumber = int32(reader.TryReadInt32())
	}
	if reader.TryReadBool() {
		userCmdInfo.TickCount = int32(reader.TryReadInt32())
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
		userCmdInfo.Buttons = int32(reader.TryReadInt32())
	}
	if reader.TryReadBool() {
		userCmdInfo.Impulse = reader.TryReadInt8()
	}
	if reader.TryReadBool() {
		value, err := reader.ReadBits(11)
		if err != nil {
			panic(err)
		}
		userCmdInfo.WeaponSelect = int(value)
		if reader.TryReadBool() {
			value, err := reader.ReadBits(6)
			if err != nil {
				panic(err)
			}
			userCmdInfo.WeaponSubtype = int(value)
		}
	}
	if reader.TryReadBool() {
		userCmdInfo.MouseDx = int16(reader.TryReadInt16())
	}
	if reader.TryReadBool() {
		userCmdInfo.MouseDy = int16(reader.TryReadInt16())
	}
	return userCmdInfo
}
