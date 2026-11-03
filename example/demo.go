// rotato demo
//
//	narrative demo
//	go run github.com/mateconpizza/rotato/example@latest -demo
//
//	all spinners
//	go run github.com/mateconpizza/rotato/example@latest -all
//
//	list by groups
//	go run github.com/mateconpizza/rotato/example@latest -list
//
//	show by group
//	go run github.com/mateconpizza/rotato/example@latest -group <name>
//
//	show by name
//	go run github.com/mateconpizza/rotato/example@latest -show <name>
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
	"text/tabwriter"
	"time"
	"unicode/utf8"

	"github.com/mateconpizza/rotato"
)

type Flags struct {
	All            bool
	NonInteractive bool
	Demo           bool
	List           bool
	Group          string
	Show           string
	Duration       time.Duration
}

//nolint:gosec // deterministic demo seed
var rng = rand.New(rand.NewSource(42))

// Color in a sequences for colors and text styles.
type Color string

func NewColor(codes ...Color) Color { return Color(combine(codes...)) }

const (
	ColorReset Color = "\x1b[0m" // Reset all attributes

	FgBlue Color = "\x1b[34m"

	FgBrightRed    Color = "\x1b[91m"
	FgBrightGreen  Color = "\x1b[92m"
	FgBrightYellow Color = "\x1b[93m"
	FgBrightBlue   Color = "\x1b[94m"
	FgBrightCyan   Color = "\x1b[96m"

	FgGray   Color = "\x1b[90m"       // Alias for bright black
	FgOrange Color = "\x1b[38;5;214m" // 256-color orange

	StyleBold   Color = "\x1b[1m" // Bold or increased intensity
	StyleDim    Color = "\x1b[2m" // Faint or dim
	StyleItalic Color = "\x1b[3m" // Italic
	StyleBlink  Color = "\x1b[5m" // Slow blink
)

func (c Color) Wrap(text string, styles ...Color) string {
	return string(c) + combine(styles...) + text + ColorReset.String()
}
func (c Color) With(styles ...Color) Color        { return Color(string(c) + combine(styles...)) }
func (c Color) Sprint(a ...any) string            { return c.Wrap(fmt.Sprint(a...)) }
func (c Color) Sprintf(f string, a ...any) string { return c.Wrap(fmt.Sprintf(f, a...)) }
func (c Color) ColorFn(s string) string           { return c.Sprint(s) }
func (c Color) String() string                    { return string(c) }

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

// Colors.
var (
	greenBold   = NewColor(StyleBold, FgBrightGreen)
	greenItalic = NewColor(FgBrightGreen, StyleItalic)
	hint        = "(ctrl+c to exit · %s)"
)

type Demo struct {
	Flags *Flags
}

func NewDemo() *Demo {
	return &Demo{
		Flags: &Flags{},
	}
}

// Run runs a simple demo (the one in the README).
func (d *Demo) Run(ctx context.Context) {
	if d.Flags.NonInteractive {
		hint = "(ctrl+c to exit)"
	}

	scenes := []func(context.Context){
		d.sceneSimple,    // basic start/done
		d.sceneMultiStep, // live updates across phases
		d.sceneError,     // failure path
		d.sceneElapse,    // shared elapse timer
		d.sceneLog,       // show log functionality
		d.sceneCountdown, // graceful shutdown with countdown
	}
	for _, scene := range scenes {
		if ctx.Err() != nil {
			return
		}
		scene(ctx)
	}
}

// All shows all registered symbols|spinners.
func (d *Demo) All(ctx context.Context) {
	maxLen := 0
	for _, s := range rotato.Spinners() {
		if l := utf8.RuneCountInString(string(s.Name)); l > maxLen {
			maxLen = l
		}
	}

	for _, sp := range rotato.Spinners() {
		if sp.Name == rotato.SpinnerDefault {
			continue
		}

		deadline := time.Now().Add(d.Flags.Duration)

		r := rotato.New(d.opts(
			rotato.WithPrefix(padRight(sp.Name, maxLen)),
			rotato.WithSpinnerStyle(sp.Name),
			rotato.WithMessage(hint+" freq: "+sp.Frequency.String()),
			rotato.WithMessageDecorator(decorator(deadline)),
		)...)
		r.Start(ctx)

		if err := pause(ctx, d.Flags.Duration); err != nil {
			r.Fail("interrupted")
			return
		}
		symbols := shorten(strings.Join(r.Symbols(), " "))
		r.Done(symbols)
	}
}

// ByGroup runs spinners by the given group.
func (d *Demo) ByGroup(ctx context.Context, g rotato.SpinnerGroup) {
	group := rotato.ByGroup(g)
	if len(group) == 0 {
		fmt.Printf("rotato: group %q not found\n", g)
		return
	}

	maxLen := 0
	for _, s := range group {
		if l := utf8.RuneCountInString(string(s.Name)); l > maxLen {
			maxLen = l
		}
	}

	maxLen = max(maxLen, 14)
	for _, sp := range group {
		if err := ctx.Err(); err != nil {
			return
		}
		sp := sp
		d.runRotato(ctx, &sp, padRight(sp.Name, maxLen))
	}
}

// ByName runs the spinner with the given name.
func (d *Demo) ByName(ctx context.Context, s string) {
	if err := ctx.Err(); err != nil {
		return
	}

	name := rotato.SpinnerName(s)
	sp, ok := rotato.ByName(name)
	if !ok {
		fmt.Printf("rotato: spinner %q not found\n", name)
		return
	}
	d.runRotato(ctx, &sp, padRight(sp.Name, 18))
}

// List pretty format spinner's groups by name.
func (d *Demo) List() {
	var (
		dim    = StyleDim.Sprint
		bold   = StyleBold.Sprint
		header = FgBrightYellow.Wrap("rotato", StyleBold)
	)

	fmt.Printf(
		"%s · %s groups, %s spinners total\n",
		header,
		bold(len(rotato.Groups)),
		bold(len(rotato.Spinners())),
	)

	for i, g := range rotato.Groups {
		spinners := rotato.ByGroup(g)
		names := make([]string, len(spinners))
		for j, s := range spinners {
			names[j] = string(s.Name)
		}

		isLast := i == len(rotato.Groups)-1

		prefix := dim("├─")
		indent := dim("│  ")
		if isLast {
			prefix = dim("└─")
			indent = "   "
		}

		fmt.Printf("%s %s (%d)\n", prefix, bold(g), len(spinners))
		fmt.Printf("%s%s\n", indent, strings.Join(names, ", "))
	}
}

// Usage prints CLI usage.
func (d *Demo) Usage() {
	const pkgPath = "github.com/mateconpizza/rotato/example@latest"
	w := flag.CommandLine.Output()

	fmt.Fprintf(w, "rotato - terminal spinner gallery & demonstration tool\n\n")
	fmt.Fprintf(w, "USAGE:\n")
	fmt.Fprintf(w, "  go run %s [flags]\n\n", pkgPath)

	fmt.Fprintf(w, "FLAGS:\n")

	// tabwriter automatically aligns columns with clean spacing
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "  -a, --all\tdemo all rotatos")
	fmt.Fprintln(tw, "  -d, --demo\trun narrative demo")
	fmt.Fprintln(tw, "  -g, --group <name>\tshow spinners by group")
	fmt.Fprintln(tw, "  -s, --show <name>\tshow spinner by name")
	fmt.Fprintln(tw, "  -l, --list\tlist spinner groups")
	fmt.Fprintln(tw, "  -t, --duration <dur>\tanimation duration (default 2s)")
	fmt.Fprintln(tw, "  -ni\tnon-interactive / dumb-terminal mode")
	_ = tw.Flush()

	fmt.Fprintf(w, "\nEXAMPLES:\n")
	fmt.Fprintf(w, "  $ go run %s --demo\n", pkgPath)
	fmt.Fprintf(w, "  $ go run %s --group dots --duration 5s\n", pkgPath)
	fmt.Fprintf(w, "  $ go run %s --show dots6\n", pkgPath)
}

// sceneSimple shows the most basic spinner usage: start, wait, done.
func (d *Demo) sceneSimple(ctx context.Context) {
	r := rotato.New(d.opts(
		rotato.WithPrefix("Fetching config"),
		rotato.WithPrefixColor(StyleItalic.ColorFn),
		rotato.WithSpinnerColor(FgBrightCyan.ColorFn),
		rotato.WithDoneMessageColor(greenItalic.ColorFn),
	)...)
	r.Start(ctx)

	if err := pause(ctx, d.Flags.Duration); err != nil {
		r.Fail("Aborted")
		return
	}
	r.Done("config.yaml loaded")
}

// sceneMultiStep shows how a single spinner can represent a multi-phase task
// by updating its message and symbol set as work progresses.
func (d *Demo) sceneMultiStep(ctx context.Context) {
	r := rotato.New(d.opts(
		rotato.WithSymbolsCircles6(),
		rotato.WithSpinnerColor(FgOrange.ColorFn),
		rotato.WithPrefix("S3 backup"),
		rotato.WithPrefixColor(StyleBold.ColorFn),
		rotato.WithMessage("connecting..."),
		rotato.WithDoneMessageColor(StyleBold.ColorFn),
	)...)
	r.Start(ctx)

	// Phase 1 - connect
	if err := pause(ctx, d.Flags.Duration); err != nil {
		r.Fail("connection aborted")
		return
	}

	// Phase 2 - confirm connection with a static check mark
	r.UpdateSymbols(rotato.WithSymbols(greenBold.Sprint("✓")))
	r.UpdateMesgColor(greenBold.ColorFn)
	r.UpdateMesg("connected")
	if err := pause(ctx, 800*time.Millisecond); err != nil {
		r.Fail("aborted after connect")
		return
	}

	// Phase 3 - transfer files
	r.UpdateSymbols(rotato.WithSymbolsBarBlock())
	r.UpdateMesgColor(FgGray.ColorFn)
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
func (d *Demo) sceneError(ctx context.Context) {
	r := rotato.New(d.opts(
		rotato.WithPrefix("AWS Health Check"),
		rotato.WithSymbolsPipe(),
		rotato.WithSpinnerColor(FgBlue.With(StyleBold).ColorFn),
		rotato.WithMessage("pinging us-east-1..."),
	)...)
	r.Start(ctx)

	if err := pause(ctx, d.Flags.Duration); err != nil {
		r.Fail("aborted by user")
		return
	}

	// Simulate an endpoint that is down.
	r.Fail("connection refused (us-east-1:443)")
}

// sceneElapse runs sequential spinner tasks (download, verify, extract) and
// shows elapsed time via a shared prefix decorator with a timeout.
func (d *Demo) sceneElapse(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 12*time.Second)
	defer cancel()

	task := "Work"
	t := time.Now()
	formatter := func(mesg string) string {
		d := time.Since(t).Round(time.Second)
		return mesg + " " + rotato.DimElapsedDecorator(d, StyleDim.ColorFn)
	}

	// Step 1: Download
	download := rotato.New(d.opts(
		rotato.WithSpinnerColor(FgBrightGreen.ColorFn),
		rotato.WithDoneSymbol("\u203A"),
		rotato.WithDoneSymbolColor(FgBrightCyan.ColorFn),
		rotato.WithSymbolsBarBlock7(),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Downloading chunks..."),
	)...)
	download.Start(ctx)
	if err := pause(ctx, d.Flags.Duration); err != nil {
		download.Fail("Aborted")
		return
	}
	download.Done("Download complete")

	// Step 2: Verify
	verify := rotato.New(d.opts(
		rotato.WithDoneSymbol("\u203A"),
		rotato.WithDoneSymbolColor(FgBrightYellow.ColorFn),
		rotato.WithSymbolsCircles7(),
		rotato.WithSpinnerColor(FgBrightCyan.ColorFn),
		rotato.WithFrequency(110*time.Millisecond),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Verifying checksum..."),
	)...)

	verify.Start(ctx)

	if err := pause(ctx, d.Flags.Duration); err != nil {
		verify.Fail("Aborted")
		return
	}
	verify.Done("Integrity OK")

	// Step 3: Extract
	extract := rotato.New(d.opts(
		rotato.WithSpinnerStyle("growvert"),
		rotato.WithSpinnerColor(greenBold.ColorFn),
		rotato.WithFrequency(140*time.Millisecond),
		rotato.WithPrefix(task),
		rotato.WithPrefixDecorator(formatter), // elapsed
		rotato.WithMessage("Extracting files..."),
	)...)
	extract.Start(ctx)
	if err := pause(ctx, d.Flags.Duration); err != nil {
		extract.Fail("Aborted")
		return
	}
	extract.Done("Ready")
}

func (d *Demo) sceneLog(ctx context.Context) {
	sp := rotato.New(d.opts(
		rotato.WithPrefix("Processing items"),
		rotato.WithMessage("starting"),
		rotato.WithMessageColor(FgBrightBlue.With(StyleItalic).ColorFn),
		rotato.WithPrefixColor(StyleDim.ColorFn),
		rotato.WithSpinnerColor(FgBrightYellow.With(StyleBold).ColorFn),
	)...)

	sp.Start(ctx)

	items := []Item{
		{"item #1", d.Flags.Duration, "processed successfully", "✓", FgBrightGreen.With(StyleBold)},
		{"item #2", d.Flags.Duration, "requires manual review", "✗", FgBrightRed.With(StyleBold)},
		{"item #3", d.Flags.Duration, "processed successfully", "✓", FgBrightGreen.With(StyleBold)},
		{"item #4", d.Flags.Duration, "requires manual review", "!", FgBrightYellow.With(StyleBold)},
	}

	var (
		current = 1
		total   = len(items)
	)

	sp.AddPrefixDecorator(func(prefix string) string {
		progress := fmt.Sprintf("[%d/%d] ", current, total)
		return FgBrightCyan.With(StyleBold).Sprint(progress) + prefix
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

// sceneCountdown demonstrates a countdown: a graceful shutdown that shows
// remaining time using the library's own formatting helpers.
func (d *Demo) sceneCountdown(ctx context.Context) {
	var color Color
	ctx, cancel := context.WithTimeout(ctx, d.Flags.Duration)
	defer cancel()

	r := rotato.New(d.opts(
		rotato.WithPrefix("Shutting Down"),
		rotato.WithPrefixColor(StyleBold.ColorFn),
		rotato.WithMessage("in"),
		rotato.WithSymbolsBlockPretty(),
		rotato.WithDoneMessageColor(StyleBlink.With(greenBold).ColorFn),
		rotato.WithDoneSymbolColor(StyleBlink.With(greenBold).ColorFn),
		rotato.WithMessageDecorator(func(mesg string) string {
			deadline, ok := ctx.Deadline()
			if !ok {
				return mesg
			}
			remaining := max(time.Until(deadline), 0)
			switch {
			case remaining > 3*time.Second:
				color = FgBrightGreen
			case remaining > time.Second:
				color = FgOrange
			default:
				color = FgBrightRed
			}

			return color.With(StyleItalic).Sprintf(
				"in %s...",
				formatDuration(remaining),
			)
		}),
	)...)
	r.Start(ctx)

	// Wait for the internal deadline; the spinner stops itself when ctx fires.
	<-ctx.Done()
	r.UpdatePrefix("Demo")
	r.UpdatePrefixColor(StyleBold.ColorFn)
	r.UpdateDoneSymbol("\u25A0")
	r.Done("goodbye!")
}

func (d *Demo) opts(extra ...rotato.Option) []rotato.Option {
	opts := make([]rotato.Option, 0, len(extra))
	if d.Flags.NonInteractive {
		opts = append(opts, rotato.WithNonInteractive())
	}

	// common
	opts = append(
		opts,
		rotato.WithDoneSymbolColor(greenBold.ColorFn),
		rotato.WithFailMessageColor(FgBrightRed.ColorFn),
		rotato.WithFailSymbolColor(StyleBold.With(FgBrightRed).ColorFn),
	)

	return append(opts, extra...)
}

func (d *Demo) parseFlags() {
	// All (-a, -all)
	flag.BoolVar(&d.Flags.All, "all", false, "")
	flag.BoolVar(&d.Flags.All, "a", false, "")

	// Demo (-d, -demo)
	flag.BoolVar(&d.Flags.Demo, "demo", false, "")
	flag.BoolVar(&d.Flags.Demo, "d", false, "")

	// List (-l, -list)
	flag.BoolVar(&d.Flags.List, "list", false, "")
	flag.BoolVar(&d.Flags.List, "l", false, "")

	// Non-Interactive (-n, -ni)
	flag.BoolVar(&d.Flags.NonInteractive, "ni", false, "")

	// Group (-g, -group)
	flag.StringVar(&d.Flags.Group, "group", "", "")
	flag.StringVar(&d.Flags.Group, "g", "", "")

	// Show (-s, -show)
	flag.StringVar(&d.Flags.Show, "show", "", "")
	flag.StringVar(&d.Flags.Show, "s", "", "")

	// Duration (-t, -duration)
	flag.DurationVar(&d.Flags.Duration, "duration", 2*time.Second, "")
	flag.DurationVar(&d.Flags.Duration, "t", 2*time.Second, "")

	flag.Usage = d.Usage
	flag.Parse()
}

func main() {
	demo := NewDemo()
	demo.parseFlags()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer stop()

	f := demo.Flags

	switch {
	case f.All:
		demo.All(ctx)
	case f.Demo:
		demo.Run(ctx)
	case f.Group != "":
		demo.ByGroup(ctx, rotato.SpinnerGroup(f.Group))
	case f.Show != "":
		demo.ByName(ctx, f.Show)
	case f.List:
		demo.List()
	default:
		demo.Usage()
	}
}

var decorator = func(deadline time.Time) rotato.MessageDecorator {
	return func(mesg string) string {
		remaining := max(time.Until(deadline).Round(time.Second), 0)
		return StyleDim.With(StyleItalic).Sprintf(mesg, remaining)
	}
}

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

type Item struct {
	text   string
	pause  time.Duration
	mesg   string
	symbol string
	color  Color
}

func (t *Item) String() string {
	return fmt.Sprintf("%s %s %s", t.color.Sprint(t.symbol), t.text, t.mesg)
}

func padRight(name rotato.SpinnerName, width int) string {
	s := string(name)
	n := width - utf8.RuneCountInString(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(" ", n)
}

func (d *Demo) runRotato(ctx context.Context, sp *rotato.SpinnerStyle, name string) {
	ctxTime, cancel := context.WithTimeout(ctx, d.Flags.Duration)
	deadline, _ := ctxTime.Deadline()
	defer cancel()

	r := rotato.New(d.opts(
		rotato.WithSpinnerStyle(sp.Name),
		rotato.WithPrefix(name),
		rotato.WithMessage(hint+" freq: "+sp.Frequency.String()),
		rotato.WithMessageDecorator(decorator(deadline)),
	)...)
	r.Start(ctx)

	if err := pause(ctx, d.Flags.Duration); err != nil {
		r.Fail("interrupted")
		return
	}
	symbols := shorten(strings.Join(r.Symbols(), " "))
	r.Done(symbols)
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
