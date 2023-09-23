package classes

import (
	"errors"
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/verification"
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
	switch sarData.Type {
	case ESarDataTimescaleCheat:
		fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parseTimescaleCheatData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataInitialCVar:

		sarData.Data = parseInitialCVarData(dataReader)
	case ESarDataEntityInputSlot:
		sarData.Slot = int(dataReader.TryReadBytes(1))
	case ESarDataEntityInput:
		sarData.Data = parseEntityInputData(dataReader)
	case ESarDataChecksum:
		// fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parseChecksumData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataChecksumV2:
		// fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

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
		// fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Slot, err = parseChallengeFlagsCrouchFlyData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
		// fmt.Printf("\t\tSlot: %d\n", sarData.Slot)
	case ESarDataPause:
		fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parsePauseData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataWaitRun:
		fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parseWaitRunData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataHWaitRun:
		fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parseHWaitRunData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataSpeedrunTime:
		fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

		sarData.Data, err = parseSpeedrunTimeData(dataReader, len)
		if err != nil {
			sarData.Data = nil
		}
	case ESarDataTimestamp:
		// fmt.Printf("\tMessage: %s (%d):\n", sarData.Type.String(), sarData.Type)

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
	fmt.Printf("\t\tTimescale: %f\n", sarDataTimescaleCheat.Timescale)
	return sarDataTimescaleCheat, nil
}

func parseInitialCVarData(reader *bitreader.Reader) SarDataInitialCVar {
	sarDataInitialCvar := SarDataInitialCVar{
		CVar: reader.TryReadString(),
		Val:  reader.TryReadString(),
	}
	// fmt.Printf("\t\tCvar: \"%s\" = \"%s\"\n", sarDataInitialCvar.CVar, sarDataInitialCvar.Val)
	return sarDataInitialCvar
}

func parseEntityInputData(reader *bitreader.Reader) SarDataEntityInput {
	sarDataEntityInput := SarDataEntityInput{
		TargetName: reader.TryReadString(),
		ClassName:  reader.TryReadString(),
		InputName:  reader.TryReadString(),
		Parameter:  reader.TryReadString(),
	}
	// fmt.Printf("\t\tTarget: %s\n", sarDataEntityInput.TargetName)
	// fmt.Printf("\t\tClass: %s\n", sarDataEntityInput.ClassName)
	// fmt.Printf("\t\tInput: %s\n", sarDataEntityInput.InputName)
	// fmt.Printf("\t\tParameter: %s\n", sarDataEntityInput.Parameter)
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
	// fmt.Printf("\t\tDemo Checksum: %d\n", sarDataChecksum.DemoSum)
	// fmt.Printf("\t\tSAR Checksum: %d\n", sarDataChecksum.SarSum)
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
	// fmt.Printf("\t\tSAR Checksum: %d\n", sarDataChecksumV2.SarSum)
	// fmt.Printf("\t\tSignature: %v\n", sarDataChecksumV2.Signature)
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
	// fmt.Printf("\t\tOrange: %t\n", orange)
	// fmt.Printf("\t\tX: %f\n", sarDataPortalPlacement.X)
	// fmt.Printf("\t\tY: %f\n", sarDataPortalPlacement.Y)
	// fmt.Printf("\t\tZ: %f\n", sarDataPortalPlacement.Z)
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
	fmt.Printf("\t\tPause Ticks: %d\n", sarDataPause.PauseTicks)
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
	fmt.Printf("\t\tTicks: %d\n", sarDataWaitRun.Ticks)
	fmt.Printf("\t\tCmd: \"%s\"\n", sarDataWaitRun.Cmd)
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
	fmt.Printf("\t\tTicks: %d\n", sarDataHWaitRun.Ticks)
	fmt.Printf("\t\tCmd: \"%s\"\n", sarDataHWaitRun.Cmd)
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
		fmt.Printf("\t\t[%d] Split Name: \"%s\"\n", splitCount, splits[splitCount].Name)
		fmt.Printf("\t\t[%d] Number of Segments: %d\n", splitCount, splits[splitCount].NSegs)
		splits[splitCount].Segs = make([]SarDataSpeedrunTimeSegs, splits[splitCount].NSegs)
		for segCount := 0; segCount < int(splits[splitCount].NSegs); segCount++ {
			splits[splitCount].Segs[segCount].Name = reader.TryReadString()
			splits[splitCount].Segs[segCount].Ticks = reader.TryReadUInt32()
			verification.Ticks += splits[splitCount].Segs[segCount].Ticks
			fmt.Printf("\t\t\t[%d] Segment Name: \"%s\"\n", segCount, splits[splitCount].Segs[segCount].Name)
			fmt.Printf("\t\t\t[%d] Segment Ticks: %d\n", segCount, splits[splitCount].Segs[segCount].Ticks)
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
	// fmt.Printf("\t\tYear: %d\n", sarDataTimeStamp.Year)
	// fmt.Printf("\t\tMonth: %d\n", sarDataTimeStamp.Month)
	// fmt.Printf("\t\tDay: %d\n", sarDataTimeStamp.Day)
	// fmt.Printf("\t\tHour: %d\n", sarDataTimeStamp.Hour)
	// fmt.Printf("\t\tMinute: %d\n", sarDataTimeStamp.Minute)
	// fmt.Printf("\t\tSecond: %d\n", sarDataTimeStamp.Second)
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
	// fmt.Printf("\t\tChecksum: %d\n", sarDataFileChecksum.Sum)
	// fmt.Printf("\t\tPath: \"%s\"\n", sarDataFileChecksum.Path)
	return sarDataFileChecksum, nil
}
