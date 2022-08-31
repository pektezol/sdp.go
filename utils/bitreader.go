package utils

import (
	"fmt"
	"math"
	"strconv"
)

func ReadButtonsDataFromInt32(input int32) []string {
	buttonList := [32]string{
		"Attack",
		"Jump",
		"Duck",
		"Forward",
		"Back",
		"Use",
		"Cancel",
		"Left",
		"Right",
		"MoveLeft",
		"MoveRight",
		"Attack2",
		"Run",
		"Reload",
		"Alt1",
		"Alt2",
		"Score",
		"Speed",
		"Walk",
		"Zoom",
		"Weapon1",
		"Weapon2",
		"BullRush",
		"Grenade1",
		"Grenade2",
		"LookSpin",
		"CurrentAbility",
		"PreviousAbility",
		"Ability1",
		"Ability2",
		"Ability3",
		"Ability4",
	}
	var buttons []string
	if input == 0 {
		buttons = append(buttons, buttonList[0])
		return buttons
	}
	for i := 1; i < 33; i++ {
		if ReadBitState(input, i) {
			buttons = append(buttons, buttonList[i])
		}
	}
	return buttons
}

func ReadBitState(input int32, index int) bool {
	value := input & (1 << index)
	return value > 0
}

func ReadBitStateLSB(input byte, index int) (bool, error) {
	if index < 0 && index > 7 {
		return false, fmt.Errorf("IndexOutOfBounds for type byte")
	}
	value := input & (1 << index)
	return (value > 0), nil
}

func Read32BitsAfterFirstBitInt32(input []byte, index int, step int) int32 {
	binary := ""
	binary += fmt.Sprintf("%08b", input[step])[8-index : 8]
	binary += fmt.Sprintf("%08b", input[step-1])
	binary += fmt.Sprintf("%08b", input[step-2])
	binary += fmt.Sprintf("%08b", input[step-3])
	binary += fmt.Sprintf("%08b", input[step-4])[:8-index]
	output, err := strconv.ParseInt(binary, 2, 32)
	CheckError(err)
	return int32(output)

}

func Read32BitsAfterFirstBitFloat32(input []byte, index int, step int) float32 {
	binary := ""
	binary += fmt.Sprintf("%08b", input[step])[8-index : 8]
	binary += fmt.Sprintf("%08b", input[step-1])
	binary += fmt.Sprintf("%08b", input[step-2])
	binary += fmt.Sprintf("%08b", input[step-3])
	binary += fmt.Sprintf("%08b", input[step-4])[:8-index]
	output, err := strconv.ParseUint(binary, 2, 32)
	CheckError(err)
	return math.Float32frombits(uint32(output))

}
