// rotato demo
//
//	narrative demo
//	go run github.com/mateconpizza/rotato/example@latest -demo
//
//	all spinners
//	go run github.com/mateconpizza/rotato/example@latest -all
//
//	more spinners
//	go run github.com/mateconpizza/rotato/example@latest -more
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
	List           bool
	Group          string
	Show           string
	More           bool
	Duration       time.Duration
}

// Colors.
var (
	greenBold   = rotato.NewColor(rotato.StyleBold, rotato.FgBrightGreen)
	greenItalic = rotato.NewColor(rotato.FgBrightGreen, rotato.StyleItalic)
	hint        = rotato.StyleDim.With(rotato.StyleItalic).Sprint("(ctrl+c to exit)")
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
func sceneSimple(ctx context.Context, d time.Duration) {
	r := rotato.New(
		rotato.WithPrefix("Fetching config"),
		rotato.WithPrefixColor(rotato.StyleItalic),
		rotato.WithSpinnerColor(rotato.FgBrightCyan),
		rotato.WithDoneSymbolColor(greenBold),
		rotato.WithDoneMessageColor(greenItalic),
	)
	r.Start(ctx)

	if err := pause(ctx, d); err != nil {
		r.Fail("Aborted")
		return
	}
	r.Done("config.yaml loaded")
}

// sceneMultiStep shows how a single spinner can represent a multi-phase task
// by updating its message and symbol set as work progresses.
func sceneMultiStep(ctx context.Context, d time.Duration) {
	r := rotato.New(
		rotato.WithSymbolsCircles6(),
		rotato.WithSpinnerColor(rotato.FgOrange),
		rotato.WithPrefix("S3 backup"),
		rotato.WithPrefixColor(rotato.StyleBold),
		rotato.WithMessage("connecting..."),
		rotato.WithDoneSymbolColor(greenBold),
		rotato.WithDoneMessageColor(rotato.StyleBold),
	)
	r.Start(ctx)

	// Phase 1 - connect
	if err := pause(ctx, d); err != nil {
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
func sceneError(ctx context.Context, d time.Duration) {
	r := rotato.New(
		rotato.WithPrefix("AWS Health Check"),
		rotato.WithSymbolsPipe(),
		rotato.WithSpinnerColor(rotato.FgBlue.With(rotato.StyleBold)),
		rotato.WithMessage("pinging us-east-1..."),
		rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
		rotato.WithFailMessageColor(rotato.FgBrightRed),
	)
	r.Start(ctx)

	if err := pause(ctx, d); err != nil {
		r.Fail("aborted by user")
		return
	}

	// Simulate an endpoint that is down.
	r.Fail("connection refused (us-east-1:443)")
}

// sceneCountdown demonstrates a countdown: a graceful shutdown that shows
// remaining time using the library's own formatting helpers.
func sceneCountdown(ctx context.Context, d time.Duration) {
	var color rotato.Color
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	r := rotato.New(
		rotato.WithPrefix("Shutting Down"),
		rotato.WithPrefixColor(rotato.StyleBold),
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
			remaining := max(time.Until(deadline), 0)
			switch {
			case remaining > 3*time.Second:
				color = rotato.FgBrightGreen
			case remaining > time.Second:
				color = rotato.FgOrange
			default:
				color = rotato.FgBrightRed
			}

			return color.With(rotato.StyleItalic).Sprintf(
				"in %s...",
				formatDuration(remaining),
			)
		}),
	)
	r.Start(ctx)

	// Wait for the internal deadline; the spinner stops itself when ctx fires.
	<-ctx.Done()
	r.UpdatePrefix("Demo")
	r.UpdatePrefixColor(rotato.StyleBold)
	r.UpdateDoneSymbol("\u25A0")
	r.Done("goodbye!")
}

// sceneElapse runs sequential spinner tasks (download, verify, extract) and
// shows elapsed time via a shared prefix decorator with a timeout.
func sceneElapse(ctx context.Context, d time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, 12*time.Second)
	defer cancel()

	task := "Work"
	t := time.Now()
	formatter := func(mesg string) string {
		d := time.Since(t).Round(time.Second)
		return mesg + " " + rotato.DimElapsedDecorator(d)
	}

	// Step 1: Download
	download := rotato.New(
		rotato.WithSpinnerColor(rotato.FgBrightGreen),
		rotato.WithDoneSymbol("\u203A"),
		rotato.WithDoneSymbolColor(rotato.FgBrightCyan),
		rotato.WithSymbolsBarBlock7(),
		rotato.WithFrequency(110*time.Millisecond),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Downloading chunks..."),
	)
	download.Start(ctx)
	if err := pause(ctx, d); err != nil {
		download.Fail("Aborted")
		return
	}
	download.Done("Download complete")

	// Step 2: Verify
	verify := rotato.New(
		rotato.WithDoneSymbol("\u203A"),
		rotato.WithDoneSymbolColor(rotato.FgBrightYellow),
		rotato.WithSymbolsCircles7(),
		rotato.WithSpinnerColor(rotato.FgBrightCyan),
		rotato.WithFrequency(110*time.Millisecond),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Verifying checksum..."),
	)
	verify.Start(ctx)
	if err := pause(ctx, d); err != nil {
		verify.Fail("Aborted")
		return
	}
	verify.Done("Integrity OK")

	// Step 3: Extract
	extract := rotato.New(
		rotato.WithSpinnerStyle("growvert"),
		rotato.WithSpinnerColor(greenBold),
		rotato.WithFrequency(140*time.Millisecond),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Extracting files..."),
		rotato.WithDoneSymbolColor(greenBold),
	)
	extract.Start(ctx)
	if err := pause(ctx, d); err != nil {
		extract.Fail("Aborted")
		return
	}
	extract.Done("Ready")
}

type Item struct {
	text   string
	pause  time.Duration
	mesg   string
	symbol string
	color  rotato.Color
}

func (t *Item) String() string {
	return fmt.Sprintf("%s %s %s", t.color.Sprint(t.symbol), t.text, t.mesg)
}

func sceneLog(ctx context.Context, d time.Duration) {
	sp := rotato.New(
		rotato.WithPrefix("Processing items"),
		rotato.WithMessage("starting"),
		rotato.WithMessageColor(rotato.FgBrightBlue.With(rotato.StyleItalic)),
		rotato.WithPrefixColor(rotato.StyleDim),
		rotato.WithSpinnerColor(rotato.FgBrightYellow.With(rotato.StyleBold)),
		rotato.WithDoneSymbolColor(rotato.FgBrightGreen.With(rotato.StyleBold)),
		rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
		rotato.WithFailMessageColor(rotato.FgBrightRed),
	)

	sp.Start(ctx)

	items := []Item{
		{"item #1", d, "processed successfully", "✓", rotato.FgBrightGreen.With(rotato.StyleBold)},
		{"item #2", d, "requires manual review", "✗", rotato.FgBrightRed.With(rotato.StyleBold)},
		{"item #3", d, "processed successfully", "✓", rotato.FgBrightGreen.With(rotato.StyleBold)},
		{"item #4", d, "requires manual review", "!", rotato.FgBrightYellow.With(rotato.StyleBold)},
	}

	var (
		current = 1
		total   = len(items)
	)

	sp.AddPrefixDecorator(func(prefix string) string {
		progress := fmt.Sprintf("[%d/%d] ", current, total)
		return rotato.FgBrightCyan.With(rotato.StyleBold).Sprint(progress) + prefix
	})

	for _, item := range items {
		sp.UpdateMesg("processing " + item.text)

		if err := pause(ctx, item.pause); err != nil {
			sp.Fail("Aborted")
			return
		}

		sp.Print(item.String())

		current++
	}

	sp.Done("completed")
}

func padRight(s string, width int) string {
	n := width - utf8.RuneCountInString(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(" ", n)
}

func showSymbols(ctx context.Context, d time.Duration) {
	maxLen := 0
	for _, s := range rotato.Spinners() {
		if l := utf8.RuneCountInString(s.Name); l > maxLen {
			maxLen = l
		}
	}

	for _, sp := range rotato.Spinners() {
		r := rotato.New(
			rotato.WithPrefix(padRight(sp.Name, maxLen)),
			rotato.WithSpinnerStyle(sp.Name),
			rotato.WithMessage(hint),
			rotato.WithDoneSymbolColor(greenBold),
			rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
		)
		r.Start(ctx)

		if err := pause(ctx, d); err != nil {
			r.Fail("interrupted")
			return
		}
		symbols := shorten(strings.Join(r.Symbols(), " "))
		r.Done(symbols)
	}
}

func listGroups() {
	fmt.Printf("rotato — %d groups, %d spinners total\n",
		len(rotato.Groups), len(rotato.Spinners()))

	for i, g := range rotato.Groups {
		spinners := rotato.ByGroup(g)
		names := make([]string, len(spinners))
		for j, s := range spinners {
			names[j] = s.Name
		}

		isLast := i == len(rotato.Groups)-1

		prefix := "├──"
		indent := "│   "
		if isLast {
			prefix = "└──"
			indent = "    "
		}

		fmt.Printf("%s %s (%d)\n", prefix, g, len(spinners))
		fmt.Printf("%s%s\n", indent, strings.Join(names, ", "))
	}
}

func showByName(ctx context.Context, name string, d time.Duration) {
	if sp, ok := rotato.ByName(name); ok {
		runRotato(ctx, sp, padRight(sp.Name, 18), d)
	}
}

func showByGroup(ctx context.Context, g rotato.SpinnerGroup, d time.Duration) {
	group := rotato.ByGroup(g)
	if len(group) == 0 {
		fmt.Printf("rotato: group %q not found\n", g)
		return
	}

	maxLen := 0
	for _, s := range group {
		if l := utf8.RuneCountInString(s.Name); l > maxLen {
			maxLen = l
		}
	}

	maxLen = max(maxLen, 14)
	for _, sp := range group {
		if err := ctx.Err(); err != nil {
			return
		}
		runRotato(ctx, sp, padRight(sp.Name, maxLen), d)
	}
}

func runRotato(ctx context.Context, sp rotato.SpinnerStyle, name string, d time.Duration) {
	r := rotato.New(
		rotato.WithSymbols(sp.Frames...),
		rotato.WithPrefix(name),
		rotato.WithMessage(hint),
		rotato.WithDoneSymbolColor(greenBold),
		rotato.WithFailSymbolColor(rotato.StyleBold, rotato.FgBrightRed),
	)
	r.Start(ctx)

	if err := pause(ctx, d); err != nil {
		r.Fail("interrupted")
		return
	}
	symbols := shorten(strings.Join(r.Symbols(), " "))
	r.Done(symbols)
}

func runDemo(ctx context.Context, d time.Duration) {
	scenes := []func(context.Context, time.Duration){
		sceneSimple,    // basic start/done
		sceneMultiStep, // live updates across phases
		sceneError,     // failure path
		sceneElapse,    // shared elapse timer
		sceneLog,       // show log functionality
		sceneCountdown, // graceful shutdown with countdown
	}
	for _, scene := range scenes {
		if ctx.Err() != nil {
			return
		}
		scene(ctx, d)
	}
}

func runMore(ctx context.Context, d time.Duration) {
	scenes := []func(context.Context, time.Duration){
		sceneElapse, // shared elapse timer
	}
	for _, scene := range scenes {
		if ctx.Err() != nil {
			return
		}
		scene(ctx, d)
	}
}

func shorten(s string) string {
	maxLen := 100
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	return string([]rune(s)[:maxLen])
}

func formatDuration(d time.Duration) string {
	switch {
	case d < time.Second:
		return d.Round(100 * time.Millisecond).String()
	case d < time.Minute:
		return d.Round(time.Second).String()
	default:
		return d.Round(time.Second).String()
	}
}

func main() {
	if flags.NonInteractive {
		rotato.SetNonInteractive()
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer stop()

	d := flags.Duration

	switch {
	case flags.All:
		showSymbols(ctx, d)
	case flags.Demo:
		runDemo(ctx, d)
	case flags.Group != "":
		showByGroup(ctx, rotato.SpinnerGroup(flags.Group), d)
	case flags.Show != "":
		showByName(ctx, flags.Show, d)
	case flags.List:
		listGroups()
	case flags.More:
		runMore(ctx, d)
	default:
		flag.Usage()
	}
}

func init() {
	flag.BoolVar(&flags.All, "all", false, "demo all rotatos")
	flag.BoolVar(&flags.Demo, "demo", false, "run narrative demo")
	flag.BoolVar(&flags.List, "list", false, "list spinners groups")
	flag.BoolVar(&flags.NonInteractive, "ni", false, "non-interactive / dumb-terminal mode")
	flag.StringVar(&flags.Group, "group", "", "show spinners by group")
	flag.StringVar(&flags.Show, "show", "", "show spinner by name")
	flag.BoolVar(&flags.More, "more", false, "show more examples")
	flag.DurationVar(&flags.Duration, "duration", 2*time.Second, "animation duration")
	flag.Parse()
}
