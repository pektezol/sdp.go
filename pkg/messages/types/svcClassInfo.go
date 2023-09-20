package messages

import (
	"math"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tCreate On Client: %t", svcClassInfo.CreateOnClient)
	if !svcClassInfo.CreateOnClient {
		writer.TempAppendLine("\t\t%d Server Classes:", svcClassInfo.ClassCount)
		for count := 0; count < int(svcClassInfo.ClassCount); count++ {
			classes = append(classes, serverClass{
				ClassId:       int16(reader.TryReadBits(uint64(math.Log2(float64(svcClassInfo.ClassCount)) + 1))),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			})
			writer.TempAppendLine("\t\t\t[%d] %s (%s)", classes[len(classes)-1].ClassId, classes[len(classes)-1].ClassName, classes[len(classes)-1].DataTableName)
		}
	} else {
		writer.TempAppendLine("\t\t%d Server Classes", svcClassInfo.ClassCount)
	}
	svcClassInfo.ServerClasses = classes
	return svcClassInfo
}
