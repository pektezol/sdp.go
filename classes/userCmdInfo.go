package classes

import (
	"fmt"
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
	Buttons       int32
	Impulse       byte
	/*WeaponSelect int
	WeaponSubtype int
	MouseDx int16
	MouseDy int16*/
}

func UserCmdInfoInit(byteArr []byte, size int) (output UserCmdInfo) {
	var class UserCmdInfo
	fmt.Println("Size: ", size)
	reader := bitreader.BitReader(byteArr)
	flag, err := reader.ReadBit()
	utils.CheckError(err)
	fmt.Println(flag)
	flag2, err := reader.ReadBit()
	utils.CheckError(err)
	fmt.Println(flag2)
	/*if size-1 >= 4 {
		class.CommandNumber = int32(utils.IntFromBytes(byteArr[:4]))
	}
	if size-1 >= 8 {
		class.TickCount = int32(utils.IntFromBytes(byteArr[4:8]))
	}
	if size-1 >= 12 {
		class.ViewAnglesX = utils.FloatFromBytes(byteArr[8:12])
	}
	if size-1 >= 16 {
		class.ViewAnglesY = utils.FloatFromBytes(byteArr[12:16])
	}
	if size-1 >= 20 {
		class.ViewAnglesZ = utils.FloatFromBytes(byteArr[16:20])
	}
	if size-1 >= 24 {
		class.ForwardMove = utils.FloatFromBytes(byteArr[20:24])
	}
	if size-1 >= 28 {
		class.SideMove = utils.FloatFromBytes(byteArr[24:28])
	}
	if size-1 >= 32 {
		class.UpMove = utils.FloatFromBytes(byteArr[28:32])
	}
	if size-1 >= 36 {
		class.Buttons = int32(utils.IntFromBytes(byteArr[32:36]))
	}
	if size-1 >= 40 {
		class.Impulse = byteArr[36]
	}*/
	return class
}
