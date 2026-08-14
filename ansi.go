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

type Palette struct {
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

func newPalette() *Palette {
	_, exists := os.LookupEnv("NO_COLOR")
	return &Palette{Enabled: !exists}
}

func (p *Palette) Format(c Color, text string) string {
	if !p.Enabled || text == "" || c == "" {
		return text
	}
	return c.Sprint(text)
}

func (p *Palette) Spinner(s string) string    { return p.Format(p.spinner, s) }     // Spinner color
func (p *Palette) Message(s string) string    { return p.Format(p.message, s) }     // Spinner message color
func (p *Palette) Prefix(s string) string     { return p.Format(p.prefixMesg, s) }  // Prefix message color
func (p *Palette) Delimiter(s string) string  { return p.Format(p.delimiter, s) }   // Delimiter color
func (p *Palette) DoneMsg(s string) string    { return p.Format(p.doneMessage, s) } // Done channel message color
func (p *Palette) DoneSymbol(s string) string { return p.Format(p.doneSymbol, s) }  // Done symbol color
func (p *Palette) FailMsg(s string) string    { return p.Format(p.failMessage, s) } // Fail message color
func (p *Palette) FailSymbol(s string) string { return p.Format(p.failSymbol, s) }  // Fail symbol color
