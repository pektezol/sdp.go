package classes

import (
	"github.com/pektezol/bitreader"
)

type ServerClassInfo struct {
	ClassId       int16
	ClassName     string
	DataTableName string
}

func ParseServerClassInfo(reader *bitreader.ReaderType, count int, numOfClasses int) ServerClassInfo {
	return ServerClassInfo{
		ClassId:       int16(reader.TryReadBits(16)),
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
