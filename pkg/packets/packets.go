package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/classes"
	"github.com/pektezol/demoparser/pkg/messages"
)

type PacketMessageInfo struct {
	PacketType int8
	TickNumber int32
	SlotNumber int8
	Data       any
}

const MSSC int = 2

func ParsePackets(reader *bitreader.Reader) PacketMessageInfo {
	packetType := reader.TryReadBits(8)
	tickNumber := reader.TryReadBits(32)
	slotNumber := reader.TryReadBits(8)
	var packetData any
	switch packetType {
	case 1: // SignOn
		signOn := SignOn{}
		for count := 0; count < MSSC; count++ {
			signOn.PacketInfo = append(signOn.PacketInfo, classes.ParseCmdInfo(reader))
		}
		signOn.InSequence = int32(reader.TryReadBits(32))
		signOn.OutSequence = int32(reader.TryReadBits(32))
		signOn.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(signOn.Size))
		packetReader := bitreader.NewReaderFromBytes(data, true)
		for {
			messageType, err := packetReader.ReadBits(6)
			if err != nil {
				break
			}
			signOn.Data = append(signOn.Data, messages.ParseMessages(int(messageType), packetReader))
		}
		packetData = signOn
	case 2: // Packet
		packet := Packet{}
		for count := 0; count < MSSC; count++ {
			packet.PacketInfo = append(packet.PacketInfo, classes.ParseCmdInfo(reader))
		}
		packet.InSequence = int32(reader.TryReadBits(32))
		packet.OutSequence = int32(reader.TryReadBits(32))
		packet.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(packet.Size))
		packetReader := bitreader.NewReaderFromBytes(data, true)
		for {
			messageType, err := packetReader.ReadBits(6)
			if err != nil {
				break
			}
			packet.Data = append(packet.Data, messages.ParseMessages(int(messageType), packetReader))
		}
		packetData = packet
	case 3: // SyncTick
		syncTick := SyncTick{}
		packetData = syncTick
	case 4: // ConsoleCmd
		size := reader.TryReadSInt32()
		consoleCmd := ConsoleCmd{
			Size: int32(size),
			Data: reader.TryReadStringLength(uint64(size)),
		}
		packetData = consoleCmd
	case 5: // UserCmd
		userCmd := UserCmd{}
		userCmd.Cmd = int32(reader.TryReadSInt32())
		userCmd.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(userCmd.Size))
		userCmd.Data = classes.ParseUserCmdInfo(data)
		packetData = userCmd
	case 6: // DataTables
		dataTables := DataTables{}
		dataTables.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(dataTables.Size))
		dataTableReader := bitreader.NewReaderFromBytes(data, true)
		count := 0
		for dataTableReader.TryReadBool() {
			count++
			dataTables.SendTable = append(dataTables.SendTable, classes.ParseSendTable(dataTableReader))
		}
		numOfClasses := dataTableReader.TryReadBits(16)
		for count = 0; count < int(numOfClasses); count++ {
			dataTables.ServerClassInfo = append(dataTables.ServerClassInfo, classes.ParseServerClassInfo(dataTableReader, count, int(numOfClasses)))
		}
		packetData = dataTables
	case 7: // Stop
		stop := Stop{}
		if reader.TryReadBool() {
			stop.RemainingData = reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
		}
		packetData = stop
	case 8: // CustomData
		customData := CustomData{
			Unknown: int32(reader.TryReadBits(32)),
			Size:    int32(reader.TryReadBits(32)),
		}
		customData.Data = string(reader.TryReadBytesToSlice(uint64(customData.Size)))
		packetData = customData
	case 9: // StringTables
		stringTables := StringTables{
			Size: int32(reader.TryReadSInt32()),
		}
		data := reader.TryReadBytesToSlice(uint64(stringTables.Size))
		stringTableReader := bitreader.NewReaderFromBytes(data, true)
		stringTables.Data = classes.ParseStringTables(stringTableReader)
		packetData = stringTables
	default: // invalid
		panic("invalid packet type")
	}
	return PacketMessageInfo{
		PacketType: int8(packetType),
		TickNumber: int32(tickNumber),
		SlotNumber: int8(slotNumber),
		Data:       packetData,
	}
}
