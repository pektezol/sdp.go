package classes

import (
	"errors"
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
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
	Type SarDataType
	Slot int
	Data any
}

type SarDataTimescaleCheat struct {
	Timescale float32
}

type SarDataInitialCVar struct {
	CVar string
	Val  string
}

type SarDataChecksum struct {
	DemoSum uint32
	SarSum  uint32
}

type SarDataChecksumV2 struct {
	SarSum    uint32
	Signature [64]byte
}

type SarDataEntityInput struct {
	TargetName string
	ClassName  string
	InputName  string
	Parameter  string
}

type SarDataPortalPlacement struct {
	Orange bool
	X      float32
	Y      float32
	Z      float32
}

type SarDataPause struct {
	PauseTicks uint32
}

type SarDataWaitRun struct {
	Ticks int
	Cmd   string
}

type SarDataHWaitRun struct {
	Ticks int
	Cmd   string
}

type SarDataSpeedrunTime struct {
	NSplits uint32
	Splits  []SarDataSpeedrunTimeSplits
}

type SarDataSpeedrunTimeSegs struct {
	Name  string
	Ticks uint32
}

type SarDataSpeedrunTimeSplits struct {
	Name  string
	NSegs uint32
	Segs  []SarDataSpeedrunTimeSegs
}

type SarDataTimestamp struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

type SarDataFileChecksum struct {
	Sum  uint32
	Path string
}

func (sarData *SarData) ParseSarData(reader *bitreader.Reader) (err error) {
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
	writer.AppendLine("\tMessage: %s (%d):", sarData.Type.String(), sarData.Type)
	switch sarData.Type {
	case ESarDataTimescaleCheat:
		sarData.Data, err = parseTimescaleCheatData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataInitialCVar:
		sarData.Data = parseInitialCVarData(dataReader)
	case ESarDataEntityInputSlot:
		sarData.Slot = int(dataReader.TryReadBytes(1))
		writer.AppendLine("\t\tSlot: %d", sarData.Slot)
	case ESarDataEntityInput:
		sarData.Data = parseEntityInputData(dataReader)
	case ESarDataChecksum:
		sarData.Data, err = parseChecksumData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataChecksumV2:
		sarData.Data, err = parseChecksumV2Data(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataPortalPlacement:
		data, slot, err := parsePortalPlacementData(dataReader, len)
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
		writer.AppendLine("\t\tSlot: %d", sarData.Slot)
	case ESarDataPause:
		sarData.Data, err = parsePauseData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataWaitRun:
		sarData.Data, err = parseWaitRunData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataHWaitRun:
		sarData.Data, err = parseHWaitRunData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataSpeedrunTime:
		sarData.Data, err = parseSpeedrunTimeData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataTimestamp:
		sarData.Data, err = parseTimestampData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataFileChecksum:
		sarData.Data, err = parseFileChecksumData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	default:
		err = errors.New("unsupported SAR data type")
		return err
	}
	return nil
}

func parseTimescaleCheatData(reader *bitreader.Reader, length uint64) (SarDataTimescaleCheat, error) {
	if length != 5 {
		return SarDataTimescaleCheat{}, errors.New("sar data invalid")
	}
	sarDataTimescaleCheat := SarDataTimescaleCheat{
		Timescale: reader.TryReadFloat32(),
	}
	writer.AppendLine("\t\tTimescale: %f", sarDataTimescaleCheat.Timescale)
	return sarDataTimescaleCheat, nil
}

func parseInitialCVarData(reader *bitreader.Reader) SarDataInitialCVar {
	sarDataInitialCvar := SarDataInitialCVar{
		CVar: reader.TryReadString(),
		Val:  reader.TryReadString(),
	}
	writer.AppendLine("\t\tCvar: \"%s\" = \"%s\"", sarDataInitialCvar.CVar, sarDataInitialCvar.Val)
	return sarDataInitialCvar
}

func parseEntityInputData(reader *bitreader.Reader) SarDataEntityInput {
	sarDataEntityInput := SarDataEntityInput{
		TargetName: reader.TryReadString(),
		ClassName:  reader.TryReadString(),
		InputName:  reader.TryReadString(),
		Parameter:  reader.TryReadString(),
	}
	writer.AppendLine("\t\tTarget: %s", sarDataEntityInput.TargetName)
	writer.AppendLine("\t\tClass: %s", sarDataEntityInput.ClassName)
	writer.AppendLine("\t\tInput: %s", sarDataEntityInput.InputName)
	writer.AppendLine("\t\tParameter: %s", sarDataEntityInput.Parameter)
	return sarDataEntityInput
}

func parseChecksumData(reader *bitreader.Reader, length uint64) (SarDataChecksum, error) {
	if length != 9 {
		return SarDataChecksum{}, errors.New("sar data invalid")
	}
	sarDataChecksum := SarDataChecksum{
		DemoSum: reader.TryReadUInt32(),
		SarSum:  reader.TryReadUInt32(),
	}
	writer.AppendLine("\t\tDemo Checksum: %d", sarDataChecksum.DemoSum)
	writer.AppendLine("\t\tSAR Checksum: %d", sarDataChecksum.SarSum)
	return sarDataChecksum, nil
}

func parseChecksumV2Data(reader *bitreader.Reader, length uint64) (SarDataChecksumV2, error) {
	if length != 69 {
		return SarDataChecksumV2{}, errors.New("sar data invalid")
	}
	sarDataChecksumV2 := SarDataChecksumV2{
		SarSum:    reader.TryReadUInt32(),
		Signature: [64]byte(reader.TryReadBytesToSlice(60)),
	}
	writer.AppendLine("\t\tSAR Checksum: %d", sarDataChecksumV2.SarSum)
	writer.AppendLine("\t\tSignature: %v", sarDataChecksumV2.Signature)
	return sarDataChecksumV2, nil
}

func parsePortalPlacementData(reader *bitreader.Reader, length uint64) (SarDataPortalPlacement, int, error) {
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
	writer.AppendLine("\t\tOrange: %t", orange)
	writer.AppendLine("\t\tX: %f", sarDataPortalPlacement.X)
	writer.AppendLine("\t\tY: %f", sarDataPortalPlacement.Y)
	writer.AppendLine("\t\tZ: %f", sarDataPortalPlacement.Z)
	return sarDataPortalPlacement, slot, nil
}

func parseChallengeFlagsCrouchFlyData(reader *bitreader.Reader, length uint64) (int, error) {
	if length != 2 {
		return 0, errors.New("sar data invalid")
	}
	return int(reader.TryReadBytes(1)), nil
}

func parsePauseData(reader *bitreader.Reader, length uint64) (SarDataPause, error) {
	if length != 5 {
		return SarDataPause{}, errors.New("sar data invalid")
	}
	sarDataPause := SarDataPause{
		PauseTicks: reader.TryReadUInt32(),
	}
	writer.AppendLine("\t\tPause Ticks: %d", sarDataPause.PauseTicks)
	return sarDataPause, nil
}

func parseWaitRunData(reader *bitreader.Reader, length uint64) (SarDataWaitRun, error) {
	if length < 6 {
		return SarDataWaitRun{}, errors.New("sar data invalid")
	}
	sarDataWaitRun := SarDataWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}
	writer.AppendLine("\t\tTicks: %d", sarDataWaitRun.Ticks)
	writer.AppendLine("\t\tCmd: \"%s\"", sarDataWaitRun.Cmd)
	return sarDataWaitRun, nil
}

func parseHWaitRunData(reader *bitreader.Reader, length uint64) (SarDataHWaitRun, error) {
	if length < 6 {
		return SarDataHWaitRun{}, errors.New("sar data invalid")
	}
	sarDataHWaitRun := SarDataHWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}
	writer.AppendLine("\t\tTicks: %d", sarDataHWaitRun.Ticks)
	writer.AppendLine("\t\tCmd: \"%s\"", sarDataHWaitRun.Cmd)
	return sarDataHWaitRun, nil
}

func parseSpeedrunTimeData(reader *bitreader.Reader, length uint64) (SarDataSpeedrunTime, error) {
	if length < 5 {
		return SarDataSpeedrunTime{}, errors.New("sar data invalid")
	}
	numberOfSplits := reader.TryReadUInt32()
	splits := make([]SarDataSpeedrunTimeSplits, numberOfSplits)
	for splitCount := 0; splitCount < int(numberOfSplits); splitCount++ {
		splits[splitCount].Name = reader.TryReadString()
		splits[splitCount].NSegs = reader.TryReadUInt32()
		writer.AppendLine("\t\t[%d] Split Name: \"%s\"", splitCount, splits[splitCount].Name)
		writer.AppendLine("\t\t[%d] Number of Segments: %d", splitCount, splits[splitCount].NSegs)
		splits[splitCount].Segs = make([]SarDataSpeedrunTimeSegs, splits[splitCount].NSegs)
		for segCount := 0; segCount < int(splits[splitCount].NSegs); segCount++ {
			splits[splitCount].Segs[segCount].Name = reader.TryReadString()
			splits[splitCount].Segs[segCount].Ticks = reader.TryReadUInt32()
			writer.AppendLine("\t\t\t[%d] Segment Name: \"%s\"", segCount, splits[splitCount].Segs[segCount].Name)
			writer.AppendLine("\t\t\t[%d] Segment Ticks: %d", segCount, splits[splitCount].Segs[segCount].Ticks)
		}
	}
	return SarDataSpeedrunTime{
		NSplits: numberOfSplits,
		Splits:  splits,
	}, nil
}

func parseTimestampData(reader *bitreader.Reader, length uint64) (SarDataTimestamp, error) {
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
	writer.AppendLine("\t\tYear: %d", sarDataTimeStamp.Year)
	writer.AppendLine("\t\tMonth: %d", sarDataTimeStamp.Month)
	writer.AppendLine("\t\tDay: %d", sarDataTimeStamp.Day)
	writer.AppendLine("\t\tHour: %d", sarDataTimeStamp.Hour)
	writer.AppendLine("\t\tMinute: %d", sarDataTimeStamp.Minute)
	writer.AppendLine("\t\tSecond: %d", sarDataTimeStamp.Second)
	return sarDataTimeStamp, nil
}

func parseFileChecksumData(reader *bitreader.Reader, length uint64) (SarDataFileChecksum, error) {
	if length < 6 {
		return SarDataFileChecksum{}, errors.New("sar data invalid")
	}
	sarDataFileChecksum := SarDataFileChecksum{
		Sum:  reader.TryReadUInt32(),
		Path: reader.TryReadString(),
	}
	writer.AppendLine("\t\tChecksum: %d", sarDataFileChecksum.Sum)
	writer.AppendLine("\t\tPath: \"%s\"", sarDataFileChecksum.Path)
	return sarDataFileChecksum, nil
}
