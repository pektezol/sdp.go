package messages

import (
	"math"

	"github.com/pektezol/bitreader"
)

type SvcClassInfo struct {
	ClassCount     uint16
	CreateOnClient bool
	ServerClasses  []serverClass
}

type serverClass struct {
	ClassId       int16
	ClassName     string
	DataTableName string
}

func ParseSvcClassInfo(reader *bitreader.Reader) SvcClassInfo {
	svcClassInfo := SvcClassInfo{
		ClassCount:     reader.TryReadUInt16(),
		CreateOnClient: reader.TryReadBool(),
	}
	classes := []serverClass{}

	if !svcClassInfo.CreateOnClient {

		for count := 0; count < int(svcClassInfo.ClassCount); count++ {
			classes = append(classes, serverClass{
				ClassId:       int16(reader.TryReadBits(uint64(math.Log2(float64(svcClassInfo.ClassCount)) + 1))),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			})

		}
	}
	svcClassInfo.ServerClasses = classes
	return svcClassInfo
}
