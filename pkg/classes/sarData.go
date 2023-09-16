package classes

import (
	"errors"
	"fmt"

	"github.com/pektezol/bitreader"
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

func (t SarDataType) String() string {
	switch t {
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
		return fmt.Sprintf("%d", int(t))
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
	Year uint16
	Mon  uint8
	Day  uint8
	Hour uint8
	Min  uint8
	Sec  uint8
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
	return SarDataTimescaleCheat{
		Timescale: reader.TryReadFloat32(),
	}, nil
}

func parseInitialCVarData(reader *bitreader.Reader) SarDataInitialCVar {
	return SarDataInitialCVar{
		CVar: reader.TryReadString(),
		Val:  reader.TryReadString(),
	}
}

func parseEntityInputData(reader *bitreader.Reader) SarDataEntityInput {
	return SarDataEntityInput{
		TargetName: reader.TryReadString(),
		ClassName:  reader.TryReadString(),
		InputName:  reader.TryReadString(),
		Parameter:  reader.TryReadString(),
	}
}

func parseChecksumData(reader *bitreader.Reader, length uint64) (SarDataChecksum, error) {
	if length != 9 {
		return SarDataChecksum{}, errors.New("sar data invalid")
	}
	return SarDataChecksum{
		DemoSum: reader.TryReadUInt32(),
		SarSum:  reader.TryReadUInt32(),
	}, nil
}

func parseChecksumV2Data(reader *bitreader.Reader, length uint64) (SarDataChecksumV2, error) {
	if length != 69 {
		return SarDataChecksumV2{}, errors.New("sar data invalid")
	}
	return SarDataChecksumV2{
		SarSum:    reader.TryReadUInt32(),
		Signature: [64]byte(reader.TryReadBytesToSlice(60)),
	}, nil
}

func parsePortalPlacementData(reader *bitreader.Reader, length uint64) (SarDataPortalPlacement, int, error) {
	if length != 15 {
		return SarDataPortalPlacement{}, 0, errors.New("sar data invalid")
	}
	slot := int(reader.TryReadBytes(1))
	orange := reader.TryReadBool()
	reader.SkipBits(7)
	return SarDataPortalPlacement{
		Orange: orange,
		X:      reader.TryReadFloat32(),
		Y:      reader.TryReadFloat32(),
		Z:      reader.TryReadFloat32(),
	}, slot, nil
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
	return SarDataPause{
		PauseTicks: reader.TryReadUInt32(),
	}, nil
}

func parseWaitRunData(reader *bitreader.Reader, length uint64) (SarDataWaitRun, error) {
	if length < 6 {
		return SarDataWaitRun{}, errors.New("sar data invalid")
	}
	return SarDataWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}, nil
}

func parseHWaitRunData(reader *bitreader.Reader, length uint64) (SarDataHWaitRun, error) {
	if length < 6 {
		return SarDataHWaitRun{}, errors.New("sar data invalid")
	}
	return SarDataHWaitRun{
		Ticks: int(reader.TryReadUInt32()),
		Cmd:   reader.TryReadString(),
	}, nil
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
		splits[splitCount].Segs = make([]SarDataSpeedrunTimeSegs, splits[splitCount].NSegs)
		for segCount := 0; segCount < int(splits[splitCount].NSegs); segCount++ {
			splits[splitCount].Segs[segCount].Name = reader.TryReadString()
			splits[splitCount].Segs[segCount].Ticks = reader.TryReadUInt32()
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
	return SarDataTimestamp{
		Year: uint16(timestamp[0]) | uint16(timestamp[1]<<8),
		Mon:  timestamp[2] + 1,
		Day:  timestamp[3],
		Hour: timestamp[4],
		Min:  timestamp[5],
		Sec:  timestamp[6],
	}, nil
}

func parseFileChecksumData(reader *bitreader.Reader, length uint64) (SarDataFileChecksum, error) {
	if length < 6 {
		return SarDataFileChecksum{}, errors.New("sar data invalid")
	}
	return SarDataFileChecksum{
		Sum:  reader.TryReadUInt32(),
		Path: reader.TryReadString(),
	}, nil
}
