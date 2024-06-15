package writer

import (
	"fmt"
	"strings"
)

type Writer struct {
	output strings.Builder
	temp   strings.Builder
}

func NewWriter() *Writer {
	return &Writer{
		output: strings.Builder{},
		temp:   strings.Builder{},
	}
}

func (w *Writer) Append(str string, a ...any) {
	_, err := w.output.WriteString(fmt.Sprintf(str, a...))
	if err != nil {
		w.output.WriteString(err.Error())
	}
}

func (w *Writer) AppendLine(str string, a ...any) {
	w.Append(str, a...)
	w.output.WriteString("\n")
}

func (w *Writer) GetOutputString() string {
	return w.output.String()
}

func (w *Writer) TempAppend(str string, a ...any) {
	_, err := w.temp.WriteString(fmt.Sprintf(str, a...))
	if err != nil {
		w.temp.WriteString(err.Error())
	}
}

func (w *Writer) TempAppendLine(str string, a ...any) {
	w.TempAppend(str, a...)
	w.temp.WriteString("\n")
}

func (w *Writer) TempGetString() string {
	return w.temp.String()
}

func (w *Writer) AppendOutputFromTemp() {
	w.output.WriteString(w.temp.String())
	w.temp.Reset()
}
