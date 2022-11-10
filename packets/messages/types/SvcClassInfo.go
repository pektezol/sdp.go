package types

import (
	"github.com/pektezol/bitreader"
)

type SvcClassInfo struct {
	CreateOnClient bool
	ServerClasses  []ServerClass
}

type ServerClass struct {
	ClassId       int32
	ClassName     string
	DataTableName string
}

func ParseSvcClassInfo(reader *bitreader.ReaderType) SvcClassInfo {
	length := reader.TryReadInt16()
	createonclient := reader.TryReadBool()
	var serverclasses []ServerClass
	if createonclient {
		serverclasses := make([]ServerClass, length)
		for i := 0; i < int(length); i++ {
			id, err := reader.ReadBits(HighestBitIndex(uint(length)) + 1)
			if err != nil {
				panic(err)
			}
			serverclasses[i] = ServerClass{
				ClassId:       int32(id),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			}
		}
	}
	return SvcClassInfo{
		CreateOnClient: createonclient,
		ServerClasses:  serverclasses,
	}
}

func HighestBitIndex(i uint) int {
	var j int
	for j = 31; j >= 0 && (i&(1<<j)) == 0; j-- {
	}
	return j
}
