package classes

import (
	"github.com/pektezol/bitreader"
)

type ConsoleCmd struct {
	Size int32
	Data string
}

func (consoleCmd *ConsoleCmd) ParseConsoleCmd(reader *bitreader.Reader) {
	consoleCmd.Size = reader.TryReadSInt32()
	consoleCmd.Data = reader.TryReadStringLength(uint64(consoleCmd.Size))
}
