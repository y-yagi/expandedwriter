package expandedwriter_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/y-yagi/expandedwriter"
)

func TestRender(t *testing.T) {
	var outbuf bytes.Buffer
	w := expandedwriter.NewWriter(&outbuf)
	w.SetFields([]string{"ID", "email"})
	w.Append([]string{"1", "test@example.com"})
	w.Append([]string{"2", "longlonglonglonglonglonglonglong@example.com"})
	w.Append([]string{"3", "3@example.com"})
	if err := w.Render(); err != nil {
		t.Fatal(err)
	}

	expected := `--[ Data 1 ]----------------------------------------
ID    | 1
email | test@example.com
--[ Data 2 ]----------------------------------------
ID    | 2
email | longlonglonglonglonglonglonglong@example.com
--[ Data 3 ]----------------------------------------
ID    | 3
email | 3@example.com
`

	got := outbuf.String()
	if got != expected {
		t.Fatalf("Exepectd \n\n%s\nbut got\n\n%s\n", expected, got)
	}
}

func TestRender_WithoutFields(t *testing.T) {
	var outbuf bytes.Buffer
	w := expandedwriter.NewWriter(&outbuf)
	w.Append([]string{"1", "test@example.com"})
	w.Append([]string{"2", "longlonglonglonglonglonglonglong@example.com"})
	w.Append([]string{"3", "3@example.com"})
	if err := w.Render(); err != nil {
		t.Fatal(err)
	}

	expected := `--[ Data 1 ]-----------------------------------
 | 1
 | test@example.com
--[ Data 2 ]-----------------------------------
 | 2
 | longlonglonglonglonglonglonglong@example.com
--[ Data 3 ]-----------------------------------
 | 3
 | 3@example.com
`

	got := outbuf.String()
	if got != expected {
		t.Fatalf("Exepectd \n\n%s\nbut got\n\n%s\n", expected, got)
	}
}

func TestRender_WithCustomerHeader(t *testing.T) {
	var outbuf bytes.Buffer
	w := expandedwriter.NewWriter(&outbuf)
	w.SetHeaderName("Row")
	w.SetFields([]string{"ID", "email"})
	w.Append([]string{"1", "test@example.com"})
	w.Append([]string{"2", "longlonglonglonglonglonglonglong@example.com"})
	w.Append([]string{"3", "3@example.com"})
	if err := w.Render(); err != nil {
		t.Fatal(err)
	}

	expected := `--[ Row 1 ]-----------------------------------------
ID    | 1
email | test@example.com
--[ Row 2 ]-----------------------------------------
ID    | 2
email | longlonglonglonglonglonglonglong@example.com
--[ Row 3 ]-----------------------------------------
ID    | 3
email | 3@example.com
`

	got := outbuf.String()
	if got != expected {
		t.Fatalf("Exepectd \n\n%s\nbut got\n\n%s\n", expected, got)
	}
}

func TestRender_WithTerminal(t *testing.T) {
	w := expandedwriter.NewWriter(os.Stdout)
	w.SetHeaderName("Row")
	w.SetFields([]string{"ID", "email"})
	w.Append([]string{"1", "test@example.com"})
	w.Append([]string{"2", strings.Repeat("long", 50) + "@example.com"})
	w.Append([]string{"3", "3@example.com"})
	if err := w.Render(); err != nil {
		t.Fatal(err)
	}
}
