package utils

import (
	"fmt"
	"os"
)

func HeaderOut(file *os.File) {
	DemoFileStamp := make([]byte, 8)
	DemoProtocol := make([]byte, 4)
	NetworkProtocol := make([]byte, 4)
	ServerName := make([]byte, 260)
	ClientName := make([]byte, 260)
	MapName := make([]byte, 260)
	GameDirectory := make([]byte, 260)
	PlaybackTime := make([]byte, 4)
	PlaybackTicks := make([]byte, 4)
	PlaybackFrames := make([]byte, 4)
	SignOnLength := make([]byte, 4)
	file.Read(DemoFileStamp)
	file.Read(DemoProtocol)
	file.Read(NetworkProtocol)
	file.Read(ServerName)
	file.Read(ClientName)
	file.Read(MapName)
	file.Read(GameDirectory)
	file.Read(PlaybackTime)
	file.Read(PlaybackTicks)
	file.Read(PlaybackFrames)
	file.Read(SignOnLength)

	fmt.Println(string(DemoFileStamp))
	fmt.Println(IntFromBytes(DemoProtocol))
	fmt.Println(IntFromBytes(NetworkProtocol))
	fmt.Println(string(ServerName))
	fmt.Println(string(ClientName))
	fmt.Println(string(MapName))
	fmt.Println(string(GameDirectory))
	fmt.Println(FloatFromBytes(PlaybackTime))
	fmt.Println(IntFromBytes(PlaybackTicks))
	fmt.Println(IntFromBytes(PlaybackFrames))
	fmt.Println(IntFromBytes(SignOnLength))
}
