// Package rotato is a simple spinner library for Go.
//
// Examples:
//
// -------- SIMPLE:
//
//	r := rotato.New(
//		rotato.WithSpinnerColor(rotato.FgBrightGreen),
//		rotato.WithPrefix("Simple Task #1"),
//		rotato.WithDoneMessageColor(rotato.FgBrightGreen, rotato.StyleItalic),
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
//		rotato.WithSymbolsCircles3(), // or rotato.WithSpinnerStyle("circle3"),
//		rotato.WithSpinnerColor(rotato.FgBrightOrange),
//		rotato.WithMessage("Connecting..."),
//		rotato.WithPrefix("S3 Backup"),
//	)
//	r.Start()
//	time.Sleep(2 * time.Second)
//
//	// connected
//	r.UpdateSymbols(rotato.WithSymbols(rotato.FgBrightGreen + "✓"))
//	r.UpdateMesg("Connected!")
//	r.UpdateMesgColor(rotato.FgBrightGreen, rotato.StyleItalic)
//
//	// updating
//	time.Sleep(1 * time.Second)
//	r.UpdateMesgColor(rotato.FgGray)
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
//		rotato.WithMessage("Trying to connect..."),
//		rotato.WithPrefix("AWS Server"),
//		rotato.WithFailMessageColor(rotato.FgBrightRed, rotato.StyleBlink),
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

func WithSpinnerStyle(name string) Option {
	return func(r *Rotato) {
		if sp, ok := ByName(name); ok {
			r.symbols = sp.Frames
		}
	}
}

// WithMessage returns an option function that sets the spinner message.
func WithMessage(s string) Option {
	return func(r *Rotato) {
		r.message = s
	}
}

// WithMessageColor returns an option function that sets the spinner message
// color.
func WithMessageColor(c ...Color) Option {
	return func(r *Rotato) {
		r.messageColor = NewColor(c...)
	}
}

// WithMessageDecorator appends a decorator to transform the spinner message
// before display.
func WithMessageDecorator(fn MessageDecorator) Option {
	return func(r *Rotato) {
		r.messageDecorators = append(r.messageDecorators, fn)
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
		r.prefixColor = NewColor(c...)
	}
}

// WithPrefixDecorator appends a decorator to transform the spinner message
// before display.
func WithPrefixDecorator(fn MessageDecorator) Option {
	return func(r *Rotato) {
		r.prefixDecorators = append(r.prefixDecorators, fn)
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
		r.delimiterColor = NewColor(c...)
	}
}

// WithSpinnerColor returns an option function that sets the spinner color.
func WithSpinnerColor(c ...Color) Option {
	return func(r *Rotato) {
		r.spinnerColor = NewColor(c...)
	}
}

// WithDoneMessage returns an option function that sets the spinner done message.
func WithDoneMessage(mesg string) Option {
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

// WithDoneMessageColor returns an option function that sets the done message
// color.
func WithDoneMessageColor(c ...Color) Option {
	return func(r *Rotato) {
		r.doneMessageColor = NewColor(c...)
	}
}

// WithDoneSymbolColor sets the combined color(s) for the completion symbol.
func WithDoneSymbolColor(c ...Color) Option {
	return func(r *Rotato) {
		r.doneSymbolColor = NewColor(c...)
	}
}

// WithFailSymbol returns an option function that sets the spinner fail symbol.
func WithFailSymbol(symbol string) Option {
	return func(r *Rotato) {
		r.failSymbol = symbol
	}
}

// WithFailMessageColor returns an option function that sets the fail message
// color.
func WithFailMessageColor(c ...Color) Option {
	return func(r *Rotato) {
		r.failMessageColor = NewColor(c...)
	}
}

// WithFailSymbolColor returns an option function that sets the fail symbol
// color.
func WithFailSymbolColor(c ...Color) Option {
	return func(r *Rotato) {
		r.failSymbolColor = NewColor(c...)
	}
}

// WithFrequency returns an option function that sets the spinner frequency.
func WithFrequency(d time.Duration) Option {
	return func(r *Rotato) {
		r.frequency = d
	}
}

// WithWriter returns an option function that sets the spinner writer.
func WithWriter(w io.Writer) Option {
	return func(r *Rotato) {
		r.writer = w
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

// WithContextDoneHandler sets a handler invoked when the context is done.
func WithContextDoneHandler(handler func(*Rotato, error)) Option {
	return func(r *Rotato) {
		r.ctxDoneHandler = handler
	}
}

// MessageDecorator defines a function that transforms a message string.
type MessageDecorator func(mesg string) string

// Option is an option function for the spinner.
type Option func(*Rotato)

// Rotato represents a CLI spinner animation.
type Rotato struct {
	// Output
	writer   io.Writer // Output writer
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
	spinnerColor Color         // Spinner color

	// Messages
	message           string             // Spinner message
	messageUpdate     sync.RWMutex       // Mutex for message update
	messageColor      Color              // Spinner message color
	messageDecorators []MessageDecorator // decorators holds message decorator functions applied before display.

	// Prefix messages
	prefixMesg       string             // Prefix message
	prefixMu         sync.RWMutex       // Synchronization mechanism for prefix updates
	prefixColor      Color              // Prefix message color
	prefixDecorators []MessageDecorator // decorators holds message decorator functions applied before display.

	// Delimiter
	delimiter      string // Delimiter between prefix and spinner symbol
	delimiterColor Color  // Delimiter color

	// Completion states
	doneMessage      string        // Done channel message
	doneMessageColor Color         // Done channel message color
	doneSymbol       string        // Done channel symbol
	doneSymbolColor  Color         // Done symbol color
	doneChan         chan struct{} // Channel for stopping the spinner

	// Failure state
	failMessageColor Color  // Fail message color
	failSymbol       string // Fail symbol
	failSymbolColor  Color  // Fail symbol color

	// Active state
	isActive bool         // State of the spinner
	activeMu sync.RWMutex // Mutex for different spinner states
}

// render displays the current frame and message of the spinner.
func (r *Rotato) render(current int) {
	mesg := r.currentMessage()
	mesg = decorate(mesg, r.messageDecorators)

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

	hideCursor(r.writer)
	r.activeMu.Lock()
	if r.isActive {
		r.activeMu.Unlock()
		return
	}
	r.isActive = true
	r.activeMu.Unlock()

	if isRedirected(r.writer) && !r.forceInteractive {
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
	defer showCursor(r.writer)

	r.stopSpinner()

	var finalMesg string
	switch {
	case len(mesg) > 0:
		finalMesg = strings.Join(mesg, " ")
	case r.doneMessage != "":
		finalMesg = r.doneMessage
	default:
		r.display(clearChars)
		return
	}

	symbol := formatSymbol(r.doneSymbol, r.doneSymbolColor)

	r.displayMessage(symbol, r.doneMessageColor, finalMesg)
}

// Fail fails the spinner animation.
func (r *Rotato) Fail(mesg ...string) {
	r.stopSpinner()
	if len(mesg) == 0 {
		mesg = append(mesg, "Failed")
	}
	symbol := formatSymbol(r.failSymbol, r.failSymbolColor)

	r.displayMessage(symbol, r.failMessageColor, mesg...)
}

// SetWriter updates the output destination safely.
func (r *Rotato) SetWriter(w io.Writer) {
	r.writerMu.Lock()
	defer r.writerMu.Unlock()
	r.writer = w
}

// Writer returns the current output destination.
func (r *Rotato) Writer() io.Writer {
	r.writerMu.Lock()
	defer r.writerMu.Unlock()
	return r.writer
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
		_, _ = fmt.Fprintf(r.writer, "%s\n", mesg)
	}
}

// UpdateMesgColor changes the color of the message.
func (r *Rotato) UpdateMesgColor(c ...Color) {
	r.messageColor = NewColor(c...)
}

// UpdatePrefix changes the prefix shown next to the spinner.
func (r *Rotato) UpdatePrefix(mesg string) {
	r.prefixMu.Lock()
	r.prefixMesg = mesg
	r.prefixMu.Unlock()
}

// UpdatePrefixColor changes the color of the prefix.
func (r *Rotato) UpdatePrefixColor(c ...Color) {
	r.prefixColor = NewColor(c...)
}

// UpdateDoneMesgColor changes the done message shown next to the spinner.
func (r *Rotato) UpdateDoneMesgColor(c ...Color) {
	r.doneMessageColor = NewColor(c...)
}

// UpdateSpinnerColor changes the color of the spinner.
func (r *Rotato) UpdateSpinnerColor(c ...Color) {
	r.spinnerColor = NewColor(c...)
}

// UpdateSymbols updates the spinner symbols.
func (r *Rotato) UpdateSymbols(opt Option) {
	r.activeMu.Lock()
	defer r.activeMu.Unlock()
	opt(r)
}

// currentMessage safely constructs and returns the current message.
func (r *Rotato) currentMessage() string {
	if r.message == "" {
		return ""
	}
	r.messageUpdate.RLock()
	defer r.messageUpdate.RUnlock()

	return r.messageColor.Sprint(r.message)
}

// currentFrame returns the spinner frame for the given iteration.
func (r *Rotato) currentFrame(i int) string {
	if len(r.symbols) == 0 {
		return ""
	}
	r.frameIdx = i % len(r.symbols)
	r.frame = r.symbols[r.frameIdx]

	return r.spinnerColor.Sprint(r.frame)
}

// parsePrefix updates the spinner prefix.
func (r *Rotato) parsePrefix(frame, mesg string) {
	r.prefixMu.RLock()
	defer r.prefixMu.RUnlock()
	prefix := r.prefixColor.Sprint(r.prefixMesg)
	prefix = decorate(prefix, r.prefixDecorators)
	del := r.delimiterColor.Sprint(r.delimiter)

	r.display(fmt.Sprintf("%s%s%s %s", prefix, del, frame, mesg))
}

// display writes the given string to the output.
func (r *Rotato) display(s string) {
	r.writerMu.Lock()
	defer r.writerMu.Unlock()

	if isRedirected(r.writer) {
		_, _ = fmt.Fprint(r.writer, removeANSI(s))
		return
	}

	_, _ = fmt.Fprintf(r.writer, "%s%s", clearChars, s)
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

	defer showCursor(r.writer)
	r.doneChan <- struct{}{}
}

// displayMessage formats and displays a message with optional prefix and color.
func (r *Rotato) displayMessage(symbol string, color Color, mesg ...string) {
	if len(mesg) == 0 {
		return
	}

	msg := strings.Join(mesg, " ")

	if !isInteractive(r) && !r.forceInteractive {
		// the leading \n moves us off the "Loading..." line
		// the trailing \n finishes the log entry
		r.display("\n" + msg + "\n")
		return
	}

	var sb strings.Builder
	// prefix path
	if r.prefixMesg != "" {
		r.buildPrefix(&sb, symbol)
	} else if symbol != "" {
		sb.WriteString(symbol)
		sb.WriteByte(' ')
	}

	// message color
	if color != "" {
		sb.WriteString(color.String())
	}

	sb.WriteString(msg)
	sb.WriteByte('\n')

	// single reset at end
	sb.WriteString(string(ColorReset))

	r.display(sb.String())
}

func (r *Rotato) buildPrefix(sb *strings.Builder, symbol string) {
	r.prefixMu.RLock()
	defer r.prefixMu.RUnlock()

	// prefix
	if r.prefixColor != "" {
		sb.WriteString(r.prefixColor.String())
	}
	sb.WriteString(r.prefixMesg)
	if r.prefixColor != "" {
		sb.WriteString(string(ColorReset))
	}

	// delimiter
	sb.WriteString(r.delimiterColor.Sprint(r.delimiter))

	// symbol
	if symbol != "" {
		sb.WriteString(symbol)
		sb.WriteByte(' ')
	}
}

func (r *Rotato) IsRunning() bool {
	r.activeMu.RLock()
	defer r.activeMu.RUnlock()
	return r.isActive
}

// decorate applies all decorators in order to the given string.
func decorate(s string, decorators []MessageDecorator) string {
	for _, d := range decorators {
		s = d(s)
	}
	return s
}

// formatSymbol wraps a symbol with color codes if provided.
func formatSymbol(symbol string, color Color) string {
	if symbol == "" {
		return ""
	}
	if color == "" {
		return symbol
	}

	return color.Sprint(symbol)
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
		writer:     os.Stdout,
	}
	for _, fn := range opt {
		fn(r)
	}

	if r.ctx == nil {
		r.ctx = context.Background()
	}

	return r
}
