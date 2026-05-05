package rotato

import (
	"fmt"
	"os"
	"strings"
)

type (
	// cursor control sequences for showing, hiding the cursor.
	cursor string

	// Color (Select Graphic Rendition) sequences for colors and text styles.
	Color string
)

const (
	// Hide the cursor.
	CursorHide cursor = "\x1b[?25l"

	// Show the cursor.
	CursorShow cursor = "\x1b[?25h"

	// clearChars represents a sequence of characters used to clear the current
	// line in the terminal.
	clearChars = "\r\033[K\r"
)

const (
	ColorReset Color = "\x1b[0m" // Reset all attributes

	// Standard foreground colors (30-37).
	ColorBlack   Color = "\x1b[30m"
	ColorRed     Color = "\x1b[31m"
	ColorGreen   Color = "\x1b[32m"
	ColorYellow  Color = "\x1b[33m"
	ColorBlue    Color = "\x1b[34m"
	ColorMagenta Color = "\x1b[35m"
	ColorCyan    Color = "\x1b[36m"
	ColorWhite   Color = "\x1b[37m"

	// Bright foreground colors (90-97).
	ColorBrightBlack   Color = "\x1b[90m"
	ColorBrightRed     Color = "\x1b[91m"
	ColorBrightGreen   Color = "\x1b[92m"
	ColorBrightYellow  Color = "\x1b[93m"
	ColorBrightBlue    Color = "\x1b[94m"
	ColorBrightMagenta Color = "\x1b[95m"
	ColorBrightCyan    Color = "\x1b[96m"
	ColorBrightWhite   Color = "\x1b[97m"

	// Aliases for common colors.
	ColorGray   Color = "\x1b[90m"       // Alias for bright black
	ColorOrange Color = "\x1b[38;5;214m" // 256-color orange
	ColorPurple Color = "\x1b[38;5;135m" // 256-color purple

	// Standard background colors (40-47).
	ColorBgBlack   Color = "\x1b[40m"
	ColorBgRed     Color = "\x1b[41m"
	ColorBgGreen   Color = "\x1b[42m"
	ColorBgYellow  Color = "\x1b[43m"
	ColorBgBlue    Color = "\x1b[44m"
	ColorBgMagenta Color = "\x1b[45m"
	ColorBgCyan    Color = "\x1b[46m"
	ColorBgWhite   Color = "\x1b[47m"

	// Bright background colors (100-107).
	ColorBgBrightBlack   Color = "\x1b[100m"
	ColorBgBrightRed     Color = "\x1b[101m"
	ColorBgBrightGreen   Color = "\x1b[102m"
	ColorBgBrightYellow  Color = "\x1b[103m"
	ColorBgBrightBlue    Color = "\x1b[104m"
	ColorBgBrightMagenta Color = "\x1b[105m"
	ColorBgBrightCyan    Color = "\x1b[106m"
	ColorBgBrightWhite   Color = "\x1b[107m"

	// Text styles.
	ColorStyleBold          Color = "\x1b[1m" // Bold or increased intensity
	ColorStyleDim           Color = "\x1b[2m" // Faint or dim
	ColorStyleItalic        Color = "\x1b[3m" // Italic
	ColorStyleUnderline     Color = "\x1b[4m" // Underline
	ColorStyleBlink         Color = "\x1b[5m" // Slow blink
	ColorStyleBlinkRapid    Color = "\x1b[6m" // Rapid blink
	ColorStyleInverse       Color = "\x1b[7m" // Inverse/reverse video
	ColorStyleHidden        Color = "\x1b[8m" // Conceal/hidden
	ColorStyleStrikethrough Color = "\x1b[9m" // Crossed-out/strikethrough
)

// isColorDisabled checks the NO_COLOR environment variable.
// https://no-color.org/
func isColorDisabled() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

// Wrap wraps the given text with the provided styles and resets afterwards.
func (c Color) Wrap(text string, styles ...Color) string {
	if isColorDisabled() || nonInteractive {
		return text
	}
	return string(c) + combine(styles...) + text + string(ColorReset)
}

// With combines the receiver style with additional styles and returns a new
// Color value.
func (c Color) With(styles ...Color) Color {
	return Color(string(c) + combine(styles...))
}

// Sprint wraps the formatted text with the receiver style and returns it as a
// string.
func (c Color) Sprint(a ...any) string {
	return c.Wrap(fmt.Sprint(a...))
}

// Sprintf wraps the formatted text using the provided format string with the
// receiver style and returns it as a string.
func (c Color) Sprintf(f string, a ...any) string {
	return c.Wrap(fmt.Sprintf(f, a...))
}

func (c Color) String() string {
	if isColorDisabled() {
		return ""
	}

	return string(c)
}

// combine merges multiple Color codes into a single string.
func combine(codes ...Color) string {
	if len(codes) == 0 {
		return ""
	}
	if len(codes) == 1 {
		return string(codes[0])
	}

	var sb strings.Builder
	for _, code := range codes {
		sb.WriteString(string(code))
	}
	return sb.String()
}

func NewColor(codes ...Color) Color {
	return Color(combine(codes...))
}
