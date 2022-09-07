package utils

import "os"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadByteFromFile(file *os.File, size int32) []byte {
	tmp := make([]byte, size)
	file.Read(tmp)
	return tmp
}
