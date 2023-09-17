package classes

import (
	"github.com/pektezol/bitreader"
)

type ServerClassInfo struct {
	ClassId       uint16
	ClassName     string
	DataTableName string
}

func ParseServerClassInfo(reader *bitreader.Reader, count int, numOfClasses int) ServerClassInfo {
	return ServerClassInfo{
		ClassId:       reader.TryReadUInt16(),
		ClassName:     reader.TryReadString(),
		DataTableName: reader.TryReadString(),
	}
}

// func serverClassBits(numOfClasses int) int {
// 	return highestBitIndex(uint(numOfClasses)) + 1
// }

// func highestBitIndex(i uint) int {
// 	var j int
// 	for j = 31; j >= 0 && (i&(1<<j)) == 0; j-- {
// 	}
// 	return j
// }
