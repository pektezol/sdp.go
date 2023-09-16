package messages

import (
	"fmt"
	"math"

	"github.com/pektezol/bitreader"
)

type SvcClassInfo struct {
	Length         int16
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
		Length:         int16(reader.TryReadBits(16)),
		CreateOnClient: reader.TryReadBool(),
	}
	classes := []serverClass{}
	if !svcClassInfo.CreateOnClient {
		for count := 0; count < int(svcClassInfo.Length); count++ {
			fmt.Println(classes)
			classes = append(classes, serverClass{
				ClassId:       int16(reader.TryReadBits(uint64(math.Log2(float64(svcClassInfo.Length)) + 1))),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			})
		}
	}
	svcClassInfo.ServerClasses = classes
	return svcClassInfo
}
