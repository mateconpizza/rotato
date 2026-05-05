// rotato demo
//
//	narrative demo
//	go run github.com/mateconpizza/rotato/example@latest -demo
//
//	all spinners
//	go run github.com/mateconpizza/rotato/example@latest -all
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/mateconpizza/rotato"
)

var flags = &Flags{}

type Flags struct {
	All            bool
	NonInteractive bool
	Demo           bool
}

// Colors.
var (
	greenBold   = rotato.NewColor(rotato.StyleBold, rotato.FgBrightGreen)
	greenItalic = rotato.NewColor(rotato.FgBrightGreen, rotato.StyleItalic)
)

//nolint:gosec // deterministic demo seed
var rng = rand.New(rand.NewSource(42))

func randFilename() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8+rng.Intn(5))
	for i := range b {
		b[i] = chars[rng.Intn(len(chars))]
	}
	return string(b) + ".zip"
}

// pause blocks until d elapses or ctx is cancelled.
func pause(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// sceneSimple shows the most basic spinner usage: start, wait, done.
func sceneSimple(ctx context.Context) {
	r := rotato.New(
		rotato.WithContext(ctx),
		rotato.WithPrefix("Fetching config"),
		rotato.WithPrefixColor(rotato.StyleItalic),
		rotato.WithSpinnerColor(rotato.FgBrightCyan),
		rotato.WithDoneSymbolColor(greenBold),
		rotato.WithDoneMessageColor(greenItalic),
	)
	r.Start()

	if err := pause(ctx, 2*time.Second); err != nil {
		r.Fail("Aborted")
		return
	}
	r.Done("config.yaml loaded")
}

// sceneMultiStep shows how a single spinner can represent a multi-phase task
// by updating its message and symbol set as work progresses.
func sceneMultiStep(ctx context.Context) {
	r := rotato.New(
		rotato.WithContext(ctx),
		rotato.WithSymbolsCircles6(),
		rotato.WithSpinnerColor(rotato.FgOrange),
		rotato.WithPrefix("S3 backup"),
		rotato.WithPrefixColor(rotato.StyleBold),
		rotato.WithMessage("connecting..."),
		rotato.WithDoneSymbolColor(greenBold),
		rotato.WithDoneMessageColor(rotato.StyleBold),
	)
	r.Start()

	// Phase 1 - connect
	if err := pause(ctx, 2*time.Second); err != nil {
		r.Fail("connection aborted")
		return
	}

	// Phase 2 - confirm connection with a static check mark
	r.UpdateSymbols(rotato.WithSymbols(greenBold.Sprint("✓")))
	r.UpdateMesgColor(greenBold)
	r.UpdateMesg("connected")
	if err := pause(ctx, 800*time.Millisecond); err != nil {
		r.Fail("aborted after connect")
		return
	}

	// Phase 3 - transfer files
	r.UpdateSymbols(rotato.WithSymbolsBarBlock())
	r.UpdateMesgColor(rotato.FgGray)
	r.UpdatePrefix("S3 uploading")
	for i := 1; i <= 15; i++ {
		r.UpdateMesg(fmt.Sprintf("[%d/15] %s", i, randFilename()))
		if err := pause(ctx, 200*time.Millisecond); err != nil {
			r.Fail(fmt.Sprintf("transfer aborted after %d files", i-1))
			return
		}
	}

	r.UpdatePrefix("S3")
	r.Done("15 files uploaded successfully")
}

// sceneError shows the failure path.
func sceneError(ctx context.Context) {
	r := rotato.New(
		rotato.WithContext(ctx),
		rotato.WithPrefix("AWS Health Check"),
		rotato.WithSymbolsPipe(),
		rotato.WithMessage("pinging us-east-1..."),
		rotato.WithSpinnerColor(rotato.FgBlue),
		rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
		rotato.WithFailMessageColor(rotato.FgBrightRed),
	)
	r.Start()

	if err := pause(ctx, 2*time.Second); err != nil {
		r.Fail("aborted by user")
		return
	}

	// Simulate an endpoint that is down.
	r.Fail("connection refused (us-east-1:443)")
}

// sceneCountdown demonstrates a countdown: a graceful shutdown that shows
// remaining time using the library's own formatting helpers.
func sceneCountdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r := rotato.New(
		rotato.WithContext(ctx),
		rotato.WithPrefix("Shutting Down"),
		rotato.WithMessage("in"),
		rotato.WithMessageColor(),
		rotato.WithSymbolsBlockPretty(),
		rotato.WithDoneMessageColor(rotato.StyleBlink, greenBold),
		rotato.WithDoneSymbolColor(rotato.StyleBlink, greenBold),
		rotato.WithMessageDecorator(func(mesg string) string {
			deadline, ok := ctx.Deadline()
			if !ok {
				return mesg
			}
			remaining := max(time.Until(deadline).Round(time.Second), 0)
			return fmt.Sprintf("in %.0fs...", remaining.Seconds())
		}),
	)
	r.Start()

	// Wait for the internal deadline; the spinner stops itself when ctx fires.
	<-ctx.Done()
	r.UpdatePrefix("Demo")
	r.UpdatePrefixColor(rotato.StyleBold)
	r.Done("goodbye!")
}

func padRight(s string, width int) string {
	n := width - utf8.RuneCountInString(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(" ", n)
}

func showSymbols(ctx context.Context) {
	maxLen := 0
	for _, s := range rotato.Spinners() {
		if l := utf8.RuneCountInString(s.Name); l > maxLen {
			maxLen = l
		}
	}

	hint := rotato.FgGray.With(rotato.StyleItalic).Sprint("(ctrl+c to exit)")

	for _, sp := range rotato.Spinners() {
		r := rotato.New(
			rotato.WithContext(ctx),
			rotato.WithPrefix(padRight(sp.Name, maxLen)),
			rotato.WithSpinnerStyle(sp.Name),
			rotato.WithMessage(hint),
			rotato.WithDoneSymbolColor(greenBold),
			rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
		)
		r.Start()

		if err := pause(ctx, 1200*time.Millisecond); err != nil {
			r.Fail("interrupted")
			return
		}
		r.Done(strings.Join(r.Symbols(), " "))
	}
}

func runDemo(ctx context.Context) {
	scenes := []func(context.Context){
		sceneSimple,    // basic start/done
		sceneMultiStep, // live updates across phases
		sceneError,     // failure path
		sceneCountdown, // graceful shutdown with countdown
	}
	for _, scene := range scenes {
		if ctx.Err() != nil {
			return
		}
		scene(ctx)
	}
}

func main() {
	if flags.NonInteractive {
		rotato.SetNonInteractive()
	}

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer stop()

	switch {
	case flags.All:
		showSymbols(ctx)
	case flags.Demo:
		runDemo(ctx)
	default:
		flag.Usage()
	}
}

func init() {
	flag.BoolVar(&flags.All, "all", false, "demo all rotatos")
	flag.BoolVar(&flags.Demo, "demo", false, "run narrative demo")
	flag.BoolVar(&flags.NonInteractive, "ni", false, "non-interactive / dumb-terminal mode")
	flag.Parse()
}
