package rotato

import (
	"fmt"
	"os"
	"strings"
)

type (
	// cursor control sequences for showing, hiding the cursor.
	cursor string

	// Color in a sequences for colors and text styles.
	Color string
)

func NewColor(codes ...Color) Color {
	return Color(combine(codes...))
}

// isColorDisabled checks the NO_COLOR environment variable.
// https://no-color.org/
func isColorDisabled() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

const (
	ColorReset Color = "\x1b[0m" // Reset all attributes

	// Standard foreground colors (30-37).
	FgBlack   Color = "\x1b[30m"
	FgRed     Color = "\x1b[31m"
	FgGreen   Color = "\x1b[32m"
	FgYellow  Color = "\x1b[33m"
	FgBlue    Color = "\x1b[34m"
	FgMagenta Color = "\x1b[35m"
	FgCyan    Color = "\x1b[36m"
	FgWhite   Color = "\x1b[37m"

	// Bright foreground colors (90-97).
	FgBrightBlack   Color = "\x1b[90m"
	FgBrightRed     Color = "\x1b[91m"
	FgBrightGreen   Color = "\x1b[92m"
	FgBrightYellow  Color = "\x1b[93m"
	FgBrightBlue    Color = "\x1b[94m"
	FgBrightMagenta Color = "\x1b[95m"
	FgBrightCyan    Color = "\x1b[96m"
	FgBrightWhite   Color = "\x1b[97m"

	// Aliases for common colors.
	FgGray   Color = "\x1b[90m"       // Alias for bright black
	FgOrange Color = "\x1b[38;5;214m" // 256-color orange
	FgPurple Color = "\x1b[38;5;135m" // 256-color purple

	// Standard background colors (40-47).
	BgBlack   Color = "\x1b[40m"
	BgRed     Color = "\x1b[41m"
	BgGreen   Color = "\x1b[42m"
	BgYellow  Color = "\x1b[43m"
	BgBlue    Color = "\x1b[44m"
	BgMagenta Color = "\x1b[45m"
	BgCyan    Color = "\x1b[46m"
	BgWhite   Color = "\x1b[47m"

	// Bright background colors (100-107).
	BgBrightBlack   Color = "\x1b[100m"
	BgBrightRed     Color = "\x1b[101m"
	BgBrightGreen   Color = "\x1b[102m"
	BgBrightYellow  Color = "\x1b[103m"
	BgBrightBlue    Color = "\x1b[104m"
	BgBrightMagenta Color = "\x1b[105m"
	BgBrightCyan    Color = "\x1b[106m"
	BgBrightWhite   Color = "\x1b[107m"

	// Text styles.
	StyleBold          Color = "\x1b[1m" // Bold or increased intensity
	StyleDim           Color = "\x1b[2m" // Faint or dim
	StyleItalic        Color = "\x1b[3m" // Italic
	StyleUnderline     Color = "\x1b[4m" // Underline
	StyleBlink         Color = "\x1b[5m" // Slow blink
	StyleBlinkRapid    Color = "\x1b[6m" // Rapid blink
	StyleInverse       Color = "\x1b[7m" // Inverse/reverse video
	StyleHidden        Color = "\x1b[8m" // Conceal/hidden
	StyleStrikethrough Color = "\x1b[9m" // Crossed-out/strikethrough
)

// Wrap wraps the given text with the provided styles and resets afterwards.
func (c Color) Wrap(text string, styles ...Color) string {
	if isColorDisabled() {
		return text
	}
	return string(c) + combine(styles...) + text + ColorReset.String()
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
		sb.WriteString(code.String())
	}
	return sb.String()
}

type colorFormatFunc func(s string) string

type Colorizer struct {
	spinner     Color
	message     Color
	prefixMesg  Color
	delimiter   Color
	doneMessage Color
	doneSymbol  Color
	failMessage Color
	failSymbol  Color

	Enabled bool
}

func newColorizer() *Colorizer {
	return &Colorizer{Enabled: !isColorDisabled()}
}

func (c *Colorizer) Format(color Color, text string) string {
	if !c.Enabled || text == "" || color == "" {
		return text
	}
	return color.Sprint(text)
}

func (c *Colorizer) Spinner(s string) string    { return c.Format(c.spinner, s) }     // Spinner color
func (c *Colorizer) Message(s string) string    { return c.Format(c.message, s) }     // Spinner message color
func (c *Colorizer) Prefix(s string) string     { return c.Format(c.prefixMesg, s) }  // Prefix message color
func (c *Colorizer) Delimiter(s string) string  { return c.Format(c.delimiter, s) }   // Delimiter color
func (c *Colorizer) DoneMsg(s string) string    { return c.Format(c.doneMessage, s) } // Done channel message color
func (c *Colorizer) DoneSymbol(s string) string { return c.Format(c.doneSymbol, s) }  // Done symbol color
func (c *Colorizer) FailMsg(s string) string    { return c.Format(c.failMessage, s) } // Fail message color
func (c *Colorizer) FailSymbol(s string) string { return c.Format(c.failSymbol, s) }  // Fail symbol color
