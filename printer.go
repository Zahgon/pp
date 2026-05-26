// printer.go: The actual pretty print implementation. Everything in this file should be private.
package pp

import (
	"bytes"
	"reflect"
	"text/tabwriter"

	"golang.org/x/text/message"
)

const (
	indentWidth = 2
)

func (pp *PrettyPrinter) format(object interface{}) string { _ = "STUB: not implemented"; return "" }

func newPrinter(object interface{}, currentScheme *ColorScheme, maxDepth int, coloringEnabled bool, decimalUint bool, exportedOnly bool, thousandsSeparator bool, omitEmpty bool) *printer {
	_ = "STUB: not implemented"
	return nil
}

type printer struct {
	*bytes.Buffer
	tw                 *tabwriter.Writer
	depth              int
	maxDepth           int
	value              reflect.Value
	visited            map[uintptr]bool
	currentScheme      *ColorScheme
	coloringEnabled    bool
	decimalUint        bool
	exportedOnly       bool
	thousandsSeparator bool
	omitEmpty          bool
	localizedPrinter   *message.Printer
}

func (p *printer) String() string { _ = "STUB: not implemented"; return "" }

func (p *printer) print(text string) { _ = "STUB: not implemented"; return }

func (p *printer) printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (p *printer) println(text string) { _ = "STUB: not implemented"; return }

func (p *printer) indentPrint(text string) { _ = "STUB: not implemented"; return }

func (p *printer) indentPrintf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (p *printer) colorPrint(text string, color uint16) { _ = "STUB: not implemented"; return }

func (p *printer) printString() { _ = "STUB: not implemented"; return }

// "\x00"

// "\u0000"

// "\U00000000"

// "\000"

func (p *printer) printMap() { _ = "STUB: not implemented"; return }

func (p *printer) printStruct() { _ = "STUB: not implemented"; return }

// ignore unexported if needed

// ignore empty fields if needed

// ignore fields with struct tags if zero value, or explicitly set

func (p *printer) printTime() { _ = "STUB: not implemented"; return }

func (p *printer) printSlice() { _ = "STUB: not implemented"; return }

// Stop travarsing cyclic reference

// Fold a large buffer

// indent for new group

// slice element

// space or newline

func (p *printer) printInterface() { _ = "STUB: not implemented"; return }

func (p *printer) printPtr() { _ = "STUB: not implemented"; return }

func (p *printer) pointerAddr() string { _ = "STUB: not implemented"; return "" }

func (p *printer) typeString() string { _ = "STUB: not implemented"; return "" }

func (p *printer) elemTypeString() string { _ = "STUB: not implemented"; return "" }

func (p *printer) colorizeType(t string) string { _ = "STUB: not implemented"; return "" }

func (p *printer) matchRegexp(text, exp string) bool { _ = "STUB: not implemented"; return false }

func (p *printer) indented(proc func()) { _ = "STUB: not implemented"; return }

func (p *printer) fmtOrLocalizedSprintf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *printer) raw() string {
	_ = "STUB: not implemented"
	// Some value causes panic when Interface() is called.
	return ""
}

func (p *printer) nil() string { _ = "STUB: not implemented"; return "" }

func (p *printer) colorize(text string, color uint16) string { _ = "STUB: not implemented"; return "" }

func (p *printer) format(object interface{}) string { _ = "STUB: not implemented"; return "" }

func (p *printer) indent() string { _ = "STUB: not implemented"; return "" }

// valueIsZero reports whether v is the zero value for its type.
// It returns false if the argument is invalid.
// This is a copy paste of reflect#IsZero from go1.15. It is not present before go1.13 (source: https://golang.org/doc/go1.13#library)
// source: https://golang.org/src/reflect/value.go?s=34297:34325#L1090
// This will need to be updated for new types or the decision should be made to drop support for Go version pre go1.13
func valueIsZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// this is the only difference between stdlib reflect#IsZero and this function. We're not going to
// panic on the default cause, even
