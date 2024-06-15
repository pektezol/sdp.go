package messages

import (
	"fmt"
	"math"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcUserMessage struct {
	Type   int8  `json:"type"`
	Length int16 `json:"length"`
	Data   any   `json:"data"`
}

func ParseSvcUserMessage(reader *bitreader.Reader, demo *types.Demo) SvcUserMessage {
	svcUserMessage := SvcUserMessage{
		Type:   int8(reader.TryReadBits(8)),
		Length: int16(reader.TryReadBits(12)),
	}
	svcUserMessage.Data = reader.TryReadBitsToSlice(uint64(svcUserMessage.Length))
	userMessageReader := bitreader.NewReaderFromBytes(svcUserMessage.Data.([]byte), true)
	demo.Writer.TempAppendLine("\t\t%s (%d):", UserMessageType(svcUserMessage.Type).String(), svcUserMessage.Type)
	switch UserMessageType(svcUserMessage.Type) {
	case EUserMessageTypeGeiger:
		svcUserMessage.parseGeiger(userMessageReader, demo)
	case EUserMessageTypeTrain:
		svcUserMessage.parseTrain(userMessageReader, demo)
	case EUserMessageTypeHudText:
		svcUserMessage.parseHUDText(userMessageReader, demo)
	case EUserMessageTypeSayText:
		svcUserMessage.parseSayText(userMessageReader, demo)
	case EUserMessageTypeSayText2:
		svcUserMessage.parseSayText2(userMessageReader, demo)
	case EUserMessageTypeTextMsg:
		svcUserMessage.parseTextMsg(userMessageReader, demo)
	case EUserMessageTypeHUDMsg:
		svcUserMessage.parseHUDMsg(userMessageReader, demo)
	case EUserMessageTypeResetHUD:
		svcUserMessage.parseResetHUD(userMessageReader, demo)
	case EUserMessageTypeShake:
		svcUserMessage.parseShake(userMessageReader, demo)
	case EUserMessageTypeFade:
		svcUserMessage.parseFade(userMessageReader, demo)
	case EUserMessageTypeVGUIMenu:
		svcUserMessage.parseVguiMenu(userMessageReader, demo)
	case EUserMessageTypeRumble:
		svcUserMessage.parseRumble(userMessageReader, demo)
	case EUserMessageTypeBattery:
		svcUserMessage.parseBattery(userMessageReader, demo)
	case EUserMessageTypeDamage:
		svcUserMessage.parseDamage(userMessageReader, demo)
	case EUserMessageTypeVoiceMask:
		svcUserMessage.parseVoiceMask(userMessageReader, demo)
	case EUserMessageTypeCloseCaption:
		svcUserMessage.parseCloseCaption(userMessageReader, demo)
	case EUserMessageTypeKeyHintText:
		svcUserMessage.parseKeyHintText(userMessageReader, demo)
	case EUserMessageTypeLogoTimeMsg:
		svcUserMessage.parseLogoTimeMsg(userMessageReader, demo)
	case EUserMessageTypeAchievementEvent:
		svcUserMessage.parseAchivementEvent(userMessageReader, demo)
	case EUserMessageTypeCurrentTimescale:
		svcUserMessage.parseCurrentTimescale(userMessageReader, demo)
	case EUserMessageTypeDesiredTimescale:
		svcUserMessage.parseDesiredTimescale(userMessageReader, demo)
	case EUserMessageTypeMPMapCompleted:
		svcUserMessage.parseMpMapCompleted(userMessageReader, demo)
	case EUserMessageTypeMPMapIncomplete:
		svcUserMessage.parseMpMapIncomplete(userMessageReader, demo)
	case EUserMessageTypeMPTauntEarned:
		svcUserMessage.parseMpTauntEarned(userMessageReader, demo)
	case EUserMessageTypeMPTauntLocked:
		svcUserMessage.parseMpTauntLocked(userMessageReader, demo)
	case EUserMessageTypePortalFX_Surface:
		svcUserMessage.parsePortalFxSurface(userMessageReader, demo)
	case EUserMessageTypePaintWorld:
		svcUserMessage.parsePaintWorld(userMessageReader, demo)
	case EUserMessageTypeTransitionFade:
		svcUserMessage.parseTransitionFade(userMessageReader, demo)
	case EUserMessageTypeScoreboardTempUpdate:
		svcUserMessage.parseScoreboardTempUpdate(userMessageReader, demo)
	default:
		demo.Writer.TempAppendLine("\t\t\tData: %v", svcUserMessage.Data)
	}
	return svcUserMessage
}

func (svcUserMessage *SvcUserMessage) parseGeiger(reader *bitreader.Reader, demo *types.Demo) {
	geiger := struct{ Range uint8 }{
		Range: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = geiger
	demo.Writer.TempAppendLine("\t\t\tGeiger Range: %d", geiger.Range)
}

func (svcUserMessage *SvcUserMessage) parseTrain(reader *bitreader.Reader, demo *types.Demo) {
	train := struct{ Pos uint8 }{
		Pos: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = train
	demo.Writer.TempAppendLine("\t\t\tPos: %d", train.Pos)
}

func (svcUserMessage *SvcUserMessage) parseHUDText(reader *bitreader.Reader, demo *types.Demo) {
	hudText := struct{ Text string }{
		Text: reader.TryReadString(),
	}
	svcUserMessage.Data = hudText
	demo.Writer.TempAppendLine("\t\t\tText: %s", hudText.Text)
}

func (svcUserMessage *SvcUserMessage) parseSayText(reader *bitreader.Reader, demo *types.Demo) {
	sayText := struct {
		Client      uint8  `json:"client"`
		Message     string `json:"message"`
		WantsToChat bool   `json:"wants_to_chat"`
	}{
		Client:      reader.TryReadUInt8(),
		Message:     reader.TryReadString(),
		WantsToChat: reader.TryReadUInt8() != 0,
	}
	svcUserMessage.Data = sayText
	demo.Writer.TempAppendLine("\t\t\tClient: %d", sayText.Client)
	demo.Writer.TempAppendLine("\t\t\tMessage: %s", sayText.Message)
	demo.Writer.TempAppendLine("\t\t\tWants To Chat: %t", sayText.WantsToChat)
}

func (svcUserMessage *SvcUserMessage) parseSayText2(reader *bitreader.Reader, demo *types.Demo) {
	sayText2 := struct {
		Client      uint8    `json:"client"`
		WantsToChat bool     `json:"wants_to_chat"`
		MessageName string   `json:"message_name"`
		Messages    []string `json:"messages"`
	}{
		Client:      reader.TryReadUInt8(),
		WantsToChat: reader.TryReadUInt8() != 0,
		MessageName: reader.TryReadString(),
		Messages:    []string{reader.TryReadString(), reader.TryReadString(), reader.TryReadString()},
	}
	svcUserMessage.Data = sayText2
	demo.Writer.TempAppendLine("\t\t\tClient: %d", sayText2.Client)
	demo.Writer.TempAppendLine("\t\t\tWants To Chat: %t", sayText2.WantsToChat)
	demo.Writer.TempAppendLine("\t\t\tName: %s", sayText2.MessageName)
	for index, message := range sayText2.Messages {
		demo.Writer.TempAppendLine("\t\t\tMessage %d: %s", index, message)
	}
}

func (svcUserMessage *SvcUserMessage) parseTextMsg(reader *bitreader.Reader, demo *types.Demo) {
	const MessageCount int = 5
	textMsg := struct {
		Destination uint8
		Messages    []string
	}{
		Destination: reader.TryReadUInt8(),
	}
	textMsg.Messages = make([]string, 5)
	for i := 0; i < MessageCount; i++ {
		textMsg.Messages[i] = reader.TryReadString()
	}
	svcUserMessage.Data = textMsg
	demo.Writer.TempAppendLine("\t\t\tDestination: %d", textMsg.Destination)
	for i := 0; i < MessageCount; i++ {
		demo.Writer.TempAppendLine("\t\t\tMessage %d: %s", i+1, textMsg.Messages)
	}
}

func (svcUserMessage *SvcUserMessage) parseHUDMsg(reader *bitreader.Reader, demo *types.Demo) {
	const MaxNetMessage uint8 = 6
	hudMsg := struct {
		Channel uint8 `json:"channel"`
		Info    struct {
			X        float32 `json:"x"` // 0-1 & resolution independent, -1 means center in each dimension
			Y        float32 `json:"y"`
			R1       uint8   `json:"r_1"`
			G1       uint8   `json:"g_1"`
			B1       uint8   `json:"b_1"`
			A1       uint8   `json:"a_1"`
			R2       uint8   `json:"r_2"`
			G2       uint8   `json:"g_2"`
			B2       uint8   `json:"b_2"`
			A2       uint8   `json:"a_2"`
			Effect   uint8   `json:"effect"`
			FadeIn   float32 `json:"fade_in"` // the fade times seem to be per character
			FadeOut  float32 `json:"fade_out"`
			HoldTime float32 `json:"hold_time"`
			FxTime   float32 `json:"fx_time"`
			Message  string  `json:"message"`
		} `json:"info"`
	}{
		Channel: reader.TryReadUInt8() % MaxNetMessage,
	}
	svcUserMessage.Data = hudMsg
	demo.Writer.TempAppendLine("\t\t\tChannel: %d", hudMsg.Channel)
	if reader.TryReadRemainingBits() >= 148 {
		hudMsg.Info = struct {
			X        float32 `json:"x"`
			Y        float32 `json:"y"`
			R1       uint8   `json:"r_1"`
			G1       uint8   `json:"g_1"`
			B1       uint8   `json:"b_1"`
			A1       uint8   `json:"a_1"`
			R2       uint8   `json:"r_2"`
			G2       uint8   `json:"g_2"`
			B2       uint8   `json:"b_2"`
			A2       uint8   `json:"a_2"`
			Effect   uint8   `json:"effect"`
			FadeIn   float32 `json:"fade_in"`
			FadeOut  float32 `json:"fade_out"`
			HoldTime float32 `json:"hold_time"`
			FxTime   float32 `json:"fx_time"`
			Message  string  `json:"message"`
		}{
			X:        reader.TryReadFloat32(),
			Y:        reader.TryReadFloat32(),
			R1:       reader.TryReadUInt8(),
			G1:       reader.TryReadUInt8(),
			B1:       reader.TryReadUInt8(),
			A1:       reader.TryReadUInt8(),
			R2:       reader.TryReadUInt8(),
			G2:       reader.TryReadUInt8(),
			B2:       reader.TryReadUInt8(),
			A2:       reader.TryReadUInt8(),
			Effect:   reader.TryReadUInt8(),
			FadeIn:   reader.TryReadFloat32(),
			FadeOut:  reader.TryReadFloat32(),
			HoldTime: reader.TryReadFloat32(),
			FxTime:   reader.TryReadFloat32(),
			Message:  reader.TryReadString(),
		}
		svcUserMessage.Data = hudMsg
		demo.Writer.TempAppendLine("\t\t\tX: %f, Y: %f", hudMsg.Info.X, hudMsg.Info.Y)
		demo.Writer.TempAppendLine("\t\t\tRGBA1: %3d %3d %3d %3d", hudMsg.Info.R1, hudMsg.Info.G1, hudMsg.Info.B1, hudMsg.Info.A1)
		demo.Writer.TempAppendLine("\t\t\tRGBA2: %3d %3d %3d %3d", hudMsg.Info.R2, hudMsg.Info.G2, hudMsg.Info.B2, hudMsg.Info.A2)
		demo.Writer.TempAppendLine("\t\t\tEffect: %d", hudMsg.Info.Effect)
		demo.Writer.TempAppendLine("\t\t\tFade In: %f", hudMsg.Info.FadeIn)
		demo.Writer.TempAppendLine("\t\t\tFade Out: %f", hudMsg.Info.FadeOut)
		demo.Writer.TempAppendLine("\t\t\tHold Time: %f", hudMsg.Info.HoldTime)
		demo.Writer.TempAppendLine("\t\t\tFX Time: %f", hudMsg.Info.FxTime)
		demo.Writer.TempAppendLine("\t\t\tMessage: %s", hudMsg.Info.Message)
	}
}

func (svcUserMessage *SvcUserMessage) parseResetHUD(reader *bitreader.Reader, demo *types.Demo) {
	resetHUD := struct{ Unknown uint8 }{
		Unknown: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = resetHUD
	demo.Writer.TempAppendLine("\t\t\tUnknown: %d", resetHUD.Unknown)
}

func (svcUserMessage *SvcUserMessage) parseShake(reader *bitreader.Reader, demo *types.Demo) {
	type ShakeCommand uint8
	const (
		Start      ShakeCommand = iota // Starts the screen shake for all players within the radius.
		Stop                           // Stops the screen shake for all players within the radius.
		Amplitude                      // Modifies the amplitude of an active screen shake for all players within the radius.
		Frequency                      // Modifies the frequency of an active screen shake for all players within the radius.
		RumbleOnly                     // Starts a shake effect that only rumbles the controller, no screen effect.
		NoRumble                       // Starts a shake that does NOT rumble the controller.
	)
	shake := struct {
		Command   uint8   `json:"command"`
		Amplitude float32 `json:"amplitude"`
		Frequency float32 `json:"frequency"`
		Duration  float32 `json:"duration"`
	}{
		Command:   reader.TryReadUInt8(),
		Amplitude: reader.TryReadFloat32(),
		Frequency: reader.TryReadFloat32(),
		Duration:  reader.TryReadFloat32(),
	}
	shakeCommandToString := func(cmd ShakeCommand) string {
		switch cmd {
		case Start:
			return "Start"
		case Stop:
			return "Stop"
		case Amplitude:
			return "Amplitude"
		case Frequency:
			return "Frequency"
		case RumbleOnly:
			return "RumbleOnly"
		case NoRumble:
			return "NoRumble"
		default:
			return "Unknown"
		}
	}
	svcUserMessage.Data = shake
	demo.Writer.TempAppendLine("\t\t\tCommand: %s", shakeCommandToString(ShakeCommand(shake.Command)))
	demo.Writer.TempAppendLine("\t\t\tAmplitude: %f", shake.Amplitude)
	demo.Writer.TempAppendLine("\t\t\tFrequency: %f", shake.Frequency)
	demo.Writer.TempAppendLine("\t\t\tDuration: %f", shake.Duration)
}

func (svcUserMessage *SvcUserMessage) parseFade(reader *bitreader.Reader, demo *types.Demo) {
	type FadeFlag uint16
	const (
		None     FadeFlag = 0
		FadeIn   FadeFlag = 1
		FadeOut  FadeFlag = 1 << 1
		Modulate FadeFlag = 1 << 2 // Modulate (don't blend)
		StayOut  FadeFlag = 1 << 3 // ignores the duration, stays faded out until new ScreenFade message received
		Purge    FadeFlag = 1 << 4 // Purges all other fades, replacing them with this one
	)
	fade := struct {
		Duration float32 `json:"duration"`
		HoldTime uint16  `json:"hold_time"`
		Flags    uint16  `json:"flags"`
		R        uint8   `json:"r"`
		G        uint8   `json:"g"`
		B        uint8   `json:"b"`
		A        uint8   `json:"a"`
	}{
		Duration: float32(reader.TryReadUInt16()) / float32(1<<9), // might be useful: #define SCREENFADE_FRACBITS 9
		HoldTime: reader.TryReadUInt16(),
		Flags:    reader.TryReadUInt16(),
		R:        reader.TryReadUInt8(),
		G:        reader.TryReadUInt8(),
		B:        reader.TryReadUInt8(),
		A:        reader.TryReadUInt8(),
	}
	getFlags := func(flags FadeFlag) []string {
		var flagStrings []string
		if flags&FadeIn != 0 {
			flagStrings = append(flagStrings, "FadeIn")
		}
		if flags&FadeOut != 0 {
			flagStrings = append(flagStrings, "FadeOut")
		}
		if flags&Modulate != 0 {
			flagStrings = append(flagStrings, "Modulate")
		}
		if flags&StayOut != 0 {
			flagStrings = append(flagStrings, "StayOut")
		}
		if flags&Purge != 0 {
			flagStrings = append(flagStrings, "Purge")
		}
		return flagStrings
	}
	svcUserMessage.Data = fade
	demo.Writer.TempAppendLine("\t\t\tDuration: %f", fade.Duration)
	demo.Writer.TempAppendLine("\t\t\tHold Time: %d", fade.HoldTime)
	demo.Writer.TempAppendLine("\t\t\tFlags: %v", getFlags(FadeFlag(fade.Flags)))
	demo.Writer.TempAppendLine("\t\t\tRGBA: %3d %3d %3d %3d", fade.R, fade.G, fade.B, fade.A)
}

func (svcUserMessage *SvcUserMessage) parseVguiMenu(reader *bitreader.Reader, demo *types.Demo) {
	vguiMenu := struct {
		Message   string              `json:"message"`
		Show      bool                `json:"show"`
		KeyValues []map[string]string `json:"key_values"`
	}{
		Message: reader.TryReadString(),
		Show:    reader.TryReadUInt8() != 0,
	}
	count := reader.TryReadUInt8()
	for i := 0; i < int(count); i++ {
		vguiMenu.KeyValues = append(vguiMenu.KeyValues, map[string]string{"Key": reader.TryReadString(), "Value": reader.TryReadString()})
	}
	svcUserMessage.Data = vguiMenu
	demo.Writer.TempAppendLine("\t\t\tMessage: %s", vguiMenu.Message)
	demo.Writer.TempAppendLine("\t\t\tShow: %t", vguiMenu.Show)
	if len(vguiMenu.KeyValues) > 0 {
		demo.Writer.TempAppendLine("\t\t\t%d Key Value Pairs:", len(vguiMenu.KeyValues))
		for _, kv := range vguiMenu.KeyValues {
			demo.Writer.TempAppendLine("\t\t\t\t%s: %s", kv["Key"], kv["Value"])
		}
	} else {
		demo.Writer.TempAppendLine("\t\t\tNo Key Value Pairs")
	}
}

func (svcUserMessage *SvcUserMessage) parseRumble(reader *bitreader.Reader, demo *types.Demo) {
	type RumbleLookup int8
	const (
		RumbleInvalid          RumbleLookup = -1
		RumbleStopAll          RumbleLookup = 0 // cease all current rumbling effects.
		PhyscannonOpen         RumbleLookup = 20
		PhyscannonPunt         RumbleLookup = 21
		PhyscannonLow          RumbleLookup = 22
		PhyscannonMedium       RumbleLookup = 23
		PhyscannonHigh         RumbleLookup = 24
		PortalgunLeft          RumbleLookup = 25
		PortalgunRight         RumbleLookup = 26
		PortalPlacementFailure RumbleLookup = 27
	)
	getRumbleLookup := func(rumbleLookup RumbleLookup) string {
		switch rumbleLookup {
		case RumbleInvalid:
			return "RumbleInvalid"
		case RumbleStopAll:
			return "RumbleStopAll"
		case PhyscannonOpen:
			return "PhyscannonOpen"
		case PhyscannonPunt:
			return "PhyscannonPunt"
		case PhyscannonLow:
			return "PhyscannonLow"
		case PhyscannonMedium:
			return "PhyscannonMedium"
		case PhyscannonHigh:
			return "PhyscannonHigh"
		case PortalgunLeft:
			return "PortalgunLeft"
		case PortalgunRight:
			return "PortalgunRight"
		case PortalPlacementFailure:
			return "PortalPlacementFailure"
		default:
			return fmt.Sprintf("%d", int(rumbleLookup))
		}
	}
	type RumbleFlag uint8
	const (
		None            RumbleFlag = 0
		Stop            RumbleFlag = 1
		Loop            RumbleFlag = 1 << 1
		Restart         RumbleFlag = 1 << 2
		UpdateScale     RumbleFlag = 1 << 3 // Apply DATA to this effect if already playing, but don't restart.   <-- DATA is scale * 100
		OnlyOne         RumbleFlag = 1 << 4 // Don't play this effect if it is already playing.
		RandomAmplitude RumbleFlag = 1 << 4 // Amplitude scale will be randomly chosen. Between 10% and 100%
		InitialScale    RumbleFlag = 1 << 4 // Data is the initial scale to start this effect ( * 100 )
	)
	rumble := struct {
		Type  int8    `json:"type"`
		Scale float32 `json:"scale"`
		Flags uint8   `json:"flags"`
	}{
		Type:  reader.TryReadSInt8(),
		Scale: float32(reader.TryReadUInt8()) / 100,
		Flags: reader.TryReadUInt8(),
	}
	getFlags := func(flags RumbleFlag) []string {
		var flagStrings []string
		if flags&Stop != 0 {
			flagStrings = append(flagStrings, "Stop")
		}
		if flags&Loop != 0 {
			flagStrings = append(flagStrings, "Loop")
		}
		if flags&Restart != 0 {
			flagStrings = append(flagStrings, "Restart")
		}
		if flags&UpdateScale != 0 {
			flagStrings = append(flagStrings, "UpdateScale")
		}
		if flags&OnlyOne != 0 {
			flagStrings = append(flagStrings, "OnlyOne")
		}
		if flags&RandomAmplitude != 0 {
			flagStrings = append(flagStrings, "RandomAmplitude")
		}
		if flags&InitialScale != 0 {
			flagStrings = append(flagStrings, "InitialScale")
		}
		return flagStrings
	}
	svcUserMessage.Data = rumble
	demo.Writer.TempAppendLine("\t\t\tType: %s", getRumbleLookup(RumbleLookup(rumble.Type)))
	demo.Writer.TempAppendLine("\t\t\tScale: %f", rumble.Scale)
	demo.Writer.TempAppendLine("\t\t\tFlags: %v", getFlags(RumbleFlag(rumble.Flags)))
}

func (svcUserMessage *SvcUserMessage) parseBattery(reader *bitreader.Reader, demo *types.Demo) {
	battery := struct{ BatteryVal uint16 }{
		BatteryVal: reader.TryReadUInt16(),
	}
	svcUserMessage.Data = battery
	demo.Writer.TempAppendLine("\t\t\tBattery: %d", battery.BatteryVal)
}

func (svcUserMessage *SvcUserMessage) parseDamage(reader *bitreader.Reader, demo *types.Demo) {
	damage := struct {
		Armor       uint8     `json:"armor"`
		DamageTaken uint8     `json:"damage_taken"`
		BitsDamage  int32     `json:"bits_damage"`
		VecFrom     []float32 `json:"vec_from"`
	}{
		Armor:       reader.TryReadUInt8(),
		DamageTaken: reader.TryReadUInt8(),
		BitsDamage:  reader.TryReadSInt32(),
		VecFrom:     []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	}
	svcUserMessage.Data = damage
	demo.Writer.TempAppendLine("\t\t\tArmor: %d", damage.Armor)
	demo.Writer.TempAppendLine("\t\t\tDamage Taken: %d", damage.DamageTaken)
	demo.Writer.TempAppendLine("\t\t\tBits Damage: %d", damage.BitsDamage)
	demo.Writer.TempAppendLine("\t\t\tVecFrom: %v", damage.VecFrom)
}

func (svcUserMessage *SvcUserMessage) parseVoiceMask(reader *bitreader.Reader, demo *types.Demo) {
	// const VoiceMaxPlayers = 2
	voiceMask := struct {
		PlayerMasks []struct {
			GameRulesMask int32 `json:"game_rules_mask"`
			BanMask       int32 `json:"ban_mask"`
		} `json:"player_masks"`
		PlayerModEnable bool `json:"player_mod_enable"`
	}{
		PlayerMasks: []struct {
			GameRulesMask int32 `json:"game_rules_mask"`
			BanMask       int32 `json:"ban_mask"`
		}{
			{
				GameRulesMask: reader.TryReadSInt32(),
				BanMask:       reader.TryReadSInt32(),
			},
			{
				GameRulesMask: reader.TryReadSInt32(),
				BanMask:       reader.TryReadSInt32(),
			},
		},
		PlayerModEnable: reader.TryReadUInt8() != 0,
	}
	svcUserMessage.Data = voiceMask
	demo.Writer.TempAppendLine("\t\t\tPlayer Masks:")
	demo.Writer.TempAppendLine("\t\t\t\t[0] Game Rules Mask: %d", voiceMask.PlayerMasks[0].GameRulesMask)
	demo.Writer.TempAppendLine("\t\t\t\t[0] Ban Mask: %d", voiceMask.PlayerMasks[0].BanMask)
	demo.Writer.TempAppendLine("\t\t\t\t[1] Game Rules Mask: %d", voiceMask.PlayerMasks[1].GameRulesMask)
	demo.Writer.TempAppendLine("\t\t\t\t[1] Ban Mask: %d", voiceMask.PlayerMasks[1].BanMask)
	demo.Writer.TempAppendLine("\t\t\t\tPlayer Mod Enable: %t", voiceMask.PlayerModEnable)
}

func (svcUserMessage *SvcUserMessage) parseCloseCaption(reader *bitreader.Reader, demo *types.Demo) {
	type CloseCaptionFlag uint8
	const (
		None          CloseCaptionFlag = 0
		WarnIfMissing CloseCaptionFlag = 1
		FromPlayer    CloseCaptionFlag = 1 << 1
		GenderMale    CloseCaptionFlag = 1 << 2
		GenderFemale  CloseCaptionFlag = 1 << 3
	)
	closeCaption := struct {
		TokenName string  `json:"token_name"`
		Duration  float32 `json:"duration"`
		Flags     uint8   `json:"flags"`
	}{
		TokenName: reader.TryReadString(),
		Duration:  float32(reader.TryReadSInt16()) * 0.1,
		Flags:     reader.TryReadUInt8(),
	}
	getFlags := func(flags CloseCaptionFlag) []string {
		var flagStrings []string
		if flags&WarnIfMissing != 0 {
			flagStrings = append(flagStrings, "WarnIfMissing")
		}
		if flags&FromPlayer != 0 {
			flagStrings = append(flagStrings, "FromPlayer")
		}
		if flags&GenderMale != 0 {
			flagStrings = append(flagStrings, "GenderMale")
		}
		if flags&GenderFemale != 0 {
			flagStrings = append(flagStrings, "GenderFemale")
		}
		return flagStrings
	}
	svcUserMessage.Data = closeCaption
	demo.Writer.TempAppendLine("\t\t\tToken Name: %s", closeCaption.TokenName)
	demo.Writer.TempAppendLine("\t\t\tDuration: %f", closeCaption.Duration)
	demo.Writer.TempAppendLine("\t\t\tFlags: %v", getFlags(CloseCaptionFlag(closeCaption.Flags)))
}

func (svcUserMessage *SvcUserMessage) parseKeyHintText(reader *bitreader.Reader, demo *types.Demo) {
	keyHintText := struct {
		Count     uint8  `json:"count"`
		KeyString string `json:"key_string"`
	}{
		Count:     reader.TryReadUInt8(),
		KeyString: reader.TryReadString(),
	}
	svcUserMessage.Data = keyHintText
	demo.Writer.TempAppendLine("\t\t\tCount: %d", keyHintText.Count)
	demo.Writer.TempAppendLine("\t\t\tString: %s", keyHintText.KeyString)
}

func (svcUserMessage *SvcUserMessage) parseLogoTimeMsg(reader *bitreader.Reader, demo *types.Demo) {
	logoTimeMsg := struct {
		Time float32 `json:"time"`
	}{
		Time: reader.TryReadFloat32(),
	}
	svcUserMessage.Data = logoTimeMsg
	demo.Writer.TempAppendLine("\t\t\tTime: %f", logoTimeMsg.Time)
}

func (svcUserMessage *SvcUserMessage) parseAchivementEvent(reader *bitreader.Reader, demo *types.Demo) {
	achivementEvent := struct{ AchivementID int32 }{
		AchivementID: reader.TryReadSInt32(),
	}
	svcUserMessage.Data = achivementEvent
	demo.Writer.TempAppendLine("\t\t\tAchivement ID: %d", achivementEvent.AchivementID)
}

func (svcUserMessage *SvcUserMessage) parseCurrentTimescale(reader *bitreader.Reader, demo *types.Demo) {
	currentTimescale := struct {
		Timescale float32 `json:"timescale"`
	}{
		Timescale: reader.TryReadFloat32(),
	}
	svcUserMessage.Data = currentTimescale
	demo.Writer.TempAppendLine("\t\t\tTimescale: %f", currentTimescale.Timescale)
}

func (svcUserMessage *SvcUserMessage) parseDesiredTimescale(reader *bitreader.Reader, demo *types.Demo) {
	desiredTimescale := struct {
		Unk1 float32 `json:"unk_1"`
		Unk2 float32 `json:"unk_2"`
		Unk3 uint8   `json:"unk_3"`
		Unk4 float32 `json:"unk_4"`
	}{
		Unk1: reader.TryReadFloat32(),
		Unk2: reader.TryReadFloat32(),
		Unk3: reader.TryReadUInt8(),
		Unk4: reader.TryReadFloat32(),
	}
	svcUserMessage.Data = desiredTimescale
	demo.Writer.TempAppendLine("\t\t\tUnk1: %f", desiredTimescale.Unk1)
	demo.Writer.TempAppendLine("\t\t\tUnk2: %f", desiredTimescale.Unk2)
	demo.Writer.TempAppendLine("\t\t\tUnk3: %d", desiredTimescale.Unk3)
	demo.Writer.TempAppendLine("\t\t\tUnk4: %f", desiredTimescale.Unk4)
}

func (svcUserMessage *SvcUserMessage) parseMpMapCompleted(reader *bitreader.Reader, demo *types.Demo) {
	mpMapCompleted := struct {
		Branch uint8 `json:"branch"`
		Level  uint8 `json:"level"`
	}{
		Branch: reader.TryReadUInt8(),
		Level:  reader.TryReadUInt8(),
	}
	svcUserMessage.Data = mpMapCompleted
	demo.Writer.TempAppendLine("\t\t\tBranch: %d", mpMapCompleted.Branch)
	demo.Writer.TempAppendLine("\t\t\tLevel: %d", mpMapCompleted.Level)
}

func (svcUserMessage *SvcUserMessage) parseMpMapIncomplete(reader *bitreader.Reader, demo *types.Demo) {
}

func (svcUserMessage *SvcUserMessage) parseMpTauntEarned(reader *bitreader.Reader, demo *types.Demo) {
	mpTauntEarned := struct {
		TauntName     string `json:"taunt_name"`
		AwardSilently bool   `json:"award_silently"`
	}{
		TauntName:     reader.TryReadString(),
		AwardSilently: reader.TryReadBool(),
	}
	svcUserMessage.Data = mpTauntEarned
	demo.Writer.TempAppendLine("\t\t\tTaunt Name: %s", mpTauntEarned.TauntName)
	demo.Writer.TempAppendLine("\t\t\tAward Silently: %t", mpTauntEarned.AwardSilently)
}

func (svcUserMessage *SvcUserMessage) parseMpTauntLocked(reader *bitreader.Reader, demo *types.Demo) {
	mpTauntLocked := struct {
		TauntName string `json:"taunt_name"`
	}{
		TauntName: reader.TryReadString(),
	}
	svcUserMessage.Data = mpTauntLocked
	demo.Writer.TempAppendLine("\t\t\tTaunt Name: %s", mpTauntLocked.TauntName)
}

func (svcUserMessage *SvcUserMessage) parsePortalFxSurface(reader *bitreader.Reader, demo *types.Demo) {
	type PortalFizzleType int8
	const (
		PortalFizzleSuccess PortalFizzleType = iota // Placed fine (no fizzle)
		PortalFizzleCantFit
		PortalFizzleOverlappedLinked
		PortalFizzleBadVolume
		PortalFizzleBadSurface
		PortalFizzleKilled
		PortalFizzleCleanser
		PortalFizzleClose
		PortalFizzleNearBlue
		PortalFizzleNearRed
		PortalFizzleNone
	)
	getPortalFizzleType := func(portalFizzleType PortalFizzleType) string {
		switch portalFizzleType {
		case PortalFizzleSuccess:
			return "PortalFizzleSuccess"
		case PortalFizzleCantFit:
			return "PortalFizzleCantFit"
		case PortalFizzleOverlappedLinked:
			return "PortalFizzleOverlappedLinked"
		case PortalFizzleBadVolume:
			return "PortalFizzleBadVolume"
		case PortalFizzleBadSurface:
			return "PortalFizzleBadSurface"
		case PortalFizzleKilled:
			return "PortalFizzleKilled"
		case PortalFizzleCleanser:
			return "PortalFizzleCleanser"
		case PortalFizzleClose:
			return "PortalFizzleClose"
		case PortalFizzleNearBlue:
			return "PortalFizzleNearBlue"
		case PortalFizzleNearRed:
			return "PortalFizzleNearRed"
		case PortalFizzleNone:
			return "PortalFizzleNone"
		default:
			return fmt.Sprintf("%d", int(portalFizzleType))
		}
	}
	portalFxSurface := struct {
		PortalEnt uint16    `json:"portal_ent"`
		OwnerEnt  uint16    `json:"owner_ent"`
		Team      uint8     `json:"team"`
		PortalNum uint8     `json:"portal_num"`
		Effect    uint8     `json:"effect"`
		Origin    []float32 `json:"origin"`
		Angles    []float32 `json:"angles"`
	}{
		PortalEnt: reader.TryReadUInt16(),
		OwnerEnt:  reader.TryReadUInt16(),
		Team:      reader.TryReadUInt8(),
		PortalNum: reader.TryReadUInt8(),
		Effect:    reader.TryReadUInt8(),
		Origin:    []float32{},
		Angles:    []float32{},
	}
	existsX, existsY, existsZ := reader.TryReadBool(), reader.TryReadBool(), reader.TryReadBool()
	if existsX {
		portalFxSurface.Origin = append(portalFxSurface.Origin, readBitCoord(reader))
	} else {
		portalFxSurface.Origin = append(portalFxSurface.Origin, 0)
	}
	if existsY {
		portalFxSurface.Origin = append(portalFxSurface.Origin, readBitCoord(reader))
	} else {
		portalFxSurface.Origin = append(portalFxSurface.Origin, 0)
	}
	if existsZ {
		portalFxSurface.Origin = append(portalFxSurface.Origin, readBitCoord(reader))
	} else {
		portalFxSurface.Origin = append(portalFxSurface.Origin, 0)
	}
	existsX, existsY, existsZ = reader.TryReadBool(), reader.TryReadBool(), reader.TryReadBool()
	if existsX {
		portalFxSurface.Angles = append(portalFxSurface.Angles, readBitCoord(reader))
	} else {
		portalFxSurface.Angles = append(portalFxSurface.Angles, 0)
	}
	if existsY {
		portalFxSurface.Angles = append(portalFxSurface.Angles, readBitCoord(reader))
	} else {
		portalFxSurface.Angles = append(portalFxSurface.Angles, 0)
	}
	if existsZ {
		portalFxSurface.Angles = append(portalFxSurface.Angles, readBitCoord(reader))
	} else {
		portalFxSurface.Angles = append(portalFxSurface.Angles, 0)
	}
	svcUserMessage.Data = portalFxSurface
	_ = getPortalFizzleType(PortalFizzleType(2))
	demo.Writer.TempAppendLine("\t\t\tPortal Entity: %d", portalFxSurface.PortalEnt)
	demo.Writer.TempAppendLine("\t\t\tOwner Entity: %d", portalFxSurface.OwnerEnt)
	demo.Writer.TempAppendLine("\t\t\tTeam: %d", portalFxSurface.Team)
	demo.Writer.TempAppendLine("\t\t\tPortal Number: %d", portalFxSurface.PortalNum)
	demo.Writer.TempAppendLine("\t\t\tEffect: %s", getPortalFizzleType(PortalFizzleType(portalFxSurface.Effect)))
	demo.Writer.TempAppendLine("\t\t\tOrigin: %v", portalFxSurface.Origin)
	demo.Writer.TempAppendLine("\t\t\tAngles: %v", portalFxSurface.Angles)
}

func (svcUserMessage *SvcUserMessage) parsePaintWorld(reader *bitreader.Reader, demo *types.Demo) {
	paintWorld := struct {
		Type      uint8       `json:"type"`
		EHandle   uint32      `json:"e_handle"`
		UnkHf1    float32     `json:"unk_hf_1"`
		UnkHf2    float32     `json:"unk_hf_2"`
		Center    []float32   `json:"center"`
		Positions [][]float32 `json:"positions"`
	}{
		Type:    reader.TryReadUInt8(),
		EHandle: reader.TryReadUInt32(),
		UnkHf1:  reader.TryReadFloat32(),
		UnkHf2:  reader.TryReadFloat32(),
		// Center:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		// Positions: [][]float32{{},{}},
	}
	length := reader.TryReadUInt8()
	paintWorld.Center = []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	paintWorld.Positions = make([][]float32, length)
	for i := 0; i < int(length); i++ {
		paintWorld.Positions[i] = []float32{paintWorld.Center[0] + float32(reader.TryReadSInt16()), paintWorld.Center[1] + float32(reader.TryReadSInt16()), paintWorld.Center[2] + float32(reader.TryReadSInt16())}
	}
	svcUserMessage.Data = paintWorld
	demo.Writer.TempAppendLine("\t\t\tType: %d", paintWorld.Type)
	demo.Writer.TempAppendLine("\t\t\tEHandle: %d", paintWorld.EHandle)
	demo.Writer.TempAppendLine("\t\t\tUnkHf1: %f", paintWorld.UnkHf1)
	demo.Writer.TempAppendLine("\t\t\tUnkHf2: %f", paintWorld.UnkHf2)
	demo.Writer.TempAppendLine("\t\t\tCenter: %v", paintWorld.Center)
	demo.Writer.TempAppendLine("\t\t\tPositions: %v", paintWorld.Positions)
}

func (svcUserMessage *SvcUserMessage) parseTransitionFade(reader *bitreader.Reader, demo *types.Demo) {
	transitionFade := struct {
		Seconds float32 `json:"seconds"`
	}{
		Seconds: reader.TryReadFloat32(),
	}
	svcUserMessage.Data = transitionFade
	demo.Writer.TempAppendLine("\t\t\tSeconds: %f", transitionFade.Seconds)
}

func (svcUserMessage *SvcUserMessage) parseScoreboardTempUpdate(reader *bitreader.Reader, demo *types.Demo) {
	scoreboardTempUpdate := struct {
		NumPortals int32 `json:"num_portals"`
		TimeTaken  int32 `json:"time_taken"`
	}{
		NumPortals: reader.TryReadSInt32(),
		TimeTaken:  reader.TryReadSInt32(),
	}
	svcUserMessage.Data = scoreboardTempUpdate
	demo.Writer.TempAppendLine("\t\t\tPortal Count: %d", scoreboardTempUpdate.NumPortals)
	demo.Writer.TempAppendLine("\t\t\tTime Taken: %.2f", float32(scoreboardTempUpdate.TimeTaken)/100.0)

	demo.Writer.TempAppendLine("\t\t\tTicks Taken: %d", int(math.Round(float64((float32(scoreboardTempUpdate.TimeTaken)/100.0)/float32(1.0/60.0)))))
}

type UserMessageType uint8

const (
	EUserMessageTypeGeiger   UserMessageType = iota // done
	EUserMessageTypeTrain                           // done
	EUserMessageTypeHudText                         // done
	EUserMessageTypeSayText                         // done
	EUserMessageTypeSayText2                        // done
	EUserMessageTypeTextMsg                         // done
	EUserMessageTypeHUDMsg                          // done
	EUserMessageTypeResetHUD                        // done // called every respawn
	EUserMessageTypeGameTitle
	EUserMessageTypeItemPickup
	EUserMessageTypeShowMenu
	EUserMessageTypeShake // done
	EUserMessageTypeTilt
	EUserMessageTypeFade      // done
	EUserMessageTypeVGUIMenu  // done // Show VGUI menu
	EUserMessageTypeRumble    // done // Send a rumble to a controller
	EUserMessageTypeBattery   // done
	EUserMessageTypeDamage    // done
	EUserMessageTypeVoiceMask // done
	EUserMessageTypeRequestState
	EUserMessageTypeCloseCaption       // done // Show a caption (by string id number)(duration in 10th of a second)
	EUserMessageTypeCloseCaptionDirect // Show a forced caption (by string id number)(duration in 10th of a second)
	EUserMessageTypeHintText           // Displays hint text display
	EUserMessageTypeKeyHintText        // done // Displays hint text display
	EUserMessageTypeSquadMemberDied
	EUserMessageTypeAmmoDenied
	EUserMessageTypeCreditsMsg
	EUserMessageTypeLogoTimeMsg      // done
	EUserMessageTypeAchievementEvent // done
	EUserMessageTypeUpdateJalopyRadar
	EUserMessageTypeCurrentTimescale // done // Send one float for the new timescale
	EUserMessageTypeDesiredTimescale // Send timescale and some blending vars
	EUserMessageTypeCreditsPortalMsg // portal 1 end
	EUserMessageTypeInventoryFlash   // portal 2 start
	EUserMessageTypeIndicatorFlash
	EUserMessageTypeControlHelperAnimate
	EUserMessageTypeTakePhoto
	EUserMessageTypeFlash
	EUserMessageTypeHudPingIndicator
	EUserMessageTypeOpenRadialMenu
	EUserMessageTypeAddLocator
	EUserMessageTypeMPMapCompleted  // done
	EUserMessageTypeMPMapIncomplete // done
	EUserMessageTypeMPMapCompletedData
	EUserMessageTypeMPTauntEarned // done
	EUserMessageTypeMPTauntUnlocked
	EUserMessageTypeMPTauntLocked // done
	EUserMessageTypeMPAllTauntsLocked
	EUserMessageTypePortalFX_Surface // done
	EUserMessageTypePaintWorld       // d
	EUserMessageTypePaintEntity
	EUserMessageTypeChangePaintColor
	EUserMessageTypePaintBombExplode
	EUserMessageTypeRemoveAllPaint
	EUserMessageTypePaintAllSurfaces
	EUserMessageTypeRemovePaint
	EUserMessageTypeStartSurvey
	EUserMessageTypeApplyHitBoxDamageEffect
	EUserMessageTypeSetMixLayerTriggerFactor
	EUserMessageTypeTransitionFade       // done
	EUserMessageTypeScoreboardTempUpdate // done
	EUserMessageTypeChallengeModCheatSession
	EUserMessageTypeChallengeModCloseAllUI
)

func (userMessageType UserMessageType) String() string {
	switch userMessageType {
	case EUserMessageTypeGeiger:
		return "Geiger"
	case EUserMessageTypeTrain:
		return "Train"
	case EUserMessageTypeHudText:
		return "HudText"
	case EUserMessageTypeSayText:
		return "SayText"
	case EUserMessageTypeSayText2:
		return "SayText2"
	case EUserMessageTypeTextMsg:
		return "TextMsg"
	case EUserMessageTypeHUDMsg:
		return "HUDMsg"
	case EUserMessageTypeResetHUD:
		return "ResetHUD"
	case EUserMessageTypeGameTitle:
		return "GameTitle"
	case EUserMessageTypeItemPickup:
		return "ItemPickup"
	case EUserMessageTypeShowMenu:
		return "ShowMenu"
	case EUserMessageTypeShake:
		return "Shake"
	case EUserMessageTypeTilt:
		return "Tilt"
	case EUserMessageTypeFade:
		return "Fade"
	case EUserMessageTypeVGUIMenu:
		return "VGUIMenu"
	case EUserMessageTypeRumble:
		return "Rumble"
	case EUserMessageTypeBattery:
		return "Battery"
	case EUserMessageTypeDamage:
		return "Damage"
	case EUserMessageTypeVoiceMask:
		return "VoiceMask"
	case EUserMessageTypeRequestState:
		return "RequestState"
	case EUserMessageTypeCloseCaption:
		return "CloseCaption"
	case EUserMessageTypeCloseCaptionDirect:
		return "CloseCaptionDirect"
	case EUserMessageTypeHintText:
		return "HintText"
	case EUserMessageTypeKeyHintText:
		return "KeyHintText"
	case EUserMessageTypeSquadMemberDied:
		return "SquadMemberDied"
	case EUserMessageTypeAmmoDenied:
		return "AmmoDenied"
	case EUserMessageTypeCreditsMsg:
		return "CreditsMsg"
	case EUserMessageTypeLogoTimeMsg:
		return "LogoTimeMsg"
	case EUserMessageTypeAchievementEvent:
		return "AchievementEvent"
	case EUserMessageTypeUpdateJalopyRadar:
		return "UpdateJalopyRadar"
	case EUserMessageTypeCurrentTimescale:
		return "CurrentTimescale"
	case EUserMessageTypeDesiredTimescale:
		return "DesiredTimescale"
	case EUserMessageTypeCreditsPortalMsg:
		return "CreditsPortalMsg"
	case EUserMessageTypeInventoryFlash:
		return "InventoryFlash"
	case EUserMessageTypeIndicatorFlash:
		return "IndicatorFlash"
	case EUserMessageTypeControlHelperAnimate:
		return "ControlHelperAnimate"
	case EUserMessageTypeTakePhoto:
		return "TakePhoto"
	case EUserMessageTypeFlash:
		return "Flash"
	case EUserMessageTypeHudPingIndicator:
		return "HudPingIndicator"
	case EUserMessageTypeOpenRadialMenu:
		return "OpenRadialMenu"
	case EUserMessageTypeAddLocator:
		return "AddLocator"
	case EUserMessageTypeMPMapCompleted:
		return "MPMapCompleted"
	case EUserMessageTypeMPMapIncomplete:
		return "MPMapIncomplete"
	case EUserMessageTypeMPMapCompletedData:
		return "MPMapCompletedData"
	case EUserMessageTypeMPTauntEarned:
		return "MPTauntEarned"
	case EUserMessageTypeMPTauntUnlocked:
		return "MPTauntUnlocked"
	case EUserMessageTypeMPTauntLocked:
		return "MPTauntLocked"
	case EUserMessageTypeMPAllTauntsLocked:
		return "MPAllTauntsLocked"
	case EUserMessageTypePortalFX_Surface:
		return "PortalFX_Surface"
	case EUserMessageTypePaintWorld:
		return "PaintWorld"
	case EUserMessageTypePaintEntity:
		return "PaintEntity"
	case EUserMessageTypeChangePaintColor:
		return "ChangePaintColor"
	case EUserMessageTypePaintBombExplode:
		return "PaintBombExplode"
	case EUserMessageTypeRemoveAllPaint:
		return "RemoveAllPaint"
	case EUserMessageTypePaintAllSurfaces:
		return "PaintAllSurfaces"
	case EUserMessageTypeRemovePaint:
		return "RemovePaint"
	case EUserMessageTypeStartSurvey:
		return "StartSurvey"
	case EUserMessageTypeApplyHitBoxDamageEffect:
		return "ApplyHitBoxDamageEffect"
	case EUserMessageTypeSetMixLayerTriggerFactor:
		return "SetMixLayerTriggerFactor"
	case EUserMessageTypeTransitionFade:
		return "TransitionFade"
	case EUserMessageTypeScoreboardTempUpdate:
		return "ScoreboardTempUpdate"
	case EUserMessageTypeChallengeModCheatSession:
		return "ChallengeModCheatSession"
	case EUserMessageTypeChallengeModCloseAllUI:
		return "ChallengeModCloseAllUI"
	default:
		return "Unknown"
	}
}

func readBitCoord(reader *bitreader.Reader) float32 {
	const (
		CoordIntBits  uint64 = 14
		CoordFracBits uint64 = 5
		CoordDenom           = 1 << CoordFracBits
		CoordRes             = 1.0 / CoordDenom
	)
	val := float32(0)
	hasInt := reader.TryReadBool()
	hasFrac := reader.TryReadBool()
	if hasInt || hasFrac {
		sign := reader.TryReadBool()
		if hasInt {
			val += float32(reader.TryReadBits(CoordIntBits) + 1)
		}
		if hasFrac {
			val += float32(reader.TryReadBits(CoordFracBits)) * CoordRes
		}
		if sign {
			val = -val
		}
	}
	return val
}
