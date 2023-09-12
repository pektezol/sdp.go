package messages

import (
	"math"

	"github.com/pektezol/bitreader"
)

type SvcCreateStringTable struct {
	Name              string
	MaxEntries        int16
	NumEntries        int8
	Length            int32
	UserDataFixedSize bool
	UserDataSize      int16
	UserDataSizeBits  int8
	Flags             int8
	StringData        int
}

func ParseSvcCreateStringTable(reader *bitreader.ReaderType) SvcCreateStringTable {
	svcCreateStringTable := SvcCreateStringTable{
		Name:       reader.TryReadString(),
		MaxEntries: int16(reader.TryReadBits(16)),
	}
	svcCreateStringTable.NumEntries = int8(reader.TryReadBits(int(math.Log2(float64(svcCreateStringTable.MaxEntries))) + 1))
	svcCreateStringTable.Length = int32(reader.TryReadBits(20))
	svcCreateStringTable.UserDataFixedSize = reader.TryReadBool()
	if svcCreateStringTable.UserDataFixedSize {
		svcCreateStringTable.UserDataSize = int16(reader.TryReadBits(12))
		svcCreateStringTable.UserDataSizeBits = int8(reader.TryReadBits(4))
	}
	svcCreateStringTable.Flags = int8(reader.TryReadBits(2))
	reader.SkipBits(int(svcCreateStringTable.Length)) // TODO: StringTable parsing
	return svcCreateStringTable
}
