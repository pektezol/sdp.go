package classes

import (
	"github.com/bisaxa/demoparser/utils"
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
	// Impulse       byte	}
	// WeaponSelect  int32	}
	// WeaponSubtype int32	Not worth the effort, no one cares about these
	// MouseDx       int16	}
	// MouseDy       int	}
}

// It is so janky it hurts, but hey it is at least working (hopefully)
// Reading the data is really weird, who even implemented this smh
func UserCmdInfoInit(byteArr []byte, size int32) (output UserCmdInfo) {
	var class UserCmdInfo
	successCount := 0
	failedCount := 0
	looped := 0
	classIndex := 0
	// fmt.Println(byteArr)
	// fmt.Printf("%08b", byteArr)
	for i := 0; i < 9; i++ {
		if successCount+failedCount > 7 {
			failedCount = -successCount
			looped++
		}
		firstBit, err := utils.ReadBitStateLSB(byteArr[successCount*4+looped], successCount+failedCount)
		utils.CheckError(err)
		if firstBit {
			successCount++
			switch classIndex {
			case 0:
				class.CommandNumber = utils.Read32BitsAfterFirstBitInt32(byteArr, successCount+failedCount, successCount*4+looped)
			case 1:
				class.TickCount = utils.Read32BitsAfterFirstBitInt32(byteArr, successCount+failedCount, successCount*4+looped)
			case 2:
				class.ViewAnglesX = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 3:
				class.ViewAnglesY = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 4:
				class.ViewAnglesZ = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 5:
				class.ForwardMove = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 6:
				class.SideMove = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 7:
				class.UpMove = utils.Read32BitsAfterFirstBitFloat32(byteArr, successCount+failedCount, successCount*4+looped)
			case 8:
				class.Buttons = utils.Read32BitsAfterFirstBitInt32(byteArr, successCount+failedCount, successCount*4+looped)
			}
			classIndex++
		} else {
			failedCount++
			classIndex++
		}
	}
	return class
}
