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

// WithMesg returns an option function that sets the spinner message.
func WithMesg(s string) Option {
	return func(r *Rotato) {
		r.message = s
	}
}

// WithMesgColor returns an option function that sets the spinner message
// color.
func WithMesgColor(c ...Color) Option {
	return func(r *Rotato) {
		r.messageColor = joinColors(c...)
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
func WithPrefixColor(c ...Color) Option {
	return func(r *Rotato) {
		r.prefixColor = joinColors(c...)
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
func WithDoneColorMesg(c ...Color) Option {
	return func(r *Rotato) {
		r.doneMessageColor = joinColors(c...)
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
func WithFailColorMesg(c ...Color) Option {
	return func(r *Rotato) {
		r.failMessageColor = joinColors(c...)
	}
}

func WithFailColorSymbol(c ...Color) Option {
	return func(r *Rotato) {
		r.failSymbolColor = joinColors(c...)
	}
}

// WithSpinnerColor returns an option function that sets the spinner color.
func WithSpinnerColor(c ...Color) Option {
	return func(r *Rotato) {
		r.spinnerColor = joinColors(c...)
	}
}

// WithFrequency returns an option function that sets the spinner frequency.
func WithFrequency(d time.Duration) Option {
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
func WithDelimiterColor(c ...Color) Option {
	return func(r *Rotato) {
		r.delimiterColor = joinColors(c...)
	}
}

// WithWriter returns an option function that sets the spinner writer.
func WithWriter(w io.Writer) Option {
	return func(r *Rotato) {
		r.Writer = w
	}
}

func WithContext(ctx context.Context) Option {
	return func(r *Rotato) {
		r.ctx = ctx
	}
}

func WithForceInteractive() Option {
	return func(r *Rotato) {
		r.forceInteractive = true
	}
}

// WithMesgDecorator appends a decorator to transform the spinner message
// before display.
func WithMesgDecorator(fn mesgDecorator) Option {
	return func(r *Rotato) {
		r.decorators = append(r.decorators, fn)
	}
}

// WithContextDoneHandler sets a handler invoked when the context is done.
func WithContextDoneHandler(handler func(*Rotato, error)) Option {
	return func(r *Rotato) {
		r.ctxDoneHandler = handler
	}
}

// mesgDecorator defines a function that transforms a message string.
type mesgDecorator func(mesg string) string

// Option is an option function for the spinner.
type Option func(*Rotato)

// Rotato represents a CLI spinner animation.
type Rotato struct {
	// Output
	Writer   io.Writer // Output writer
	writerMu sync.Mutex

	// Context
	ctx            context.Context
	ctxDoneHandler func(r *Rotato, err error)

	// for Testing
	forceInteractive bool

	// Spinner animation
	symbols      []string      // Spinner symbols
	frame        string        // Current spinner frame
	frameIdx     int           // Current spinner frame index
	frequency    time.Duration // Spinner animation frequency
	spinnerColor string        // Spinner color

	// Messages
	message       string       // Spinner message
	messageUpdate sync.RWMutex // Mutex for message update
	messageColor  string       // Spinner message color

	// Prefix messages
	prefixMesg  string       // Prefix message
	prefixMu    sync.RWMutex // Synchronization mechanism for prefix updates
	prefixColor string       // Prefix message color

	// Delimiter
	delimiter      string // Delimiter between prefix and spinner symbol
	delimiterColor string // Delimiter color

	// Completion states
	doneMessage      string        // Done channel message
	doneMessageColor string        // Done channel message color
	doneSymbol       string        // Done channel symbol
	doneChan         chan struct{} // Channel for stopping the spinner

	// Failure state
	failMessageColor string // Fail message color
	failSymbol       string // Fail symbol
	failSymbolColor  string // Fail symbol color

	// decorators holds message decorator functions applied before display.
	decorators []mesgDecorator

	// Active state
	isActive bool         // State of the spinner
	activeMu sync.RWMutex // Mutex for different spinner states
}

// render displays the current frame and message of the spinner.
func (r *Rotato) render(current int) {
	mesg := r.currentMessage()
	mesg = r.decorateMessage(mesg)

	frameFormatted := r.currentFrame(current)

	if r.prefixMesg != "" {
		r.parsePrefix(frameFormatted, mesg)
		return
	}

	r.display(fmt.Sprintf("%s %s", frameFormatted, mesg))
}

// Start starts the spinning animation in a goroutine.
func (r *Rotato) Start() {
	if !isInteractive(r) && !r.forceInteractive {
		r.activeMu.Lock()
		defer r.activeMu.Unlock()

		if r.isActive {
			return
		}

		r.isActive = true

		if r.prefixMesg != "" {
			r.message = fmt.Sprintf("%s%s%s", r.prefixMesg, r.delimiter, r.message)
		}
		r.display(r.message)

		return
	}

	hideCursor(r.Writer)
	r.activeMu.Lock()
	if r.isActive {
		r.activeMu.Unlock()
		return
	}
	r.isActive = true
	r.activeMu.Unlock()

	if isRedirected(r.Writer) && !r.forceInteractive {
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
			case <-r.ctx.Done():
				if r.ctxDoneHandler != nil {
					r.ctxDoneHandler(r, r.ctx.Err())
				}

				r.stopSpinner()
				return
			case <-ticker.C:
				r.activeMu.Lock()
				if !r.isActive {
					r.activeMu.Unlock()
					return
				}
				r.render(i)
				r.activeMu.Unlock()
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
func (r *Rotato) UpdateMesgColor(c ...Color) {
	r.messageColor = joinColors(c...)
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
	r.activeMu.Lock()
	opt(r)
	r.activeMu.Unlock()
}

// currentMessage safely constructs and returns the current message.
func (r *Rotato) currentMessage() string {
	if r.message == "" {
		return ""
	}
	r.messageUpdate.RLock()
	defer r.messageUpdate.RUnlock()

	return r.messageColor + r.message + string(ColorReset)
}

// currentFrame returns the spinner frame for the given iteration.
func (r *Rotato) currentFrame(i int) string {
	if len(r.symbols) == 0 {
		return ""
	}
	r.frameIdx = i % len(r.symbols)
	r.frame = r.symbols[r.frameIdx]

	return r.spinnerColor + r.frame + string(ColorReset)
}

// parsePrefix updates the spinner prefix.
func (r *Rotato) parsePrefix(frame, mesg string) {
	r.prefixMu.RLock()
	prefix := r.prefixColor + r.prefixMesg + string(ColorReset)
	r.prefixMu.RUnlock()
	del := r.delimiterColor + r.delimiter + string(ColorReset)

	r.display(fmt.Sprintf("%s%s%s %s", prefix, del, frame, mesg))
}

// display writes the given string to the output.
func (r *Rotato) display(s string) {
	r.writerMu.Lock()
	defer r.writerMu.Unlock()

	if isRedirected(r.Writer) {
		_, _ = fmt.Fprint(r.Writer, removeANSI(s))
		return
	}

	_, _ = fmt.Fprintf(r.Writer, "%s%s", clearChars, s)
}

// stopSpinner handles the common logic for stopping the spinner.
func (r *Rotato) stopSpinner() {
	r.activeMu.Lock()
	defer r.activeMu.Unlock()

	if !r.isActive {
		return
	}

	r.isActive = false

	if !isInteractive(r) && !r.forceInteractive {
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

	if !isInteractive(r) && !r.forceInteractive {
		r.display(s)
		return
	}

	if r.prefixMesg != "" {
		r.parsePrefix(symbol, s)
		fmt.Print(string(ColorReset))
		return
	}

	var result string
	if symbol != "" {
		result = symbol + " "
	}

	r.display(result + s + string(ColorReset))
}

func (r *Rotato) IsRunning() bool {
	r.activeMu.RLock()
	defer r.activeMu.RUnlock()
	return r.isActive
}

func (r *Rotato) decorateMessage(mesg string) string {
	for _, d := range r.decorators {
		mesg = d(mesg)
	}
	return mesg
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
		delimiter:  NBSP,
		isActive:   false,
		message:    "Loading...",
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

	if r.ctx == nil {
		r.ctx = context.Background()
	}

	return r
}
