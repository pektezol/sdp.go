package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/packets"
	"github.com/pektezol/demoparser/pkg/verification"
)

const littleEndian bool = true

func main() {
	fmt.Println("Portal 2 Run Validity Checker Tool")
	fmt.Println()
	if len(os.Args) != 2 {
		panic("specify file in command line arguments")
	}
	files, err := os.ReadDir(os.Args[1])
	if err != nil { // If it's not a directory
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		reader := bitreader.NewReader(file, littleEndian)
		demoParserHandler(reader)
		defer file.Close()
		fmt.Println()
		fmt.Println("Total Ticks from SAR:", verification.Ticks)
		fmt.Println()
		if verification.IsContinuous(verification.ServerNumbers) {
			fmt.Println("Server Numbers: VALID")
			fmt.Println(verification.ServerNumbers)
		} else {
			fmt.Println("[!] Server Numbers: NOT VALID")
			fmt.Println(verification.ServerNumbers)
		}
		fmt.Scanln()
		return
	}
	sort.Slice(files, func(i, j int) bool {
		// Extract numeric parts from file names
		numA := extractNumber(files[i].Name())
		numB := extractNumber(files[j].Name())

		// Compare the extracted numbers for sorting
		return numA < numB
	})
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		if err != nil {
			panic(err)
		}
		reader := bitreader.NewReader(file, littleEndian)
		demoParserHandler(reader)
		defer file.Close()
	}
	fmt.Println()
	fmt.Println("Total Ticks from SAR:", verification.Ticks)
	fmt.Println()
	if verification.IsContinuous(verification.ServerNumbers) {
		fmt.Println("Server Numbers: VALID")
		fmt.Println(verification.ServerNumbers)
	} else {
		fmt.Println("[!] Server Numbers: NOT VALID")
		fmt.Println(verification.ServerNumbers)
	}
	fmt.Scanln()
}

func demoParserHandler(reader *bitreader.Reader) {
	packets.ParseHeaders(reader)
	for {
		packet := packets.ParsePackets(reader)
		if packet.PacketType == 7 {
			break
		}
	}
}

func extractNumber(filename string) int {
	// Split the filename by underscores
	parts := strings.Split(filename, "_")

	// Get the last part which should be the number
	lastPart := parts[len(parts)-1]

	// Remove any non-digit characters
	numStr := strings.TrimSuffix(lastPart, ".dem") // Assuming filenames end with ".txt"
	num, _ := strconv.Atoi(numStr)

	return num
}
