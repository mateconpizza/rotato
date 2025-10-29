package rotato

import "strings"

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

func Colorize(s string, codes ...Color) string {
	var result string
	for _, code := range codes {
		result += string(code)
	}
	return result + s + string(ColorReset)
}

func joinColors(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, c := range colors {
		sb.WriteString(string(c))
	}
	return sb.String()
}
