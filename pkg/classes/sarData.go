package classes

import (
	"errors"
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SarDataType uint8

const (
	ESarDataTimescaleCheat  SarDataType = 0x01
	ESarDataInitialCVar     SarDataType = 0x02
	ESarDataEntityInput     SarDataType = 0x03
	ESarDataEntityInputSlot SarDataType = 0x04
	ESarDataPortalPlacement SarDataType = 0x05
	ESarDataChallengeFlags  SarDataType = 0x06
	ESarDataCrouchFly       SarDataType = 0x07
	ESarDataPause           SarDataType = 0x08
	ESarDataWaitRun         SarDataType = 0x09
	ESarDataSpeedrunTime    SarDataType = 0x0A
	ESarDataTimestamp       SarDataType = 0x0B
	ESarDataFileChecksum    SarDataType = 0x0C
	ESarDataHWaitRun        SarDataType = 0x0D
	ESarDataChecksum        SarDataType = 0xFF
	ESarDataChecksumV2      SarDataType = 0xFE
	ESarDataInvalid         SarDataType = iota
)

func (sarDataType SarDataType) String() string {
	switch sarDataType {
	case ESarDataTimescaleCheat:
		return "SarDataTimescaleCheat"
	case ESarDataInitialCVar:
		return "SarDataInitialCVar"
	case ESarDataEntityInput:
		return "SarDataEntityInput"
	case ESarDataEntityInputSlot:
		return "SarDataEntityInputSlot"
	case ESarDataPortalPlacement:
		return "SarDataPortalPlacement"
	case ESarDataChallengeFlags:
		return "SarDataChallengeFlags"
	case ESarDataCrouchFly:
		return "SarDataCrouchFly"
	case ESarDataPause:
		return "SarDataPause"
	case ESarDataWaitRun:
		return "SarDataWaitRun"
	case ESarDataSpeedrunTime:
		return "SarDataSpeedrunTime"
	case ESarDataTimestamp:
		return "SarDataTimestamp"
	case ESarDataFileChecksum:
		return "SarDataFileChecksum"
	case ESarDataHWaitRun:
		return "SarDataHWaitRun"
	case ESarDataChecksum:
		return "SarDataChecksum"
	case ESarDataChecksumV2:
		return "SarDataChecksumV2"
	case ESarDataInvalid:
		return "SarDataInvalid"
	default:
		return fmt.Sprintf("%d", int(sarDataType))
	}
}

type SarData struct {
	Type SarDataType `json:"type"`
	Slot int         `json:"slot"`
	Data any         `json:"data"`
}

type SarDataTimescaleCheat struct {
	Timescale float32 `json:"timescale"`
}

type SarDataInitialCVar struct {
	CVar string `json:"cvar"`
	Val  string `json:"val"`
}

type SarDataChecksum struct {
	DemoSum uint32 `json:"demo_sum"`
	SarSum  uint32 `json:"sar_sum"`
}

type SarDataChecksumV2 struct {
	SarSum    uint32   `json:"sar_sum"`
	Signature [64]byte `json:"signature"`
}

type SarDataEntityInput struct {
	TargetName string `json:"target_name"`
	ClassName  string `json:"class_name"`
	InputName  string `json:"input_name"`
	Parameter  string `json:"parameter"`
}

type SarDataPortalPlacement struct {
	Orange bool    `json:"orange"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Z      float32 `json:"z"`
}

type SarDataPause struct {
	PauseTicks uint32 `json:"pause_ticks"`
}

type SarDataWaitRun struct {
	Ticks int    `json:"ticks"`
	Cmd   string `json:"cmd"`
}

type SarDataHWaitRun struct {
	Ticks int    `json:"ticks"`
	Cmd   string `json:"cmd"`
}

type SarDataSpeedrunTime struct {
	NSplits uint32                      `json:"n_splits"`
	Splits  []SarDataSpeedrunTimeSplits `json:"splits"`
}

type SarDataSpeedrunTimeSegs struct {
	Name  string `json:"name"`
	Ticks uint32 `json:"ticks"`
}

type SarDataSpeedrunTimeSplits struct {
	Name  string                    `json:"name"`
	NSegs uint32                    `json:"n_segs"`
	Segs  []SarDataSpeedrunTimeSegs `json:"segs"`
}

type SarDataTimestamp struct {
	Year   uint16 `json:"year"`
	Month  uint8  `json:"month"`
	Day    uint8  `json:"day"`
	Hour   uint8  `json:"hour"`
	Minute uint8  `json:"minute"`
	Second uint8  `json:"second"`
}

type SarDataFileChecksum struct {
	Sum  uint32 `json:"sum"`
	Path string `json:"path"`
}

func (sarData *SarData) ParseSarData(reader *bitreader.Reader, demo *types.Demo) (err error) {
	reader.SkipBytes(8)
	len := reader.TryReadRemainingBits() / 8
	if len == 0 {
		sarData.Type = ESarDataInvalid
		err = errors.New("sar data invalid")
		return err
	}
	sarData.Type = SarDataType(reader.TryReadBytes(1))
	if sarData.Type == ESarDataChecksum && len == 5 {
		len = 9
	}
	dataReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(len-1), true)
	demo.Writer.AppendLine("\tMessage: %s (%d):", sarData.Type.String(), sarData.Type)
	switch sarData.Type {
	case ESarDataTimescaleCheat:
		sarData.Data, err = parseTimescaleCheatData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataInitialCVar:
		sarData.Data = parseInitialCVarData(dataReader, demo)
	case ESarDataEntityInputSlot:
		sarData.Slot = int(dataReader.TryReadBytes(1))
		demo.Writer.AppendLine("\t\tSlot: %d", sarData.Slot)
	case ESarDataEntityInput:
		sarData.Data = parseEntityInputData(dataReader, demo)
	case ESarDataChecksum:
		sarData.Data, err = parseChecksumData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataChecksumV2:
		sarData.Data, err = parseChecksumV2Data(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataPortalPlacement:
		data, slot, err := parsePortalPlacementData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		} else {
			sarData.Data = data
			sarData.Slot = slot
		}
	case ESarDataChallengeFlags, ESarDataCrouchFly:
		sarData.Slot, err = parseChallengeFlagsCrouchFlyData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
		demo.Writer.AppendLine("\t\tSlot: %d", sarData.Slot)
	case ESarDataPause:
		sarData.Data, err = parsePauseData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataWaitRun:
		sarData.Data, err = parseWaitRunData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataHWaitRun:
		sarData.Data, err = parseHWaitRunData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataSpeedrunTime:
		sarData.Data, err = parseSpeedrunTimeData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataTimestamp:
		sarData.Data, err = parseTimestampData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataFileChecksum:
		sarData.Data, err = parseFileChecksumData(dataReader, len, demo)
		if err != nil {
			sarData.Data = nil
		}
	default:
		err = errors.New("unsupported SAR data type")
		return err
	}
	return nil
}

func parseTimescaleCheatData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataTimescaleCheat, error) {
	if length != 5 {
		return SarDataTimescaleCheat{}, errors.New("sar data invalid")
	}
	sarDataTimescaleCheat := SarDataTimescaleCheat{
		Timescale: reader.TryReadFloat32(),
	}
	demo.Writer.AppendLine("\t\tTimescale: %f", sarDataTimescaleCheat.Timescale)
	return sarDataTimescaleCheat, nil
}

func parseInitialCVarData(reader *bitreader.Reader, demo *types.Demo) SarDataInitialCVar {
	sarDataInitialCvar := SarDataInitialCVar{
		CVar: reader.TryReadString(),
		Val:  reader.TryReadString(),
	}
	demo.Writer.AppendLine("\t\tCvar: \"%s\" = \"%s\"", sarDataInitialCvar.CVar, sarDataInitialCvar.Val)
	return sarDataInitialCvar
}

func parseEntityInputData(reader *bitreader.Reader, demo *types.Demo) SarDataEntityInput {
	sarDataEntityInput := SarDataEntityInput{
		TargetName: reader.TryReadString(),
		ClassName:  reader.TryReadString(),
		InputName:  reader.TryReadString(),
		Parameter:  reader.TryReadString(),
	}
	demo.Writer.AppendLine("\t\tTarget: %s", sarDataEntityInput.TargetName)
	demo.Writer.AppendLine("\t\tClass: %s", sarDataEntityInput.ClassName)
	demo.Writer.AppendLine("\t\tInput: %s", sarDataEntityInput.InputName)
	demo.Writer.AppendLine("\t\tParameter: %s", sarDataEntityInput.Parameter)
	return sarDataEntityInput
}

func parseChecksumData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataChecksum, error) {
	if length != 9 {
		return SarDataChecksum{}, errors.New("sar data invalid")
	}
	sarDataChecksum := SarDataChecksum{
		DemoSum: reader.TryReadUInt32(),
		SarSum:  reader.TryReadUInt32(),
	}
	demo.Writer.AppendLine("\t\tDemo Checksum: %d", sarDataChecksum.DemoSum)
	demo.Writer.AppendLine("\t\tSAR Checksum: %d", sarDataChecksum.SarSum)
	return sarDataChecksum, nil
}

func parseChecksumV2Data(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataChecksumV2, error) {
	if length != 69 {
		return SarDataChecksumV2{}, errors.New("sar data invalid")
	}
	sarDataChecksumV2 := SarDataChecksumV2{
		SarSum:    reader.TryReadUInt32(),
		Signature: [64]byte(reader.TryReadBytesToSlice(60)),
	}
	demo.Writer.AppendLine("\t\tSAR Checksum: %d", sarDataChecksumV2.SarSum)
	demo.Writer.AppendLine("\t\tSignature: %v", sarDataChecksumV2.Signature)
	return sarDataChecksumV2, nil
}

func parsePortalPlacementData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataPortalPlacement, int, error) {
	if length != 15 {
		return SarDataPortalPlacement{}, 0, errors.New("sar data invalid")
	}
	slot := int(reader.TryReadBytes(1))
	orange := reader.TryReadBool()
	reader.SkipBits(7)
	sarDataPortalPlacement := SarDataPortalPlacement{
		Orange: orange,
		X:      reader.TryReadFloat32(),
		Y:      reader.TryReadFloat32(),
		Z:      reader.TryReadFloat32(),
	}
	demo.Writer.AppendLine("\t\tOrange: %t", orange)
	demo.Writer.AppendLine("\t\tX: %f", sarDataPortalPlacement.X)
	demo.Writer.AppendLine("\t\tY: %f", sarDataPortalPlacement.Y)
	demo.Writer.AppendLine("\t\tZ: %f", sarDataPortalPlacement.Z)
	return sarDataPortalPlacement, slot, nil
}

func parseChallengeFlagsCrouchFlyData(reader *bitreader.Reader, length uint64) (int, error) {
	if length != 2 {
		return 0, errors.New("sar data invalid")
	}
	return int(reader.TryReadBytes(1)), nil
}

func parsePauseData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataPause, error) {
	if length != 5 {
		return SarDataPause{}, errors.New("sar data invalid")
	}
	sarDataPause := SarDataPause{
		PauseTicks: reader.TryReadUInt32(),
	}
	demo.Writer.AppendLine("\t\tPause Ticks: %d", sarDataPause.PauseTicks)
	return sarDataPause, nil
}

func parseWaitRunData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataWaitRun, error) {
	if length < 6 {
		return SarDataWaitRun{}, errors.New("sar data invalid")
	}
	sarDataWaitRun := SarDataWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}
	demo.Writer.AppendLine("\t\tTicks: %d", sarDataWaitRun.Ticks)
	demo.Writer.AppendLine("\t\tCmd: \"%s\"", sarDataWaitRun.Cmd)
	return sarDataWaitRun, nil
}

func parseHWaitRunData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataHWaitRun, error) {
	if length < 6 {
		return SarDataHWaitRun{}, errors.New("sar data invalid")
	}
	sarDataHWaitRun := SarDataHWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}
	demo.Writer.AppendLine("\t\tTicks: %d", sarDataHWaitRun.Ticks)
	demo.Writer.AppendLine("\t\tCmd: \"%s\"", sarDataHWaitRun.Cmd)
	return sarDataHWaitRun, nil
}

func parseSpeedrunTimeData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataSpeedrunTime, error) {
	if length < 5 {
		return SarDataSpeedrunTime{}, errors.New("sar data invalid")
	}
	numberOfSplits := reader.TryReadUInt32()
	splits := make([]SarDataSpeedrunTimeSplits, numberOfSplits)
	for splitCount := 0; splitCount < int(numberOfSplits); splitCount++ {
		splits[splitCount].Name = reader.TryReadString()
		splits[splitCount].NSegs = reader.TryReadUInt32()
		demo.Writer.AppendLine("\t\t[%d] Split Name: \"%s\"", splitCount, splits[splitCount].Name)
		demo.Writer.AppendLine("\t\t[%d] Number of Segments: %d", splitCount, splits[splitCount].NSegs)
		splits[splitCount].Segs = make([]SarDataSpeedrunTimeSegs, splits[splitCount].NSegs)
		for segCount := 0; segCount < int(splits[splitCount].NSegs); segCount++ {
			splits[splitCount].Segs[segCount].Name = reader.TryReadString()
			splits[splitCount].Segs[segCount].Ticks = reader.TryReadUInt32()
			demo.Writer.AppendLine("\t\t\t[%d] Segment Name: \"%s\"", segCount, splits[splitCount].Segs[segCount].Name)
			demo.Writer.AppendLine("\t\t\t[%d] Segment Ticks: %d", segCount, splits[splitCount].Segs[segCount].Ticks)
		}
	}
	return SarDataSpeedrunTime{
		NSplits: numberOfSplits,
		Splits:  splits,
	}, nil
}

func parseTimestampData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataTimestamp, error) {
	if length != 8 {
		return SarDataTimestamp{}, errors.New("sar data invalid")
	}
	timestamp := reader.TryReadBytesToSlice(7)
	sarDataTimeStamp := SarDataTimestamp{
		Year:   uint16(timestamp[0]) | uint16(timestamp[1])<<8,
		Month:  timestamp[2] + 1,
		Day:    timestamp[3],
		Hour:   timestamp[4],
		Minute: timestamp[5],
		Second: timestamp[6],
	}
	demo.Writer.AppendLine("\t\tYear: %d", sarDataTimeStamp.Year)
	demo.Writer.AppendLine("\t\tMonth: %d", sarDataTimeStamp.Month)
	demo.Writer.AppendLine("\t\tDay: %d", sarDataTimeStamp.Day)
	demo.Writer.AppendLine("\t\tHour: %d", sarDataTimeStamp.Hour)
	demo.Writer.AppendLine("\t\tMinute: %d", sarDataTimeStamp.Minute)
	demo.Writer.AppendLine("\t\tSecond: %d", sarDataTimeStamp.Second)
	return sarDataTimeStamp, nil
}

func parseFileChecksumData(reader *bitreader.Reader, length uint64, demo *types.Demo) (SarDataFileChecksum, error) {
	if length < 6 {
		return SarDataFileChecksum{}, errors.New("sar data invalid")
	}
	sarDataFileChecksum := SarDataFileChecksum{
		Sum:  reader.TryReadUInt32(),
		Path: reader.TryReadString(),
	}
	demo.Writer.AppendLine("\t\tChecksum: %d", sarDataFileChecksum.Sum)
	demo.Writer.AppendLine("\t\tPath: \"%s\"", sarDataFileChecksum.Path)
	return sarDataFileChecksum, nil
}
