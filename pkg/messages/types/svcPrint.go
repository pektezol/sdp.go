package messages

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/verification"
)

type SvcPrint struct {
	Message string
}

func ParseSvcPrint(reader *bitreader.Reader) SvcPrint {
	svcPrint := SvcPrint{
		Message: reader.TryReadString(),
	}
	// common psycopath behaviour
	print := fmt.Sprintf("\t\t%s\n", strings.Replace(strings.ReplaceAll(strings.ReplaceAll(svcPrint.Message, "\n", "\n\t\t"), "\n\t\t\n\t\t", ""), "\n\t\t", "", 1))
	// Define a regular expression pattern to match the "Server Number" line and capture the integer value.
	pattern := `Server Number: (\d+)`

	// Compile the regular expression pattern.
	re := regexp.MustCompile(pattern)

	// Find the match in the text.
	match := re.FindStringSubmatch(print)
	if len(match) >= 1 {
		serverNumber := match[1]
		n, _ := strconv.Atoi(serverNumber)
		verification.ServerNumbers = append(verification.ServerNumbers, n)
	}
	return svcPrint
}
