package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type UserCmd struct {
	Cmd  uint32      `json:"cmd"`
	Size uint32      `json:"size"`
	Data UserCmdInfo `json:"data"`
}

type UserCmdInfo struct {
	CommandNumber uint32  `json:"command_number"`
	TickCount     uint32  `json:"tick_count"`
	ViewAnglesX   float32 `json:"view_angles_x"`
	ViewAnglesY   float32 `json:"view_angles_y"`
	ViewAnglesZ   float32 `json:"view_angles_z"`
	ForwardMove   float32 `json:"forward_move"`
	SideMove      float32 `json:"side_move"`
	UpMove        float32 `json:"up_move"`
	Buttons       uint32  `json:"buttons"`
	Impulse       uint8   `json:"impulse"`
	WeaponSelect  uint16  `json:"weapon_select"`
	WeaponSubType uint8   `json:"weapon_sub_type"`
	MouseDx       uint16  `json:"mouse_dx"`
	MouseDy       uint16  `json:"mouse_dy"`
}

func (userCmd *UserCmd) ParseUserCmd(reader *bitreader.Reader, demo *types.Demo) {
	userCmd.Cmd = reader.TryReadUInt32()
	userCmd.Size = reader.TryReadUInt32()
	userCmdReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(userCmd.Size)), true)
	userCmd.ParseUserCmdInfo(userCmdReader, demo)
}

func (userCmd *UserCmd) ParseUserCmdInfo(reader *bitreader.Reader, demo *types.Demo) {
	if reader.TryReadBool() {
		userCmd.Data.CommandNumber = reader.TryReadUInt32()
	}
	if reader.TryReadBool() {
		userCmd.Data.TickCount = reader.TryReadUInt32()
	}
	if reader.TryReadBool() {
		userCmd.Data.ViewAnglesX = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.ViewAnglesY = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.ViewAnglesZ = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.ForwardMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.SideMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.UpMove = reader.TryReadFloat32()
	}
	if reader.TryReadBool() {
		userCmd.Data.Buttons = reader.TryReadUInt32()
	}
	if reader.TryReadBool() {
		userCmd.Data.Impulse = reader.TryReadUInt8()
	}
	if reader.TryReadBool() {
		userCmd.Data.WeaponSelect = uint16(reader.TryReadBits(11))
		if reader.TryReadBool() {
			userCmd.Data.WeaponSubType = uint8(reader.TryReadBits(6))
		}
	}
	if reader.TryReadBool() {
		userCmd.Data.MouseDx = reader.TryReadUInt16()
	}
	if reader.TryReadBool() {
		userCmd.Data.MouseDy = reader.TryReadUInt16()
	}
	demo.Writer.AppendLine("\tCommand Number: %v", userCmd.Data.CommandNumber)
	demo.Writer.AppendLine("\tTick Count: %v", userCmd.Data.TickCount)
	demo.Writer.AppendLine("\tView Angles: %v", []float32{userCmd.Data.ViewAnglesX, userCmd.Data.ViewAnglesY, userCmd.Data.ViewAnglesZ})
	demo.Writer.AppendLine("\tMovement: %v", []float32{userCmd.Data.ForwardMove, userCmd.Data.SideMove, userCmd.Data.UpMove})
	demo.Writer.AppendLine("\tButtons: %v", Buttons(userCmd.Data.Buttons).GetButtons())
	demo.Writer.AppendLine("\tImpulse: %v", userCmd.Data.Impulse)
	demo.Writer.AppendLine("\tWeapon, Subtype: %v, %v", userCmd.Data.WeaponSelect, userCmd.Data.WeaponSubType)
	demo.Writer.AppendLine("\tMouse Dx, Mouse Dy: %v, %v", userCmd.Data.MouseDx, userCmd.Data.MouseDy)
}

func (button Buttons) GetButtons() []string {
	flags := []string{}
	if button == 0 {
		flags = append(flags, EButtonsNone)
	}
	if checkBit(uint32(button), 0) {
		flags = append(flags, EButtonsAttack)
	}
	if checkBit(uint32(button), 1) {
		flags = append(flags, EButtonsJump)
	}
	if checkBit(uint32(button), 2) {
		flags = append(flags, EButtonsDuck)
	}
	if checkBit(uint32(button), 3) {
		flags = append(flags, EButtonsForward)
	}
	if checkBit(uint32(button), 4) {
		flags = append(flags, EButtonsBack)
	}
	if checkBit(uint32(button), 5) {
		flags = append(flags, EButtonsUse)
	}
	if checkBit(uint32(button), 6) {
		flags = append(flags, EButtonsCancel)
	}
	if checkBit(uint32(button), 7) {
		flags = append(flags, EButtonsLeft)
	}
	if checkBit(uint32(button), 8) {
		flags = append(flags, EButtonsRight)
	}
	if checkBit(uint32(button), 9) {
		flags = append(flags, EButtonsMoveLeft)
	}
	if checkBit(uint32(button), 10) {
		flags = append(flags, EButtonsMoveRight)
	}
	if checkBit(uint32(button), 11) {
		flags = append(flags, EButtonsAttack2)
	}
	if checkBit(uint32(button), 12) {
		flags = append(flags, EButtonsRun)
	}
	if checkBit(uint32(button), 13) {
		flags = append(flags, EButtonsReload)
	}
	if checkBit(uint32(button), 14) {
		flags = append(flags, EButtonsAlt1)
	}
	if checkBit(uint32(button), 15) {
		flags = append(flags, EButtonsAlt2)
	}
	if checkBit(uint32(button), 16) {
		flags = append(flags, EButtonsScore)
	}
	if checkBit(uint32(button), 17) {
		flags = append(flags, EButtonsSpeed)
	}
	if checkBit(uint32(button), 18) {
		flags = append(flags, EButtonsWalk)
	}
	if checkBit(uint32(button), 19) {
		flags = append(flags, EButtonsZoom)
	}
	if checkBit(uint32(button), 20) {
		flags = append(flags, EButtonsWeapon1)
	}
	if checkBit(uint32(button), 21) {
		flags = append(flags, EButtonsWeapon2)
	}
	if checkBit(uint32(button), 22) {
		flags = append(flags, EButtonsBullRush)
	}
	if checkBit(uint32(button), 23) {
		flags = append(flags, EButtonsGrenade1)
	}
	if checkBit(uint32(button), 24) {
		flags = append(flags, EButtonsGrenade2)
	}
	if checkBit(uint32(button), 25) {
		flags = append(flags, EButtonsLookSpin)
	}
	if checkBit(uint32(button), 26) {
		flags = append(flags, EButtonsCurrentAbility)
	}
	if checkBit(uint32(button), 27) {
		flags = append(flags, EButtonsPreviousAbility)
	}
	if checkBit(uint32(button), 28) {
		flags = append(flags, EButtonsAbility1)
	}
	if checkBit(uint32(button), 29) {
		flags = append(flags, EButtonsAbility2)
	}
	if checkBit(uint32(button), 30) {
		flags = append(flags, EButtonsAbility3)
	}
	if checkBit(uint32(button), 31) {
		flags = append(flags, EButtonsAbility4)
	}
	return flags
}

type Buttons int

const (
	EButtonsNone            string = "None"
	EButtonsAttack          string = "Attack"
	EButtonsJump            string = "Jump"
	EButtonsDuck            string = "Duck"
	EButtonsForward         string = "Forward"
	EButtonsBack            string = "Back"
	EButtonsUse             string = "Use"
	EButtonsCancel          string = "Cancel"
	EButtonsLeft            string = "Left"
	EButtonsRight           string = "Right"
	EButtonsMoveLeft        string = "MoveLeft"
	EButtonsMoveRight       string = "MoveRight"
	EButtonsAttack2         string = "Attack2"
	EButtonsRun             string = "Run"
	EButtonsReload          string = "Reload"
	EButtonsAlt1            string = "Alt1"
	EButtonsAlt2            string = "Alt2"
	EButtonsScore           string = "Score"
	EButtonsSpeed           string = "Speed"
	EButtonsWalk            string = "Walk"
	EButtonsZoom            string = "Zoom"
	EButtonsWeapon1         string = "Weapon1"
	EButtonsWeapon2         string = "Weapon2"
	EButtonsBullRush        string = "BullRush"
	EButtonsGrenade1        string = "Grenade1"
	EButtonsGrenade2        string = "Grenade2"
	EButtonsLookSpin        string = "LookSpin"
	EButtonsCurrentAbility  string = "CurrentAbility"
	EButtonsPreviousAbility string = "PreviousAbility"
	EButtonsAbility1        string = "Ability1"
	EButtonsAbility2        string = "Ability2"
	EButtonsAbility3        string = "Ability3"
	EButtonsAbility4        string = "Ability4"
)
