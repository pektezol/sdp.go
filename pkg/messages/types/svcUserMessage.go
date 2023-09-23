package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcUserMessage struct {
	Type   int8
	Length int16
	Data   []byte
}

type UserMessageType int

const (
	EUserMessageTypeUnknown UserMessageType = iota
	EUserMessageTypeInvalid
	EUserMessageTypeGeiger
	EUserMessageTypeTrain
	EUserMessageTypeHudText
	EUserMessageTypeSayText
	EUserMessageTypeSayText2
	EUserMessageTypeTextMsg
	EUserMessageTypeHUDMsg
	EUserMessageTypeResetHUD
	EUserMessageTypeGameTitle
	EUserMessageTypeItemPickup
	EUserMessageTypeShowMenu
	EUserMessageTypeShake
	EUserMessageTypeFade
	EUserMessageTypeVGUIMenu
	EUserMessageTypeRumble
	EUserMessageTypeBattery
	EUserMessageTypeDamage
	EUserMessageTypeVoiceMask
	EUserMessageTypeRequestState
	EUserMessageTypeCloseCaption
	EUserMessageTypeHintText
	EUserMessageTypeKeyHintText
	EUserMessageTypeSquadMemberDied
	EUserMessageTypeAmmoDenied
	EUserMessageTypeCreditsMsg
	EUserMessageTypeCreditsPortalMsg
	EUserMessageTypeLogoTimeMsg
	EUserMessageTypeAchievementEvent
	EUserMessageTypeEntityPortalled
	EUserMessageTypeKillCam
	EUserMessageTypeTilt
	EUserMessageTypeCloseCaptionDirect
	EUserMessageTypeUpdateJalopyRadar
	EUserMessageTypeCurrentTimescale
	EUserMessageTypeDesiredTimescale
	EUserMessageTypeInventoryFlash
	EUserMessageTypeIndicatorFlash
	EUserMessageTypeControlHelperAnimate
	EUserMessageTypeTakePhoto
	EUserMessageTypeFlash
	EUserMessageTypeHudPingIndicator
	EUserMessageTypeOpenRadialMenu
	EUserMessageTypeAddLocator
	EUserMessageTypeMPMapCompleted
	EUserMessageTypeMPMapIncomplete
	EUserMessageTypeMPMapCompletedData
	EUserMessageTypeMPTauntEarned
	EUserMessageTypeMPTauntUnlocked
	EUserMessageTypeMPTauntLocked
	EUserMessageTypeMPAllTauntsLocked
	EUserMessageTypePortalFX_Surface
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
	EUserMessageTypeScoreboardTempUpdate
	EUserMessageTypeChallengeModCheatSession
	EUserMessageTypeChallengeModCloseAllUI
)

func ParseSvcUserMessage(reader *bitreader.Reader) SvcUserMessage {
	svcUserMessage := SvcUserMessage{
		Type:   int8(reader.TryReadBits(8)),
		Length: int16(reader.TryReadBits(12)),
	}
	svcUserMessage.Data = reader.TryReadBitsToSlice(uint64(svcUserMessage.Length))

	return svcUserMessage
}

// func byteToUserMessageType() {

// }
