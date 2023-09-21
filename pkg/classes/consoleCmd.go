package classes

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type ConsoleCmd struct {
	Size int32
	Data string
}

func (consoleCmd *ConsoleCmd) ParseConsoleCmd(reader *bitreader.Reader) {
	consoleCmd.Size = reader.TryReadSInt32()
	consoleCmd.Data = reader.TryReadStringLength(uint64(consoleCmd.Size))
	writer.AppendLine("\t%s", strings.TrimSpace(consoleCmd.Data))
}
