package messages

import (
	"math"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcClassInfo struct {
	ClassCount     uint16        `json:"class_count"`
	CreateOnClient bool          `json:"create_on_client"`
	ServerClasses  []serverClass `json:"server_classes"`
}

type serverClass struct {
	ClassId       int16  `json:"class_id"`
	ClassName     string `json:"class_name"`
	DataTableName string `json:"data_table_name"`
}

func ParseSvcClassInfo(reader *bitreader.Reader, demo *types.Demo) SvcClassInfo {
	svcClassInfo := SvcClassInfo{
		ClassCount:     reader.TryReadUInt16(),
		CreateOnClient: reader.TryReadBool(),
	}
	classes := []serverClass{}
	demo.Writer.TempAppendLine("\t\tCreate On Client: %t", svcClassInfo.CreateOnClient)
	if !svcClassInfo.CreateOnClient {
		demo.Writer.TempAppendLine("\t\t%d Server Classes:", svcClassInfo.ClassCount)
		for count := 0; count < int(svcClassInfo.ClassCount); count++ {
			classes = append(classes, serverClass{
				ClassId:       int16(reader.TryReadBits(uint64(math.Log2(float64(svcClassInfo.ClassCount)) + 1))),
				ClassName:     reader.TryReadString(),
				DataTableName: reader.TryReadString(),
			})
			demo.Writer.TempAppendLine("\t\t\t[%d] %s (%s)", classes[len(classes)-1].ClassId, classes[len(classes)-1].ClassName, classes[len(classes)-1].DataTableName)
		}
	} else {
		demo.Writer.TempAppendLine("\t\t%d Server Classes", svcClassInfo.ClassCount)
	}
	svcClassInfo.ServerClasses = classes
	return svcClassInfo
}
