package expandedwriter

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

type Expandedwriter struct {
	w              io.Writer
	fields         []string
	values         [][]string
	valueMaxLength int
	fieldMaxLength int
	headerName     string
	terminalWidth  int
}

func NewWriter(w io.Writer) *Expandedwriter {
	writer := &Expandedwriter{w: w, headerName: "Data", terminalWidth: -1}
	if w, ok := w.(*os.File); ok {
		if width, _, err := term.GetSize(int(w.Fd())); err == nil {
			writer.terminalWidth = width
		}
	}

	return writer
}

func (ew *Expandedwriter) SetFields(fields []string) {
	for _, v := range fields {
		ew.fieldMaxLength = ew.max(ew.fieldMaxLength, len(v))
	}
	ew.fields = fields
}

func (ew *Expandedwriter) SetHeaderName(headername string) {
	ew.headerName = headername
}

func (ew *Expandedwriter) Append(value []string) {
	for _, v := range value {
		ew.valueMaxLength = ew.max(ew.valueMaxLength, len(v))
	}
	ew.values = append(ew.values, value)
}

func (ew *Expandedwriter) Render() error {
	result := ""
	delimiterSizeForFieldAndValue := 3

	headerSize := ew.valueMaxLength + ew.fieldMaxLength + delimiterSizeForFieldAndValue
	if ew.terminalWidth != -1 && headerSize > ew.terminalWidth {
		headerSize = ew.terminalWidth
	}

	for i, value := range ew.values {
		header := fmt.Sprintf("--[ "+ew.headerName+" %d ]", i+1)
		if len(header) < headerSize {
			header += strings.Repeat("-", headerSize-len(header))
		}

		result += header

		for i, v := range value {
			field := ""
			if len(ew.fields) > i {
				field = ew.fields[i]
			}

			result += fmt.Sprintf("\n%-*s | %s", ew.fieldMaxLength, field, v)
		}
		result += "\n"
	}

	_, err := ew.w.Write([]byte(result))
	return err
}

func (ew *Expandedwriter) max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
