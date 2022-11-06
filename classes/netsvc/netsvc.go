package netsvc

import (
	"bytes"
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/utils"
)

const NET_TICK_SCALEUP = 10000

func ParseNetSvcMessage(file []byte) {
	reader := bitreader.Reader(bytes.NewReader(file), true)
	bitsRead := 0
	for {
		messageType, err := reader.ReadBits(6)
		if err != nil { // No remaining bits left
			break
		}
		switch messageType {
		case 16:
			var svcprint SvcPrint
			svcprint.Message = utils.ReadStringFromSlice(file)
			fmt.Println(svcprint)
			bitsRead += len(svcprint.Message) * 8
		default:
			//fmt.Println("default")
			break
		}
	}
}
