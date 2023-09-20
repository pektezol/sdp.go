package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/classes"
	"github.com/pektezol/demoparser/pkg/messages"
	"github.com/pektezol/demoparser/pkg/writer"
)

type PacketMessageInfo struct {
	PacketType uint8
	TickNumber int32
	SlotNumber uint8
	Data       any
}

const MSSC int = 2

func ParsePackets(reader *bitreader.Reader) PacketMessageInfo {
	packetType := reader.TryReadUInt8()
	tickNumber := reader.TryReadSInt32()
	slotNumber := reader.TryReadUInt8()
	var packetData any
	switch packetType {
	case 1: // SignOn
		signOn := SignOn{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "SIGNON", packetType)
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
		writer.AppendLine("[%d] %s (%d):", tickNumber, "PACKET", packetType)
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
		writer.AppendLine("[%d] %s (%d):", tickNumber, "SYNCTICK", packetType)
		packetData = syncTick
	case 4: // ConsoleCmd
		consoleCmd := ConsoleCmd{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "CONSOLECMD", packetType)
		consoleCmd.Size = reader.TryReadSInt32()
		consoleCmd.Data = reader.TryReadStringLength(uint64(consoleCmd.Size))
		writer.AppendLine("\t%s", consoleCmd.Data)
		packetData = consoleCmd
	case 5: // UserCmd TODO: usercmdinfo refactor
		userCmd := UserCmd{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "USERCMD", packetType)
		userCmd.Cmd = int32(reader.TryReadSInt32())
		userCmd.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(userCmd.Size))
		userCmd.Data = classes.ParseUserCmdInfo(data)
		packetData = userCmd
	case 6: // DataTables TODO: prop stuff
		dataTables := DataTables{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "DATATABLES", packetType)
		dataTables.Size = int32(reader.TryReadSInt32())
		data := reader.TryReadBytesToSlice(uint64(dataTables.Size))
		dataTableReader := bitreader.NewReaderFromBytes(data, true)
		count := 0
		for dataTableReader.TryReadBool() {
			count++
			dataTables.SendTable = append(dataTables.SendTable, classes.ParseSendTable(dataTableReader))
		}
		writer.AppendLine("\t%d Send Tables:", count)
		writer.AppendOutputFromTemp()
		numOfClasses := dataTableReader.TryReadBits(16)
		for count = 0; count < int(numOfClasses); count++ {
			dataTables.ServerClassInfo = append(dataTables.ServerClassInfo, classes.ParseServerClassInfo(dataTableReader, count, int(numOfClasses)))
		}
		writer.AppendLine("\t%d Classes:", count)
		writer.AppendOutputFromTemp()
		packetData = dataTables
	case 7: // Stop
		stop := Stop{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "STOP", packetType)
		if reader.TryReadBool() {
			stop.RemainingData = reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
			writer.AppendLine("\tRemaining Data: %v", stop.RemainingData)
		}
		packetData = stop
	case 8: // CustomData TODO: not sar data
		customData := CustomData{}
		customData.Type = reader.TryReadSInt32()
		customData.Size = reader.TryReadSInt32()
		if customData.Type != 0 || customData.Size == 8 {
			// Not SAR data
			writer.AppendLine("[%d] %s (%d):", tickNumber, "CUSTOMDATA", packetType)
			customData.Data = string(reader.TryReadBytesToSlice(uint64(customData.Size)))
			writer.AppendLine("\t%s", customData.Data)
			packetData = customData
			break
		}
		// SAR data
		sarData := classes.SarData{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "SARDATA", packetType)
		data := reader.TryReadBytesToSlice(uint64(customData.Size))
		sarReader := bitreader.NewReaderFromBytes(data, true)
		sarData.ParseSarData(sarReader)
		packetData = sarData
	case 9: // StringTables TODO: parsing string table data
		stringTables := StringTables{}
		writer.AppendLine("[%d] %s (%d):", tickNumber, "STRINGTABLES", packetType)
		stringTables.Size = reader.TryReadSInt32()
		data := reader.TryReadBytesToSlice(uint64(stringTables.Size))
		stringTableReader := bitreader.NewReaderFromBytes(data, true)
		stringTables.Data = classes.ParseStringTables(stringTableReader)
		packetData = stringTables
	default: // invalid
		writer.AppendLine("[%d] %s (%d):", tickNumber, "INVALID", packetType)
		panic("invalid packet type")
	}
	return PacketMessageInfo{
		PacketType: packetType,
		TickNumber: tickNumber,
		SlotNumber: slotNumber,
		Data:       packetData,
	}
}
