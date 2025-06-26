// Package rotato is a simple spinner library for Go.
//
// Examples:
//
// -------- SIMPLE:
//
//	r := rotato.New(
//		rotato.WithSpinnerColor(rotato.ColorBrightGreen),
//		rotato.WithPrefix("Simple Task #1"),
//		rotato.WithDoneColorMesg(rotato.ColorBrightGreen, rotato.ColorStyleItalic),
//	)
//	r.Start()
//
//	// working...
//	time.Sleep(2 * time.Second)
//	r.Done("Task Completed!")
//
// -------- COLORS, STYLES:
//
//	r := rotato.New(
//		rotato.WithSymbolsCircles3(),
//		rotato.WithSpinnerColor(rotato.ColorBrightOrange),
//		rotato.WithMesg("Connecting..."),
//		rotato.WithPrefix("S3 Backup"),
//	)
//	r.Start()
//	time.Sleep(2 * time.Second)
//
//	// connected
//	r.UpdateSymbols(rotato.WithSymbols(rotato.ColorBrightGreen + "✓"))
//	r.UpdateMesg("Connected!")
//	r.UpdateMesgColor(rotato.ColorBrightGreen, rotato.ColorStyleItalic)
//
//	// updating
//	time.Sleep(1 * time.Second)
//	r.UpdateMesgColor(rotato.ColorGray)
//	r.UpdateSymbols(rotato.WithSymbolsBarBlock())
//	for i := 0; i < 15; i++ {
//		r.UpdateMesg(randomString(12) + ".zip")
//		time.Sleep(200 * time.Millisecond)
//	}
//
//	// end
//	r.Done("Backup completed!")
//
// -------- FAILING:
//
//	r := rotato.New(
//		rotato.WithMesg("Trying to connect..."),
//		rotato.WithPrefix("AWS Server"),
//		rotato.WithFailColorMesg(rotato.ColorBrightRed, rotato.ColorStyleBlink),
//	)
//	r.Start()
//	// trying to connect
//	time.Sleep(2 * time.Second)
//	// fail
//	if true {
//		r.Fail("Connection Failed!")
//	}
//
//	--------------------------------
package rotato

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// nbsp represents a non-breaking space character.
const nbsp = "\u00A0"

var (
	// normal colors.
	ColorBlack   = "\x1b[30m"
	ColorBlue    = "\x1b[34m"
	ColorCyan    = "\x1b[36m"
	ColorGray    = "\x1b[90m"
	ColorGreen   = "\x1b[32m"
	ColorMagenta = "\x1b[95m"
	ColorOrange  = "\x1b[33m"
	ColorPurple  = "\x1b[35m"
	ColorRed     = "\x1b[31m"
	ColorWhite   = "\x1b[37m"
	ColorYellow  = "\x1b[93m"

	// bright colors.
	ColorBrightBlack   = "\x1b[90m"
	ColorBrightBlue    = "\x1b[94m"
	ColorBrightCyan    = "\x1b[96m"
	ColorBrightGray    = "\x1b[37m"
	ColorBrightGreen   = "\x1b[92m"
	ColorBrightMagenta = "\x1b[95m"
	ColorBrightOrange  = "\x1b[38;5;214m"
	ColorBrightPurple  = "\x1b[38;5;135m"
	ColorBrightRed     = "\x1b[91m"
	ColorBrightWhite   = "\x1b[97m"
	ColorBrightYellow  = "\x1b[93m"

	// styles.
	ColorStyleBold          = "\x1b[1m"
	ColorStyleDim           = "\x1b[2m"
	ColorStyleInverse       = "\x1b[7m"
	ColorStyleItalic        = "\x1b[3m"
	ColorStyleStrikethrough = "\x1b[9m"
	ColorStyleUnderline     = "\x1b[4m"
	ColorStyleBlink         = "\x1b[5m"

	// reset.
	ColorReset = "\x1b[0m"
)

// WithMesg returns an option function that sets the spinner message.
func WithMesg(s string) Option {
	return func(r *Rotato) {
		r.message = s
	}
}

// WithMesgColor returns an option function that sets the spinner message
// color.
func WithMesgColor(color ...string) Option {
	return func(r *Rotato) {
		r.messageColor = strings.Join(color, "")
	}
}

// WithPrefix returns an option function that sets the spinner prefix.
func WithPrefix(prefix string) Option {
	return func(r *Rotato) {
		r.prefixMesg = prefix
	}
}

// WithPrefixColor returns an option function that sets the spinner color
// prefix.
func WithPrefixColor(color ...string) Option {
	return func(r *Rotato) {
		r.prefixColor = strings.Join(color, "")
	}
}

// WithDoneMesg returns an option function that sets the spinner done message.
func WithDoneMesg(mesg string) Option {
	return func(r *Rotato) {
		r.doneMessage = mesg
	}
}

// WithDoneSymbol returns an option function that sets the spinner stop symbol.
func WithDoneSymbol(symbol string) Option {
	return func(r *Rotato) {
		r.doneSymbol = symbol
	}
}

// WithDoneColorMesg returns an option function that sets the done message
// color.
func WithDoneColorMesg(color ...string) Option {
	return func(r *Rotato) {
		r.doneMessageColor = strings.Join(color, "")
	}
}

// WithFailSymbol returns an option function that sets the spinner fail symbol.
func WithFailSymbol(symbol string) Option {
	return func(r *Rotato) {
		r.failSymbol = symbol
	}
}

// WithFailColorMesg returns an option function that sets the fail message
// color.
func WithFailColorMesg(color ...string) Option {
	return func(r *Rotato) {
		r.failMessageColor = strings.Join(color, "")
	}
}

// WithSpinnerColor returns an option function that sets the spinner color.
func WithSpinnerColor(color ...string) Option {
	return func(r *Rotato) {
		r.spinnerColor = strings.Join(color, "")
	}
}

// WithSpinnerFrequency returns an option function that sets the spinner frequency.
func WithSpinnerFrequency(d time.Duration) Option {
	return func(r *Rotato) {
		r.frequency = d
	}
}

// WithDelimiter returns an option function that sets the spinner delimiter.
func WithDelimiter(s string) Option {
	return func(r *Rotato) {
		r.delimiter = s
	}
}

// WithDelimiterColor returns an option function that sets the spinner color
// delimiter, only visible with `prefix`.
func WithDelimiterColor(color ...string) Option {
	return func(r *Rotato) {
		r.delimiterColor = strings.Join(color, "")
	}
}

// WithWriter returns an option function that sets the spinner writer.
func WithWriter(w io.Writer) Option {
	return func(r *Rotato) {
		r.Writer = w
	}
}

// Option is an option function for the spinner.
type Option func(*Rotato)

// Rotato represents a CLI spinner animation.
type Rotato struct {
	Writer           io.Writer     // Output writer
	delimiter        string        // Delimiter between prefix and spinner symbol
	delimiterColor   string        // Delimiter color
	doneMessage      string        // Done channel message
	doneChan         chan struct{} // Channel for stopping the spinner
	doneMessageColor string        // Done channel message color
	doneSymbol       string        // Done channel symbol
	failMessageColor string        // Fail message color
	failSymbol       string        // Fail symbol
	frame            string        // Current spinner frame
	frameIdx         int           // Current spinner frame index
	frequency        time.Duration // Spinner animation frequency
	isActive         bool          // State of the spinner
	message          string        // Spinner message
	messageColor     string        // Spinner message color
	messageUpdate    sync.RWMutex  // Mutex for message update
	mu               *sync.RWMutex // Mutex for different spinner states
	prefixColor      string        // Prefix message color
	prefixMesg       string        // Prefix message
	prefixMu         sync.RWMutex  // Synchronization mechanism for prefix updates.
	spinnerColor     string        // Spinner color
	symbols          []string      // Spinner symbols
}

// render displays the current frame and message of the spinner.
func (r *Rotato) render(current int) {
	mesg := r.currentMessage()
	frameFormatted := r.currentFrame(current)

	if r.prefixMesg != "" {
		r.parsePrefix(frameFormatted, mesg)
		return
	}
	r.display(fmt.Sprintf("%s %s", frameFormatted, mesg))
}

// Start starts the spinning animation in a goroutine.
func (r *Rotato) Start() {
	if !isInteractive(r) {
		r.mu.Lock()
		defer r.mu.Unlock()

		if r.isActive {
			return
		}

		r.isActive = true
		// add prefix
		if r.prefixMesg != "" {
			r.message = fmt.Sprintf("%s%s%s", r.prefixMesg, r.delimiter, r.message)
		}
		r.display(r.message)

		return
	}

	hideCursor(r.Writer)
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.isActive {
		return
	}

	r.isActive = true
	if isRedirected(r.Writer) {
		r.render(0)
		return
	}

	ticker := time.NewTicker(r.frequency)
	go func() {
		defer ticker.Stop()

		for i := 0; ; i++ {
			select {
			case <-r.doneChan:
				return
			case <-ticker.C:
				r.mu.Lock()
				if !r.isActive {
					r.mu.Unlock()
					return
				}
				r.render(i)
				r.mu.Unlock()
			}
		}
	}()
}

// Done stops the spinner animation.
func (r *Rotato) Done(mesg ...string) {
	if !r.isActive {
		return
	}
	defer showCursor(r.Writer)

	r.stopSpinner()

	var finalMesg string
	switch {
	case len(mesg) > 0:
		finalMesg = strings.Join(mesg, " ")
	case r.doneMessage != "":
		finalMesg = r.doneMessage
	default:
		fmt.Print(clearChars)
		return
	}

	r.displayMessage(r.doneSymbol, r.doneMessageColor, finalMesg)
}

// Fail fails the spinner animation.
func (r *Rotato) Fail(mesg ...string) {
	if !r.isActive {
		return
	}
	r.stopSpinner()
	if len(mesg) == 0 {
		mesg = append(mesg, "Failed")
	}
	r.displayMessage(r.failSymbol, r.failMessageColor, mesg...)
}

// Symbols returns the spinner symbols.
func (r *Rotato) Symbols() []string {
	return r.symbols
}

// UpdateMesg changes the message shown next to the spinner.
func (r *Rotato) UpdateMesg(mesg string) {
	r.messageUpdate.Lock()
	r.message = mesg
	r.messageUpdate.Unlock()
	if !isInteractive(r) {
		_, _ = fmt.Fprintf(r.Writer, "%s\n", mesg)
	}
}

// UpdateMesgColor changes the color of the message.
func (r *Rotato) UpdateMesgColor(color ...string) {
	r.messageColor = strings.Join(color, "")
}

// UpdatePrefix changes the prefix shown next to the spinner.
func (r *Rotato) UpdatePrefix(mesg string) {
	r.prefixMu.Lock()
	r.prefixMesg = mesg
	r.prefixMu.Unlock()
}

// UpdatePrefixColor changes the color of the prefix.
func (r *Rotato) UpdatePrefixColor(color ...string) {
	r.prefixColor = strings.Join(color, "")
}

// UpdateDoneMesgColor changes the done message shown next to the spinner.
func (r *Rotato) UpdateDoneMesgColor(mesg ...string) {
	r.doneMessageColor = strings.Join(mesg, "")
}

// UpdateSpinnerColor changes the color of the spinner.
func (r *Rotato) UpdateSpinnerColor(color ...string) {
	r.spinnerColor = strings.Join(color, "")
}

// UpdateSymbols updates the spinner symbols.
func (r *Rotato) UpdateSymbols(opt Option) {
	r.mu.Lock()
	opt(r)
	r.mu.Unlock()
}

// currentMessage safely constructs and returns the current message.
func (r *Rotato) currentMessage() string {
	if r.message == "" {
		return ""
	}
	r.messageUpdate.RLock()
	defer r.messageUpdate.RUnlock()

	return r.messageColor + r.message + ColorReset
}

// currentFrame returns the spinner frame for the given iteration.
func (r *Rotato) currentFrame(i int) string {
	if len(r.symbols) == 0 {
		return ""
	}
	r.frameIdx = i % len(r.symbols)
	r.frame = r.symbols[r.frameIdx]

	return r.spinnerColor + r.frame + ColorReset
}

// parsePrefix updates the spinner prefix.
func (r *Rotato) parsePrefix(frame, mesg string) {
	r.prefixMu.RLock()
	prefix := r.prefixColor + r.prefixMesg + ColorReset
	r.prefixMu.RUnlock()
	del := r.delimiterColor + r.delimiter + ColorReset

	r.display(fmt.Sprintf("%s%s%s %s", prefix, del, frame, mesg))
}

// display writes the given string to the output.
func (r *Rotato) display(s string) {
	if isRedirected(r.Writer) {
		_, _ = fmt.Fprint(r.Writer, removeANSI(s))
		return
	}
	_, _ = fmt.Fprintf(r.Writer, "%s%s", clearChars, s)
}

// stopSpinner handles the common logic for stopping the spinner.
func (r *Rotato) stopSpinner() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.isActive {
		return
	}

	r.isActive = false

	if !isInteractive(r) {
		return
	}

	defer showCursor(r.Writer)
	r.doneChan <- struct{}{}
}

// displayMessage formats and displays a message with optional prefix and color.
func (r *Rotato) displayMessage(symbol, color string, mesg ...string) {
	if len(mesg) == 0 {
		return
	}

	s := strings.Join(mesg, " ")
	s = color + s + "\n"

	if !isInteractive(r) {
		r.display(s)
		return
	}

	if r.prefixMesg != "" {
		r.parsePrefix(symbol, s)
		fmt.Print(ColorReset)
		return
	}

	r.display(symbol + " " + s + ColorReset)
}

// removeANSI removes ANSI codes from a given string.
func removeANSI(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}

// New returns a new spinner.
func New(opt ...Option) *Rotato {
	r := &Rotato{
		frequency:  100 * time.Millisecond,
		delimiter:  nbsp,
		isActive:   false,
		message:    "Loading...",
		mu:         &sync.RWMutex{},
		prefixMesg: "",
		doneChan:   make(chan struct{}, 1),
		doneSymbol: "✓",
		failSymbol: "✗",
		symbols:    defaultSymbols,
		Writer:     os.Stdout,
	}
	for _, fn := range opt {
		fn(r)
	}

	setupInterruptHandler(context.Background(), func() {
		showCursor(r.Writer)
	})

	return r
}
