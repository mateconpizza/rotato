package rotato

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

// NBSP represents a non-breaking space character.
const NBSP = "\u00A0"

// nonInteractive indicates whether the terminal is non-interactive.
var nonInteractive = false

// SetNonInteractive sets the terminal to non-interactive mode.
func SetNonInteractive() {
	nonInteractive = true
}

// hideCursor hides the cursor.
func hideCursor(output io.Writer) {
	if !isRedirected(output) {
		fmt.Print(CursorHide)
	}
}

// showCursor shows the cursor.
func showCursor(output io.Writer) {
	if !isRedirected(output) {
		fmt.Print(CursorShow)
	}
}

// isInteractive checks if the output is interactive.
func isInteractive(r *Rotato) bool {
	return !isRedirected(r.Writer)
}

// isRedirected checks if the provided output writer is redirected.
// It returns true if the writer is not a terminal.
func isRedirected(output io.Writer) bool {
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
