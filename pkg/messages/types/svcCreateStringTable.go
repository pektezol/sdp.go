package messages

import (
	"math"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcCreateStringTable struct {
	Name              string `json:"name"`
	MaxEntries        int16  `json:"max_entries"`
	NumEntries        int8   `json:"num_entries"`
	Length            int32  `json:"length"`
	UserDataFixedSize bool   `json:"user_data_fixed_size"`
	UserDataSize      int16  `json:"user_data_size"`
	UserDataSizeBits  int8   `json:"user_data_size_bits"`
	Flags             int8   `json:"flags"`
	StringData        int    `json:"string_data"`
}

func ParseSvcCreateStringTable(reader *bitreader.Reader, demo *types.Demo) SvcCreateStringTable {
	svcCreateStringTable := SvcCreateStringTable{
		Name:       reader.TryReadString(),
		MaxEntries: reader.TryReadSInt16(),
	}
	svcCreateStringTable.NumEntries = int8(reader.TryReadBits(uint64(math.Log2(float64(svcCreateStringTable.MaxEntries))) + 1))
	svcCreateStringTable.Length = int32(reader.TryReadBits(20))
	svcCreateStringTable.UserDataFixedSize = reader.TryReadBool()
	if svcCreateStringTable.UserDataFixedSize {
		svcCreateStringTable.UserDataSize = int16(reader.TryReadBits(12))
		svcCreateStringTable.UserDataSizeBits = int8(reader.TryReadBits(4))
	}
	svcCreateStringTable.Flags = int8(reader.TryReadBits(2))
	demo.Writer.TempAppendLine("\t\tName: %s", svcCreateStringTable.Name)
	demo.Writer.TempAppendLine("\t\tMax Enties: %d", svcCreateStringTable.MaxEntries)
	demo.Writer.TempAppendLine("\t\tNumber Of Entiries: %d", svcCreateStringTable.NumEntries)
	demo.Writer.TempAppendLine("\t\tUser Data Fixed Size: %t", svcCreateStringTable.UserDataFixedSize)
	demo.Writer.TempAppendLine("\t\tUser Data Size: %d", svcCreateStringTable.UserDataSize)
	demo.Writer.TempAppendLine("\t\tUser Data Size In Bits: %d", svcCreateStringTable.UserDataSizeBits)
	demo.Writer.TempAppendLine("\t\tFlags: %d", svcCreateStringTable.Flags)
	reader.SkipBits(uint64(svcCreateStringTable.Length)) // TODO: StringTable parsing
	return svcCreateStringTable
}
