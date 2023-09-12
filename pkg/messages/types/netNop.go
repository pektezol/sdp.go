package messages

import "github.com/pektezol/bitreader"

type NetNop struct{}

func ParseNetNop(reader *bitreader.ReaderType) NetNop {
	return NetNop{}
}
