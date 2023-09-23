package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcUserMessage struct {
	Type   int8
	Length int16
	Data   any
}

func ParseSvcUserMessage(reader *bitreader.Reader) SvcUserMessage {
	svcUserMessage := SvcUserMessage{
		Type:   int8(reader.TryReadBits(8)),
		Length: int16(reader.TryReadBits(12)),
	}
	svcUserMessage.Data = reader.TryReadBitsToSlice(uint64(svcUserMessage.Length))
	userMessageReader := bitreader.NewReaderFromBytes(svcUserMessage.Data.([]byte), true)
	writer.TempAppendLine("\t\t%s (%d):", UserMessageType(svcUserMessage.Type).String(), svcUserMessage.Type)
	switch UserMessageType(svcUserMessage.Type) {
	case EUserMessageTypeGeiger:
		svcUserMessage.parseGeiger(userMessageReader)
	case EUserMessageTypeTrain:
		svcUserMessage.parseTrain(userMessageReader)
	case EUserMessageTypeHudText:
		svcUserMessage.parseHUDText(userMessageReader)
	case EUserMessageTypeSayText:
		svcUserMessage.parseSayText(userMessageReader)
	case EUserMessageTypeSayText2:
		svcUserMessage.parseSayText2(userMessageReader)
	case EUserMessageTypeTextMsg:
		svcUserMessage.parseTextMsg(userMessageReader)
	case EUserMessageTypeHUDMsg:
		svcUserMessage.parseHUDMsg(userMessageReader)
	case EUserMessageTypeResetHUD:
		svcUserMessage.parseResetHUD(userMessageReader)
	case EUserMessageTypeShake:
		svcUserMessage.parseShake(userMessageReader)
	case EUserMessageTypeFade:
		svcUserMessage.parseFade(userMessageReader)
	case EUserMessageTypeVGUIMenu:
		svcUserMessage.parseVguiMenu(userMessageReader)
	case EUserMessageTypeRumble:
		svcUserMessage.parseRumble(userMessageReader)
	case EUserMessageTypeBattery:
		svcUserMessage.parseBattery(userMessageReader)
	case EUserMessageTypeDamage:
		svcUserMessage.parseDamage(userMessageReader)
	case EUserMessageTypeVoiceMask:
		svcUserMessage.parseVoiceMask(userMessageReader)
	case EUserMessageTypeCloseCaption:
		svcUserMessage.parseCloseCaption(userMessageReader)
	case EUserMessageTypeKeyHintText:
		svcUserMessage.parseKeyHintText(userMessageReader)
	case EUserMessageTypeLogoTimeMsg:
		svcUserMessage.parseLogoTimeMsg(userMessageReader)
	case EUserMessageTypeAchievementEvent:
		svcUserMessage.parseAchivementEvent(userMessageReader)
	case EUserMessageTypeMPMapCompleted:
		svcUserMessage.parseMpMapCompleted(userMessageReader)
	case EUserMessageTypeMPMapIncomplete:
		svcUserMessage.parseMpMapIncomplete(userMessageReader)
	case EUserMessageTypeMPTauntEarned:
		svcUserMessage.parseMpTauntEarned(userMessageReader)
	case EUserMessageTypeMPTauntLocked:
		svcUserMessage.parseMpTauntLocked(userMessageReader)
	case EUserMessageTypePortalFX_Surface:
		svcUserMessage.parsePortalFxSurface(userMessageReader)
	case EUserMessageTypeScoreboardTempUpdate:
		svcUserMessage.parseScoreboardTempUpdate(userMessageReader)
	default:
		writer.TempAppendLine("\t\t\tData: %v", svcUserMessage.Data)
	}
	return svcUserMessage
}

func (svcUserMessage *SvcUserMessage) parseGeiger(reader *bitreader.Reader) {
	geiger := struct{ Range uint8 }{
		Range: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = geiger
	writer.TempAppendLine("\t\t\tGeiger Range: %d", geiger.Range)
}

func (svcUserMessage *SvcUserMessage) parseTrain(reader *bitreader.Reader) {
	train := struct{ Pos uint8 }{
		Pos: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = train
	writer.TempAppendLine("\t\t\tPos: %d", train.Pos)
}

func (svcUserMessage *SvcUserMessage) parseHUDText(reader *bitreader.Reader) {
	hudText := struct{ Text string }{
		Text: reader.TryReadString(),
	}
	svcUserMessage.Data = hudText
	writer.TempAppendLine("\t\t\tText: %s", hudText.Text)
}

func (svcUserMessage *SvcUserMessage) parseSayText(reader *bitreader.Reader) {
	sayText := struct {
		Client      uint8
		Message     string
		WantsToChat bool
	}{
		Client:      reader.TryReadUInt8(),
		Message:     reader.TryReadString(),
		WantsToChat: reader.TryReadUInt8() != 0,
	}
	svcUserMessage.Data = sayText
	writer.TempAppendLine("\t\t\tClient: %d", sayText.Client)
	writer.TempAppendLine("\t\t\tMessage: %s", sayText.Message)
	writer.TempAppendLine("\t\t\tWants To Chat: %t", sayText.WantsToChat)
}

func (svcUserMessage *SvcUserMessage) parseSayText2(reader *bitreader.Reader) {
	sayText2 := struct {
		Client      uint8
		WantsToChat bool
		MessageName string
		Messages    []string
	}{
		Client:      reader.TryReadUInt8(),
		WantsToChat: reader.TryReadUInt8() != 0,
		MessageName: reader.TryReadString(),
		Messages:    []string{reader.TryReadString(), reader.TryReadString(), reader.TryReadString()},
	}
	svcUserMessage.Data = sayText2
	writer.TempAppendLine("\t\t\tClient: %d", sayText2.Client)
	writer.TempAppendLine("\t\t\tWants To Chat: %t", sayText2.WantsToChat)
	writer.TempAppendLine("\t\t\tName: %s", sayText2.MessageName)
	for index, message := range sayText2.Messages {
		writer.TempAppendLine("\t\t\tMessage %d: %s", index, message)
	}
}

func (svcUserMessage *SvcUserMessage) parseTextMsg(reader *bitreader.Reader) {
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
	writer.TempAppendLine("\t\t\tDestination: %d", textMsg.Destination)
	for i := 0; i < MessageCount; i++ {
		writer.TempAppendLine("\t\t\tMessage %d: %s", i+1, textMsg.Messages)
	}
}

func (svcUserMessage *SvcUserMessage) parseHUDMsg(reader *bitreader.Reader) {
	const MaxNetMessage uint8 = 6
	hudMsg := struct {
		Channel uint8
		Info    struct {
			X, Y                              float32 // 0-1 & resolution independent, -1 means center in each dimension
			R1, G1, B1, A1                    uint8
			R2, G2, B2, A2                    uint8
			Effect                            uint8
			FadeIn, FadeOut, HoldTime, FxTime float32 // the fade times seem to be per character
			Message                           string
		}
	}{
		Channel: reader.TryReadUInt8() % MaxNetMessage,
	}
	svcUserMessage.Data = hudMsg
	writer.TempAppendLine("\t\t\tChannel: %d", hudMsg.Channel)
	if reader.TryReadRemainingBits() >= 148 {
		hudMsg.Info = struct {
			X        float32
			Y        float32
			R1       uint8
			G1       uint8
			B1       uint8
			A1       uint8
			R2       uint8
			G2       uint8
			B2       uint8
			A2       uint8
			Effect   uint8
			FadeIn   float32
			FadeOut  float32
			HoldTime float32
			FxTime   float32
			Message  string
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
		writer.TempAppendLine("\t\t\tX: %f, Y: %f", hudMsg.Info.X, hudMsg.Info.Y)
		writer.TempAppendLine("\t\t\tRGBA1: %3d %3d %3d %3d", hudMsg.Info.R1, hudMsg.Info.G1, hudMsg.Info.B1, hudMsg.Info.A1)
		writer.TempAppendLine("\t\t\tRGBA2: %3d %3d %3d %3d", hudMsg.Info.R2, hudMsg.Info.G2, hudMsg.Info.B2, hudMsg.Info.A2)
		writer.TempAppendLine("\t\t\tEffect: %d", hudMsg.Info.Effect)
		writer.TempAppendLine("\t\t\tFade In: %f", hudMsg.Info.FadeIn)
		writer.TempAppendLine("\t\t\tFade Out: %f", hudMsg.Info.FadeOut)
		writer.TempAppendLine("\t\t\tHold Time: %f", hudMsg.Info.HoldTime)
		writer.TempAppendLine("\t\t\tFX Time: %f", hudMsg.Info.FxTime)
		writer.TempAppendLine("\t\t\tMessage: %s", hudMsg.Info.Message)
	}
}

func (svcUserMessage *SvcUserMessage) parseResetHUD(reader *bitreader.Reader) {
	resetHUD := struct{ Unknown uint8 }{
		Unknown: reader.TryReadUInt8(),
	}
	svcUserMessage.Data = resetHUD
	writer.TempAppendLine("\t\t\tUnknown: %d", resetHUD.Unknown)
}

func (svcUserMessage *SvcUserMessage) parseShake(reader *bitreader.Reader) {
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
		Command   uint8
		Amplitude float32
		Frequency float32
		Duration  float32
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
	writer.TempAppendLine("\t\t\tCommand: %v", shakeCommandToString(ShakeCommand(shake.Command)))
	writer.TempAppendLine("\t\t\tAmplitude: %v", shake.Amplitude)
	writer.TempAppendLine("\t\t\tFrequency: %v", shake.Frequency)
	writer.TempAppendLine("\t\t\tDuration: %v", shake.Duration)
}

func (svcUserMessage *SvcUserMessage) parseFade(reader *bitreader.Reader) {
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
		Duration float32
		HoldTime uint16
		Flags    uint16
		R        uint8
		G        uint8
		B        uint8
		A        uint8
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
	writer.TempAppendLine("\t\t\tDuration: %f", fade.Duration)
	writer.TempAppendLine("\t\t\tHold Time: %d", fade.HoldTime)
	writer.TempAppendLine("\t\t\tFlags: %v", getFlags(FadeFlag(fade.Flags)))
	writer.TempAppendLine("\t\t\tRGBA: %3d %3d %3d %3d", fade.R, fade.G, fade.B, fade.A)
}

func (svcUserMessage *SvcUserMessage) parseVguiMenu(reader *bitreader.Reader) {
	vguiMenu := struct {
		Message   string
		Show      bool
		KeyValues []map[string]string
	}{
		Message: reader.TryReadString(),
		Show:    reader.TryReadUInt8() != 0,
	}
	count := reader.TryReadUInt8()
	for i := 0; i < int(count); i++ {
		vguiMenu.KeyValues = append(vguiMenu.KeyValues, map[string]string{"Key": reader.TryReadString(), "Value": reader.TryReadString()})
	}
	svcUserMessage.Data = vguiMenu
	writer.TempAppendLine("\t\t\tMessage: %s", vguiMenu.Message)
	writer.TempAppendLine("\t\t\tShow: %t", vguiMenu.Show)
	if len(vguiMenu.KeyValues) > 0 {
		writer.TempAppendLine("\t\t\t%d Key Value Pairs:", len(vguiMenu.KeyValues))
		for _, kv := range vguiMenu.KeyValues {
			writer.TempAppendLine("\t\t\t\t%s: %s", kv["Key"], kv["Value"])
		}
	} else {
		writer.TempAppendLine("\t\t\tNo Key Value Pairs")
	}
}

func (svcUserMessage *SvcUserMessage) parseRumble(reader *bitreader.Reader) {
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
		Type  int8
		Scale float32
		Flags uint8
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
	writer.TempAppendLine("\t\t\tType: %s", getRumbleLookup(RumbleLookup(rumble.Type)))
	writer.TempAppendLine("\t\t\tScale: %f", rumble.Scale)
	writer.TempAppendLine("\t\t\tFlags: %v", getFlags(RumbleFlag(rumble.Flags)))
}

func (svcUserMessage *SvcUserMessage) parseBattery(reader *bitreader.Reader) {
	battery := struct{ BatteryVal uint16 }{
		BatteryVal: reader.TryReadUInt16(),
	}
	svcUserMessage.Data = battery
	writer.TempAppendLine("\t\t\tBattery: %d", battery.BatteryVal)
}

func (svcUserMessage *SvcUserMessage) parseDamage(reader *bitreader.Reader) {
	damage := struct {
		Armor       uint8
		DamageTaken uint8
		BitsDamage  int32
		VecFrom     []float32
	}{
		Armor:       reader.TryReadUInt8(),
		DamageTaken: reader.TryReadUInt8(),
		BitsDamage:  reader.TryReadSInt32(),
		VecFrom:     []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	}
	svcUserMessage.Data = damage
	writer.TempAppendLine("\t\t\tArmor: %d", damage.Armor)
	writer.TempAppendLine("\t\t\tDamage Taken: %d", damage.DamageTaken)
	writer.TempAppendLine("\t\t\tBits Damage: %d", damage.BitsDamage)
	writer.TempAppendLine("\t\t\tVecFrom: %v", damage.VecFrom)
}

func (svcUserMessage *SvcUserMessage) parseVoiceMask(reader *bitreader.Reader) {
	// const VoiceMaxPlayers = 2
	voiceMask := struct {
		PlayerMasks []struct {
			GameRulesMask int32
			BanMask       int32
		}
		PlayerModEnable bool
	}{
		PlayerMasks: []struct {
			GameRulesMask int32
			BanMask       int32
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
	writer.TempAppendLine("\t\t\tPlayer Masks:")
	writer.TempAppendLine("\t\t\t\t[0] Game Rules Mask: %d", voiceMask.PlayerMasks[0].GameRulesMask)
	writer.TempAppendLine("\t\t\t\t[0] Ban Mask: %d", voiceMask.PlayerMasks[0].BanMask)
	writer.TempAppendLine("\t\t\t\t[1] Game Rules Mask: %d", voiceMask.PlayerMasks[1].GameRulesMask)
	writer.TempAppendLine("\t\t\t\t[1] Ban Mask: %d", voiceMask.PlayerMasks[1].BanMask)
	writer.TempAppendLine("\t\t\t\tPlayer Mod Enable: %t", voiceMask.PlayerModEnable)
}

func (svcUserMessage *SvcUserMessage) parseCloseCaption(reader *bitreader.Reader) {
	type CloseCaptionFlag uint8
	const (
		None          CloseCaptionFlag = 0
		WarnIfMissing CloseCaptionFlag = 1
		FromPlayer    CloseCaptionFlag = 1 << 1
		GenderMale    CloseCaptionFlag = 1 << 2
		GenderFemale  CloseCaptionFlag = 1 << 3
	)
	closeCaption := struct {
		TokenName string
		Duration  float32
		Flags     uint8
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
	writer.TempAppendLine("\t\t\tToken Name: %s", closeCaption.TokenName)
	writer.TempAppendLine("\t\t\tDuration: %f", closeCaption.Duration)
	writer.TempAppendLine("\t\t\tFlags: %v", getFlags(CloseCaptionFlag(closeCaption.Flags)))
}

func (svcUserMessage *SvcUserMessage) parseKeyHintText(reader *bitreader.Reader) {
	keyHintText := struct {
		Count     uint8
		KeyString string
	}{
		Count:     reader.TryReadUInt8(),
		KeyString: reader.TryReadString(),
	}
	svcUserMessage.Data = keyHintText
	writer.TempAppendLine("\t\t\tCount: %d", keyHintText.Count)
	writer.TempAppendLine("\t\t\tString: %s", keyHintText.KeyString)
}

func (svcUserMessage *SvcUserMessage) parseLogoTimeMsg(reader *bitreader.Reader) {
	logoTimeMsg := struct{ Time float32 }{
		Time: reader.TryReadFloat32(),
	}
	svcUserMessage.Data = logoTimeMsg
	writer.TempAppendLine("\t\t\tTime: %f", logoTimeMsg.Time)
}

func (svcUserMessage *SvcUserMessage) parseAchivementEvent(reader *bitreader.Reader) {
	achivementEvent := struct{ AchivementID int32 }{
		AchivementID: reader.TryReadSInt32(),
	}
	svcUserMessage.Data = achivementEvent
	writer.TempAppendLine("\t\t\tPortal Count: %v", achivementEvent.AchivementID)
}

func (svcUserMessage *SvcUserMessage) parseMpMapCompleted(reader *bitreader.Reader) {
	mpMapCompleted := struct {
		Branch uint8
		Level  uint8
	}{
		Branch: reader.TryReadUInt8(),
		Level:  reader.TryReadUInt8(),
	}
	svcUserMessage.Data = mpMapCompleted
	writer.TempAppendLine("\t\t\tBranch: %d", mpMapCompleted.Branch)
	writer.TempAppendLine("\t\t\tLevel: %d", mpMapCompleted.Level)
}

func (svcUserMessage *SvcUserMessage) parseMpMapIncomplete(reader *bitreader.Reader) {}

func (svcUserMessage *SvcUserMessage) parseMpTauntEarned(reader *bitreader.Reader) {
	mpTauntEarned := struct {
		TauntName     string
		AwardSilently bool
	}{
		TauntName:     reader.TryReadString(),
		AwardSilently: reader.TryReadBool(),
	}
	svcUserMessage.Data = mpTauntEarned
	writer.TempAppendLine("\t\t\tTaunt Name: %s", mpTauntEarned.TauntName)
	writer.TempAppendLine("\t\t\tAward Silently: %t", mpTauntEarned.AwardSilently)
}

func (svcUserMessage *SvcUserMessage) parseMpTauntLocked(reader *bitreader.Reader) {
	mpTauntLocked := struct{ TauntName string }{
		TauntName: reader.TryReadString(),
	}
	svcUserMessage.Data = mpTauntLocked
	writer.TempAppendLine("\t\t\tTaunt Name: %s", mpTauntLocked.TauntName)
}

func (svcUserMessage *SvcUserMessage) parsePortalFxSurface(reader *bitreader.Reader) {
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
		PortalEnt uint16
		OwnerEnt  uint16
		Team      uint8
		PortalNum uint8
		Effect    uint8
		Origin    []float32
		Angles    []float32
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
	writer.TempAppendLine("\t\t\tPortal Entity: %d", portalFxSurface.PortalEnt)
	writer.TempAppendLine("\t\t\tOwner Entity: %d", portalFxSurface.OwnerEnt)
	writer.TempAppendLine("\t\t\tTeam: %d", portalFxSurface.Team)
	writer.TempAppendLine("\t\t\tPortal Number: %d", portalFxSurface.PortalNum)
	writer.TempAppendLine("\t\t\tEffect: %s", getPortalFizzleType(PortalFizzleType(portalFxSurface.Effect)))
	writer.TempAppendLine("\t\t\tOrigin: %v", portalFxSurface.Origin)
	writer.TempAppendLine("\t\t\tAngles: %v", portalFxSurface.Angles)
}

func (svcUserMessage *SvcUserMessage) parseScoreboardTempUpdate(reader *bitreader.Reader) {
	scoreboardTempUpdate := struct {
		NumPortals int32
		TimeTaken  int32
	}{
		NumPortals: reader.TryReadSInt32(),
		TimeTaken:  reader.TryReadSInt32(),
	}
	svcUserMessage.Data = scoreboardTempUpdate
	writer.TempAppendLine("\t\t\tPortal Count: %v", scoreboardTempUpdate.NumPortals)
	writer.TempAppendLine("\t\t\tCM Ticks: %v", scoreboardTempUpdate.TimeTaken)
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
	EUserMessageTypeCurrentTimescale // Send one float for the new timescale
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
	EUserMessageTypePaintWorld
	EUserMessageTypePaintEntity
	EUserMessageTypeChangePaintColor
	EUserMessageTypePaintBombExplode
	EUserMessageTypeRemoveAllPaint
	EUserMessageTypePaintAllSurfaces
	EUserMessageTypeRemovePaint
	EUserMessageTypeStartSurvey
	EUserMessageTypeApplyHitBoxDamageEffect
	EUserMessageTypeSetMixLayerTriggerFactor
	EUserMessageTypeTransitionFade
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
