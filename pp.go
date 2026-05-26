// pp.go: API definitions. The core implementation is delegated to printer.go.
package pp

import (
	"io"
	"sync"

	"github.com/mattn/go-colorable"
)

// Global variable API
// see also: color.go
var (
	// Default pretty printer. It's public so that you can modify config globally.
	Default = newPrettyPrinter(3) // pp.* => PrettyPrinter.* => formatAll
	// If the length of array or slice is larger than this,
	// the buffer will be shorten as {...}.
	BufferFoldThreshold = 1024
	// PrintMapTypes when set to true will have map types will always appended to maps.
	PrintMapTypes = true
	// WithLineInfo add file name and line information to output
	// call this function with care, because getting stack has performance penalty
	WithLineInfo bool
)

// Internals
var (
	defaultOut          = colorable.NewColorableStdout()
	defaultWithLineInfo = false
)

type PrettyPrinter struct {
	// WithLineInfo adds file name and line information to output.
	// Call this function with care, because getting stack has performance penalty.
	WithLineInfo bool
	// To support WithLineInfo, we need to know which frame we should look at.
	// Thus callerLevel sets the number of frames it needs to skip.
	callerLevel        int
	out                io.Writer
	currentScheme      ColorScheme
	outLock            sync.Mutex
	maxDepth           int
	coloringEnabled    bool
	decimalUint        bool
	thousandsSeparator bool
	// This skips unexported fields of structs.
	exportedOnly bool

	// This skips empty fields of structs.
	omitEmpty bool
}

// New creates a new PrettyPrinter that can be used to pretty print values
func New() *PrettyPrinter { _ = "STUB: not implemented"; return nil }

// PrettyPrinter.* => formatAll

func newPrettyPrinter(callerLevel int) *PrettyPrinter { _ = "STUB: not implemented"; return nil }

// adjustFormat adjust format, if print line.
func adjustFormat(format string, withLine bool) string { _ = "STUB: not implemented"; return "" }

// Print prints given arguments.
func (pp *PrettyPrinter) Print(a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Printf prints a given format.
func (pp *PrettyPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Println prints given arguments with newline.
func (pp *PrettyPrinter) Println(a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Sprint formats given arguments and returns the result as string.
func (pp *PrettyPrinter) Sprint(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Sprintf formats with pretty print and returns the result as string.
func (pp *PrettyPrinter) Sprintf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Sprintln formats given arguments with newline and returns the result as string.
func (pp *PrettyPrinter) Sprintln(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Fprint prints given arguments to a given writer.
func (pp *PrettyPrinter) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fprintf prints format to a given writer.
func (pp *PrettyPrinter) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fprintln prints given arguments to a given writer with newline.
func (pp *PrettyPrinter) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Errorf formats given arguments and returns it as error type.
func (pp *PrettyPrinter) Errorf(format string, a ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Fatal prints given arguments and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatal(a ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf prints a given format and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatalf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalln prints given arguments with newline and finishes execution with exit status 1.
func (pp *PrettyPrinter) Fatalln(a ...interface{}) { _ = "STUB: not implemented"; return }

func (pp *PrettyPrinter) SetColoringEnabled(enabled bool) { _ = "STUB: not implemented"; return }

func (pp *PrettyPrinter) SetDecimalUint(enabled bool) { _ = "STUB: not implemented"; return }

func (pp *PrettyPrinter) SetExportedOnly(enabled bool) { _ = "STUB: not implemented"; return }

// SetOmitEmpty makes empty fields in struct not be printed.
func (pp *PrettyPrinter) SetOmitEmpty(enabled bool) { _ = "STUB: not implemented"; return }

func (pp *PrettyPrinter) SetThousandsSeparator(enabled bool) { _ = "STUB: not implemented"; return }

// SetOutput sets pp's output
func (pp *PrettyPrinter) SetOutput(o io.Writer) { _ = "STUB: not implemented"; return }

// GetOutput returns pp's output.
func (pp *PrettyPrinter) GetOutput() io.Writer {
	_ = "STUB: not implemented"

	// ResetOutput sets pp's output back to the default output
	return *new(io.Writer)
}

func (pp *PrettyPrinter) ResetOutput() { _ = "STUB: not implemented"; return }

// SetColorScheme takes a colorscheme used by all future Print calls.
func (pp *PrettyPrinter) SetColorScheme(scheme ColorScheme) { _ = "STUB: not implemented"; return }

// ResetColorScheme resets colorscheme to default.
func (pp *PrettyPrinter) ResetColorScheme() { _ = "STUB: not implemented"; return }

// SetMaxDepth sets the printer's Depth, -1 prints all
func (pp *PrettyPrinter) SetMaxDepth(v int) { _ = "STUB: not implemented"; return }

func (pp *PrettyPrinter) formatAll(objects []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (pp *PrettyPrinter) formatAllWithLineFlag(objects []interface{}) ([]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil,

		// fix for backwards capability
		false
}

// Print prints given arguments.
func Print(a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// Printf prints a given format.
		nil
}

func Printf(format string, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Println prints given arguments with newline.
func Println(a ...interface{}) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Sprint formats given arguments and returns the result as string.
func Sprint(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Sprintf formats with pretty print and returns the result as string.
func Sprintf(format string, a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Sprintln formats given arguments with newline and returns the result as string.
func Sprintln(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Fprint prints given arguments to a given writer.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fprintf prints format to a given writer.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fprintln prints given arguments to a given writer with newline.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Errorf formats given arguments and returns it as error type.
func Errorf(format string, a ...interface{}) error { _ = "STUB: not implemented"; return nil }

// Fatal prints given arguments and finishes execution with exit status 1.
func Fatal(a ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf prints a given format and finishes execution with exit status 1.
func Fatalf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalln prints given arguments with newline and finishes execution with exit status 1.
func Fatalln(a ...interface{}) { _ = "STUB: not implemented"; return }

// Change Print* functions' output to a given writer.
// For example, you can limit output by ENV.
//
//	func init() {
//		if os.Getenv("DEBUG") == "" {
//			pp.SetDefaultOutput(ioutil.Discard)
//		}
//	}
func SetDefaultOutput(o io.Writer) { _ = "STUB: not implemented"; return }

// GetOutput returns pp's default output.
func GetDefaultOutput() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// Change Print* functions' output to default one.
func ResetDefaultOutput() { _ = "STUB: not implemented"; return }

// SetColorScheme takes a colorscheme used by all future Print calls.
func SetColorScheme(scheme ColorScheme) { _ = "STUB: not implemented"; return }

// ResetColorScheme resets colorscheme to default.
func ResetColorScheme() { _ = "STUB: not implemented"; return }

// SetMaxDepth sets the printer's Depth, -1 prints all
func SetDefaultMaxDepth(v int) { _ = "STUB: not implemented"; return }
