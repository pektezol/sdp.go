package classes

import (
	"encoding/binary"
	"os"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/utils"
)

func ParseCmdInfo(file *os.File, MSSC int) []CmdInfo {
	array := utils.ReadByteFromFile(file, 76*int32(MSSC))
	var cmdinfoarray []CmdInfo
	for count := 0; count < MSSC; count++ {
		var cmdinfo CmdInfo
		cmdinfo.Flags = int32(binary.LittleEndian.Uint32(array[0+76*count : 4+76*count]))
		cmdinfo.ViewOrigin = utils.FloatArrFromBytes(array[4+76*count : 16+76*count])
		cmdinfo.ViewAngles = utils.FloatArrFromBytes(array[16+76*count : 28+76*count])
		cmdinfo.LocalViewAngles = utils.FloatArrFromBytes(array[28+76*count : 40+76*count])
		cmdinfo.ViewOrigin2 = utils.FloatArrFromBytes(array[40+76*count : 52+76*count])
		cmdinfo.ViewAngles2 = utils.FloatArrFromBytes(array[52+76*count : 64+76*count])
		cmdinfo.LocalViewAngles2 = utils.FloatArrFromBytes(array[64+76*count : 76+76*count])
		cmdinfoarray = append(cmdinfoarray, cmdinfo)
	}
	return cmdinfoarray
}

func ParseUserCmdInfo(file *os.File, size int) UserCmdInfo {
	reader := bitreader.Reader(file, true)
	var usercmd UserCmdInfo
	count := 0
	flag, err := reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.CommandNumber = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.TickCount = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesX = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesY = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ViewAnglesZ = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.ForwardMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.SideMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.UpMove = reader.TryReadFloat32()
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.Buttons = int32(reader.TryReadInt32())
		count += 32
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		//reader.SkipBits(8)
		usercmd.Impulse = int8(reader.TryReadInt8())
		count += 8
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		value, err := reader.ReadBits(11)
		utils.CheckError(err)
		usercmd.WeaponSelect = int(value)
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		count += 11
		if flag {
			value, err := reader.ReadBits(6)
			utils.CheckError(err)
			usercmd.WeaponSubtype = int(value)
			count += 6
		}
		count++
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.MouseDx = int16(reader.TryReadInt16())
		count += 16
	}
	count++
	flag, err = reader.ReadBool()
	utils.CheckError(err)
	if flag {
		usercmd.MouseDy = int16(reader.TryReadInt16())
		count += 16
	}
	count++
	reader.SkipBits(size*8 - count) // Skip remaining bits from specified size
	return usercmd
}

/*func ParseStringTable(file *os.File, size int) []StringTable {
	reader := bitreader.Reader(file, true)
	var stringtable StringTable
	var stringtablearray []StringTable
	//count := 0
	stringtable.NumOfTables = int8(reader.TryReadInt8())
	for i := 0; i < int(stringtable.NumOfTables); i++ {
		stringtable.TableName = utils.ReadStringFromFile(file)
		stringtable.NumOfEntries = int16(reader.TryReadInt16())
		stringtable.EntryName = utils.ReadStringFromFile(file)
		flag, err := reader.ReadBool()
		utils.CheckError(err)
		if flag {
			stringtable.EntrySize = int16(reader.TryReadInt16())
		}
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		if flag {
			fmt.Println(int(stringtable.EntrySize))
			reader.SkipBytes(int(stringtable.EntrySize))
			var bytearray []byte
			for i := 0; i < int(stringtable.EntrySize); i++ {
				value, err := reader.ReadBytes(1)
				utils.CheckError(err)
				bytearray = append(bytearray, byte(value))
			}
			stringtable.EntryData = bytearrray
		}
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		if flag {
			stringtable.NumOfClientEntries = int16(reader.TryReadInt16())
		}
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		if flag {
			stringtable.ClientEntryName = utils.ReadStringFromFile(file)
		}
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		if flag {
			stringtable.ClientEntrySize = int16(reader.TryReadInt16())
		}
		flag, err = reader.ReadBool()
		utils.CheckError(err)
		if flag {
			reader.SkipBytes(int(stringtable.ClientEntrySize))
			/*var bytearray []byte
			for i := 0; i < int(stringtable.ClientEntrySize); i++ {
				value, err := reader.ReadBytes(1)
				utils.CheckError(err)
				bytearray = append(bytearray, byte(value))
			}
			stringtable.ClientEntryData = bytearrray
		}
		stringtablearray = append(stringtablearray, stringtable)
	}

	//reader.SkipBits(size*8 - 8)
	return stringtablearray
}*/
