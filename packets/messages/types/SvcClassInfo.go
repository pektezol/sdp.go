package types

import (
	"math"

	"github.com/pektezol/bitreader"
)

type SvcClassInfo struct {
	Length         uint16
	CreateOnClient bool
	ServerClasses  []ServerClass
}

type ServerClass struct {
	ClassId       int32
	ClassName     string
	DataTableName string
}

func ParseSvcClassInfo(reader *bitreader.ReaderType) SvcClassInfo {
	var serverclasses []ServerClass
	svcclassinfo := SvcClassInfo{
		Length:         reader.TryReadInt16(),
		CreateOnClient: reader.TryReadBool(),
	}
	if svcclassinfo.CreateOnClient {
		for i := 0; i < int(svcclassinfo.Length); i++ {
			id, err := reader.ReadBits(int(math.Log2(float64(svcclassinfo.Length))) + 1)
			if err != nil {
				panic(err)
			}
			serverclasses = append(serverclasses, ServerClass{
				ClassId:       int32(id),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			})
		}
	}
	svcclassinfo.ServerClasses = serverclasses
	return svcclassinfo
}
