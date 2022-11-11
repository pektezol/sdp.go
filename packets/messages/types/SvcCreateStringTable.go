package types

import (
	"github.com/pektezol/bitreader"
)

type SvcCreateStringTable struct {
	Name              string
	MaxEntries        uint16
	NumEntries        uint8
	UserDataFixedSize bool
	UserDataSize      uint16
	UserDataSizeBits  uint8
	Flags             uint8
	StringData        []byte
}

func ParseSvcCreateStringTable(reader *bitreader.ReaderType) SvcCreateStringTable {
	svccreatestringtable := SvcCreateStringTable{
		Name:       reader.TryReadString(),
		MaxEntries: reader.TryReadInt16(),
	}
	svccreatestringtable.NumEntries = uint8(reader.TryReadBits(HighestBitIndex(uint(svccreatestringtable.MaxEntries)) + 1))
	length := reader.TryReadBits(20)
	svccreatestringtable.UserDataFixedSize = reader.TryReadBool()
	if svccreatestringtable.UserDataFixedSize {
		svccreatestringtable.UserDataSize = uint16(reader.TryReadBits(12))
		svccreatestringtable.UserDataSizeBits = uint8(reader.TryReadBits(4))
	}
	svccreatestringtable.Flags = uint8(reader.TryReadBits(2))
	svccreatestringtable.StringData = reader.TryReadBitsToSlice(int(length))
	return svccreatestringtable

}
