package classes

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type ConsoleCmd struct {
	Size int32  `json:"size"`
	Data string `json:"data"`
}

func (consoleCmd *ConsoleCmd) ParseConsoleCmd(reader *bitreader.Reader, demo *types.Demo) {
	consoleCmd.Size = reader.TryReadSInt32()
	consoleCmd.Data = reader.TryReadStringLength(uint64(consoleCmd.Size))
	demo.Writer.AppendLine("\t%s", strings.TrimSpace(consoleCmd.Data))
}
