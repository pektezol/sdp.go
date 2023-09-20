package writer

import (
	"fmt"
	"strings"
)

var output strings.Builder

var temp strings.Builder

func Append(str string, a ...any) {
	_, err := output.WriteString(fmt.Sprintf(str, a...))
	if err != nil {
		output.WriteString(err.Error())
	}
}

func AppendLine(str string, a ...any) {
	Append(str, a...)
	output.WriteString("\n")
}

func GetString() string {
	return output.String()
}

func TempAppend(str string, a ...any) {
	_, err := temp.WriteString(fmt.Sprintf(str, a...))
	if err != nil {
		temp.WriteString(err.Error())
	}
}

func TempAppendLine(str string, a ...any) {
	TempAppend(str, a...)
	temp.WriteString("\n")
}

func TempGetString() string {
	return temp.String()
}

func AppendOutputFromTemp() {
	output.WriteString(temp.String())
	temp.Reset()
}
