package rotato

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"syscall"
)

// NBSP represents a non-breaking space character.
const NBSP = "\u00A0"

const (
	cursorHide cursor = "\x1b[?25l"
	cursorShow cursor = "\x1b[?25h"
	clearChars        = "\r\033[K\r"
)

type Term struct {
	clearLine      string
	hideCursor     string
	showCursor     string
	nonInteractive bool
}

func NewTerm() Term {
	return Term{
		clearLine:  clearChars,
		hideCursor: string(cursorHide),
		showCursor: string(cursorShow),
	}
}

// WithNonInteractive returns a copy of t configured to always report
// redirected/non-interactive output, regardless of the actual writer.
func (t Term) WithNonInteractive() Term {
	t.nonInteractive = true
	return t
}

// ClearLine clears line and carriage return.
func (t Term) ClearLine() string {
	return t.clearLine
}

// RemoveANSI removes ANSI codes from a given string.
func (t Term) RemoveANSI(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}

// HideCursor hides the cursor.
func (t Term) HideCursor(output io.Writer) {
	if t.hideCursor != "" && !isRedirected(output, t.nonInteractive) {
		fmt.Fprint(output, t.hideCursor)
	}
}

// ShowCursor shows the cursor.
func (t Term) ShowCursor(output io.Writer) {
	if t.showCursor != "" && !isRedirected(output, t.nonInteractive) {
		fmt.Fprint(output, t.showCursor)
	}
}

// isInteractive checks if the output is interactive.
func isInteractive(r *Rotato) bool {
	return !isRedirected(r.writer, r.term.nonInteractive)
}

// isRedirected checks if the provided output writer is redirected.
// It returns true if the writer is not a terminal.
func isRedirected(output io.Writer, nonInteractive bool) bool {
	if nonInteractive {
		return true
	}
	file, ok := output.(*os.File)
	if !ok {
		// If it's not an *os.File, assume it's redirected,
		return true
	}

	stat, err := file.Stat()
	if err != nil {
		return false
	}
	// Check if the file mode indicates a character device (terminal).
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	// Additional check using syscall.
	var st syscall.Stat_t
	if err := syscall.Fstat(int(file.Fd()), &st); err != nil {
		return false
	}
	// If the mode does not indicate a character device, the output is redirected.
	return (st.Mode & syscall.S_IFMT) != syscall.S_IFCHR
}
