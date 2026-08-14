package rotato

import (
	"bytes"
	"context"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSpinnerOutput(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	mesg := "Testing"
	r := New(
		WithWriter(&buf),
		WithMessage(mesg),
		WithDoneMessage("Done"),
		WithFrequency(10*time.Millisecond),
	)

	ctx := context.Background()
	r.Start(ctx)
	time.Sleep(50 * time.Millisecond)
	r.Done()

	output := buf.String()
	if output == "" {
		t.Error("expected spinner output, got empty string")
	}
	if !strings.Contains(output, mesg) {
		t.Errorf("expected spinner output to contain 'Testing', got %q", output)
	}
}

// TestSpinnerState verifies that after Stop() the spinner is no longer
// running.
func TestSpinnerState(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	r := New(
		WithWriter(&buf),
		WithFrequency(10*time.Millisecond),
		WithSymbols([]string{"-", "\\", "|", "/"}...),
		WithDoneMessage("Stopped"),
	)

	ctx := context.Background()
	r.Start(ctx)

	time.Sleep(20 * time.Millisecond)
	// verify that the spinner state is true.
	if !r.isActive {
		t.Error("expected spinner to be running")
	}
	// verify that the spinner state is false.
	r.Done()
	if r.isActive {
		t.Error("expected spinner to be stopped after calling Stop()")
	}
}

// TestSpinnerMessageUpdate verifies that the spinner's message can be updated
// while running.
func TestSpinnerMessageUpdate(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	r := New(
		WithWriter(&buf),
		WithFrequency(10*time.Millisecond),
		WithMessage("Initial"),
		WithDoneMessage("Done"),
	)

	ctx := context.Background()
	r.Start(ctx)

	time.Sleep(20 * time.Millisecond)

	// Update the message.
	r.UpdateMesg("Updated")

	time.Sleep(50 * time.Millisecond)

	r.Done()

	out := buf.String()
	if !strings.Contains(out, "Updated") {
		t.Errorf("expected spinner output to contain updated message, got %q", out)
	}
}

func TestFailMesg(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	r := New(WithWriter(&buf))
	ctx := context.Background()
	r.Start(ctx)
	r.Fail("Failed")
	out := buf.String()
	if !strings.Contains(out, "Failed") {
		t.Errorf("expected spinner output to contain 'Failed', got %q", out)
	}
}

func TestRemoveANSI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Simple ANSI codes",
			input: FgRed.With(StyleBold).Sprint("Hello"),
			want:  "Hello",
		},
		{
			name:  "No ANSI codes",
			input: "No ANSI codes here",
			want:  "No ANSI codes here",
		},
		{
			name:  "Multiple ANSI sequences",
			input: "Text " + FgRed.Sprint("Red") + " and " + FgGreen.Sprint("Green"),
			want:  "Text Red and Green",
		},
		{
			name:  "ANSI only",
			input: FgBlue.With(StyleBold).Sprint("Blue Bold Text"),
			want:  "Blue Bold Text",
		},
		{
			name:  "Empty ANSI",
			input: string(FgBlue),
			want:  "",
		},
	}

	term := Term{}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := term.RemoveANSI(tt.input)
			if got != tt.want {
				t.Errorf("removeANSI(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestContext(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	mesg := "Testing"
	ctx, cancel := context.WithCancel(context.Background())
	r := New(
		WithWriter(&buf),
		WithMessage(mesg),
		WithDoneMessage("Done"),
		WithForceInteractive(),
		WithFrequency(10*time.Millisecond),
	)

	r.Start(ctx)

	time.Sleep(50 * time.Millisecond)

	cancel()

	time.Sleep(50 * time.Millisecond)

	isActive := r.IsRunning()
	if isActive {
		t.Fatalf("spinner should be inactive after context cancellation: %v", isActive)
	}
}

func TestContextCancelThenDone(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())
	r := New(
		WithWriter(&buf),
		WithMessage("Running..."),
		WithDoneMessage("Completed"),
		WithForceInteractive(),
		WithFrequency(10*time.Millisecond),
	)

	r.Start(ctx)

	time.Sleep(50 * time.Millisecond)

	cancel()

	time.Sleep(50 * time.Millisecond)

	if r.IsRunning() {
		t.Fatal("spinner should be inactive after context cancellation")
	}

	output := buf.String()
	if !strings.Contains(output, "Running...") || !strings.Contains(output, r.symbols[0]) {
		t.Errorf("Expected final output to contain completion message after cancel, got:\n%s", output)
	}
}

func TestContextTimeoutStopsSpinner(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	timeout := 10 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	r := New(
		WithWriter(&buf),
		WithMessage("Timing out..."),
		WithForceInteractive(),
		WithFrequency(5*time.Millisecond),
	)

	r.Start(ctx)
	time.Sleep(timeout + 50*time.Millisecond)

	isActive := r.IsRunning()

	if isActive {
		t.Fatalf("spinner should be inactive after context timeout: %v", isActive)
	}
}

func TestStartWithPreCancelledContext(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	mesg := "Test Message"
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	r := New(
		WithWriter(&buf),
		WithMessage(mesg),
		WithFrequency(5*time.Millisecond),
		WithForceInteractive(),
	)

	r.Start(ctx)
	time.Sleep(5 * time.Millisecond)

	isActive := r.IsRunning()

	if isActive {
		t.Fatalf("spinner should not become active with pre-cancelled context")
	}

	output := buf.String()
	if output != "" {
		t.Errorf("expected message to be empty")
	}
}

func TestDisplayMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		symbol    string
		delimiter string
		color     Color
		msg       []string
		prefix    string
		expected  string
	}{
		{
			name:     "symbol + message",
			symbol:   "✓",
			color:    "",
			msg:      []string{"done"},
			expected: "✓ done\n",
		},
		{
			name:     "color applied - ANSI stripped on redirect",
			symbol:   "✓",
			color:    "\x1b[32m",
			msg:      []string{"ok"},
			expected: "✓ ok\n",
		},
		{
			name:      "no symbol",
			delimiter: "--",
			color:     "",
			msg:       []string{"hello"},
			expected:  "hello\n",
		},
		{
			name:      "prefix overrides symbol position",
			symbol:    "✓",
			color:     "",
			delimiter: NBSP,
			msg:       []string{"done"},
			prefix:    "step1",
			expected:  "step1" + NBSP + "✓ done\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}

			r := New(
				WithWriter(buf),
				WithDelimiter(tt.delimiter),
				WithDoneSymbol(tt.symbol),
				WithPrefix(tt.prefix),
				WithForceInteractive(), // bypass isInteractive early-return
			)

			r.displayMessage(tt.symbol, tt.color, tt.msg...)

			got := buf.String()
			if got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestDecorate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		input      string
		decorators []MessageDecorator
		want       string
	}{
		{
			name:  "single_decorator",
			input: "hello",
			decorators: []MessageDecorator{
				func(s string) string { return s + "!" },
			},
			want: "hello!",
		},
		{
			name:  "multiple_decorators_order",
			input: "go",
			decorators: []MessageDecorator{
				func(s string) string { return s + "lang" },
				func(s string) string { return "[" + s + "]" },
			},
			want: "[golang]",
		},
		{
			name:       "no_decorators",
			input:      "plain",
			decorators: []MessageDecorator{},
			want:       "plain",
		},
		{
			name:       "nil_decorators_slice",
			input:      "nilcase",
			decorators: nil,
			want:       "nilcase",
		},
		{
			name:  "empty_string_with_decorators",
			input: "",
			decorators: []MessageDecorator{
				func(s string) string { return s + "prefix" },
				func(s string) string { return s + "suffix" },
			},
			want: "prefixsuffix",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := decorate(tt.input, tt.decorators)
			if got != tt.want {
				t.Fatalf("decorate(%q, decorators) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestContextCancelled(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	r := New(
		WithWriter(&buf),
		WithForceInteractive(),
		WithContextDoneHandler(func(r *Rotato, err error) {
			r.Fail(err.Error())
		}),
	)
	r.Start(ctx)

	time.Sleep(20 * time.Millisecond)

	want := context.Canceled.Error()
	r.writerMu.Lock()
	got := buf.String()
	r.writerMu.Unlock()

	if !strings.Contains(got, want) {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}

func TestContextDeadlineExceeded(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	r := New(
		WithWriter(&buf),
		WithPrefix("CtxDeadlineExceeded"),
		WithForceInteractive(),
		WithContextDoneHandler(func(r *Rotato, err error) {
			r.Fail(err.Error())
		}),
	)
	r.Start(ctx)

	time.Sleep(20 * time.Millisecond)

	want := context.DeadlineExceeded.Error()
	r.writerMu.Lock()
	got := buf.String()
	r.writerMu.Unlock()

	if !strings.Contains(got, want) {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

// TestStopSpinner_NoDeadlockDone is the same scenario but exercises Done
// instead of Fail.
func TestStopSpinner_NoDeadlockDone(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	r := New(
		WithWriter(io.Discard),
		WithMessage("Loading..."),
		WithForceInteractive(),
	)
	r.Start(ctx)

	time.Sleep(50 * time.Millisecond)

	cancel()

	time.Sleep(50 * time.Millisecond)

	done := make(chan struct{})
	go func() {
		defer close(done)
		r.Done("Finished!")
	}()

	select {
	case <-done:
		// good.
	case <-time.After(2 * time.Second):
		t.Fatal("Done() blocked: stopSpinner deadlock detected (goroutine already exited)")
	}
}

// TestRotato_DoneConcurrency tests the data race by firing multiple Done()
// calls simultaneously.
func TestRotato_DoneConcurrency(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	r := New(
		WithWriter(&buf),
		WithForceInteractive(),
	)
	r.isActive = true

	const workers = 50
	var wg sync.WaitGroup
	wg.Add(workers)

	// fire 50 goroutines trying to stop the spinner at the exact same time
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			r.Done("Concurrent Done")
		}()
	}

	wg.Wait()

	// verify the spinner state was safely transitioned
	if r.isActive {
		t.Errorf("expected isActive to be false")
	}

	// the doneChan should only receive `exactly` one signal, regardless of
	// how many goroutines called `Done()`
	if len(r.doneChan) != 1 {
		t.Errorf("expected doneChan to have exactly 1 item, got %d", len(r.doneChan))
	}
}

// TestStopSpinner_DoneFailRace simulates the logical race condition where a context
// timeout triggers Fail() at the exact same millisecond the deferred Done() fires.
func TestStopSpinner_DoneFailRace(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	r := New(
		WithWriter(&buf),
		WithForceInteractive(),
	)
	r.isActive = true

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		r.Fail("Context Timeout")
	}()

	go func() {
		defer wg.Done()
		r.Done("Graceful Exit")
	}()

	wg.Wait()

	if r.isActive {
		t.Errorf("expected isActive to be false")
	}

	if len(r.doneChan) != 1 {
		t.Errorf("expected doneChan to receive exactly 1 stop signal, got %d", len(r.doneChan))
	}
}

func TestWriteAbove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		msg         string
		isRunning   bool
		prefix      string
		message     string
		setupFn     func(*Rotato)
		checkOutput func(t *testing.T, r *Rotato, output string)
	}{
		{
			name:      "spinner_not_running",
			msg:       "test message",
			isRunning: false,
			prefix:    "",
			message:   "Loading",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "test message") {
					t.Fatalf("expected output to contain 'test message', got %q", output)
				}
				if !strings.Contains(output, "\n") {
					t.Fatalf("expected newline in output, got %q", output)
				}
			},
		},
		{
			name:      "spinner_running_without_prefix",
			msg:       "log line",
			isRunning: true,
			prefix:    "",
			message:   "Processing",
			setupFn: func(r *Rotato) {
				r.frame = "⠋"
				r.isActive = true
			},
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "log line") {
					t.Fatalf("expected output to contain 'log line', got %q", output)
				}
				if !strings.Contains(output, r.term.ClearLine()) {
					t.Fatalf("expected clearChars in output, got %q", output)
				}
			},
		},
		{
			name:      "spinner_running_with_prefix",
			msg:       "status update",
			isRunning: true,
			prefix:    "[1/10]",
			message:   "Syncing",
			setupFn: func(r *Rotato) {
				r.frame = "⠙"
				r.isActive = true
				r.prefixMesg = "[1/10]"
			},
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "status update") {
					t.Fatalf("expected output to contain 'status update', got %q", output)
				}
				if !strings.Contains(output, "[1/10]") {
					t.Fatalf("expected prefix '[1/10]' in output, got %q", output)
				}
			},
		},
		{
			name:      "empty_message",
			msg:       "",
			isRunning: false,
			prefix:    "",
			message:   "Loading",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				// Empty message still results in newline
				if output == "" {
					t.Fatalf("expected output to contain newline for empty message")
				}
			},
		},
		{
			name:      "message_with_special_chars",
			msg:       "Error: file not found (exit code: 1)",
			isRunning: false,
			prefix:    "",
			message:   "Loading",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "Error: file not found (exit code: 1)") {
					t.Fatalf("expected output to contain special chars message, got %q", output)
				}
			},
		},
		{
			name:      "multiline_message_only_first_line_written",
			msg:       "line 1\nline 2",
			isRunning: false,
			prefix:    "",
			message:   "Loading",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "line 1\nline 2") {
					t.Fatalf("expected output to contain multiline message, got %q", output)
				}
			},
		},
		{
			name:      "spinner_running_no_frame_set",
			msg:       "test",
			isRunning: true,
			prefix:    "",
			message:   "",
			setupFn: func(r *Rotato) {
				r.isActive = true
				r.frame = ""
			},
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "test") {
					t.Fatalf("expected output to contain 'test', got %q", output)
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			r := New(
				WithWriter(&buf),
				WithMessage(tt.message),
				WithForceInteractive(),
			)

			if tt.setupFn != nil {
				tt.setupFn(r)
			}

			r.writeAbove(tt.msg)

			output := buf.String()
			tt.checkOutput(t, r, output)
		})
	}
}

func TestRenderCurrentLocked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		frame       string
		message     string
		prefix      string
		decorator   MessageDecorator
		checkOutput func(t *testing.T, r *Rotato, output string)
	}{
		{
			name:    "simple_frame_and_message",
			frame:   "⠋",
			message: "Loading",
			prefix:  "",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "⠋") {
					t.Fatalf("expected output to contain frame '⠋', got %q", output)
				}
				if !strings.Contains(output, "Loading") {
					t.Fatalf("expected output to contain 'Loading', got %q", output)
				}
			},
		},
		{
			name:    "with_prefix",
			frame:   "⠙",
			message: "Syncing",
			prefix:  "[2/5]",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "[2/5]") {
					t.Fatalf("expected output to contain prefix '[2/5]', got %q", output)
				}
				if !strings.Contains(output, "⠙") {
					t.Fatalf("expected output to contain frame, got %q", output)
				}
				if !strings.Contains(output, "Syncing") {
					t.Fatalf("expected output to contain message, got %q", output)
				}
			},
		},
		{
			name:    "empty_message",
			frame:   "⠹",
			message: "",
			prefix:  "",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "⠹") {
					t.Fatalf("expected output to contain frame, got %q", output)
				}
				if !strings.Contains(output, r.term.ClearLine()) {
					t.Fatalf("expected clearChars in output, got %q", output)
				}
			},
		},
		{
			name:    "empty_prefix_with_message",
			frame:   "⠸",
			message: "Processing",
			prefix:  "",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "⠸") {
					t.Fatalf("expected frame in output, got %q", output)
				}
				if !strings.Contains(output, "Processing") {
					t.Fatalf("expected message in output, got %q", output)
				}
			},
		},
		{
			name:      "with_message_decorator",
			frame:     "⠼",
			message:   "decorated",
			prefix:    "",
			decorator: func(s string) string { return "[" + s + "]" },
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "[decorated]") {
					t.Fatalf("expected decorated message '[decorated]' in output, got %q", output)
				}
			},
		},
		{
			name:    "prefix_and_message",
			frame:   "⠾",
			message: "Done",
			prefix:  "[final]",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "[final]") {
					t.Fatalf("expected prefix in output, got %q", output)
				}
				if !strings.Contains(output, "⠾") {
					t.Fatalf("expected frame in output, got %q", output)
				}
				if !strings.Contains(output, "Done") {
					t.Fatalf("expected message in output, got %q", output)
				}
			},
		},
		{
			name:    "empty_frame_empty_message",
			frame:   "",
			message: "",
			prefix:  "",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, r.term.ClearLine()) {
					t.Fatalf("expected clearChars in output, got %q", output)
				}
			},
		},
		{
			name:    "frame_only_no_message",
			frame:   "⠷",
			message: "",
			prefix:  "",
			checkOutput: func(t *testing.T, r *Rotato, output string) {
				t.Helper()
				if !strings.Contains(output, "⠷") {
					t.Fatalf("expected frame in output, got %q", output)
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			r := New(
				WithWriter(&buf),
				WithMessage(tt.message),
				WithForceInteractive(),
				withTerm(Term{}),
			)

			r.frame = tt.frame
			if tt.prefix != "" {
				r.prefixMesg = tt.prefix
			}
			if tt.decorator != nil {
				r.AddMessageDecorator(tt.decorator)
			}

			r.writerMu.Lock()
			r.renderCurrentLocked()
			r.writerMu.Unlock()

			output := buf.String()
			tt.checkOutput(t, r, output)
		})
	}
}

func TestNew_WithSymbols(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer

	tests := []struct {
		name           string
		opts           []Option
		spinnerName    SpinnerName
		wantDoneSymbol string
		wantFailSymbol string
		wantMessage    string
		wantDelimiter  string
		wantActive     bool
		wantWriter     io.Writer // nil = expect default os.Stdout
	}{
		{
			name:           "no_options_defaults",
			opts:           nil,
			spinnerName:    "",
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "with_symbols_default_style",
			opts:           []Option{WithSymbolsDefault()},
			spinnerName:    SpinnerDefault,
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "with_symbols_block_style",
			opts:           []Option{WithSymbolsBlock()},
			spinnerName:    SpinnerBlock,
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "with_symbols_barblock4_style",
			opts:           []Option{WithSymbolsBarBlock4()},
			spinnerName:    SpinnerBlockbar4,
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "multiple_options_last_style_wins",
			opts:           []Option{WithSymbolsBlock(), WithSymbolsBarBlock2()},
			spinnerName:    SpinnerBlockbar2,
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "unknown_spinner_style_keeps_defaults",
			opts:           []Option{WithSpinnerStyle(SpinnerName("nonexistent-style"))},
			spinnerName:    "",
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     nil,
		},
		{
			name:           "with_custom_writer",
			opts:           []Option{WithWriter(&buf)},
			spinnerName:    "",
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     &buf,
		},
		{
			name:           "with_nil_writer_overrides_default",
			opts:           []Option{WithWriter(nil)},
			spinnerName:    "",
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     io.Writer(nil),
		},
		{
			name:           "writer_option_combined_with_style",
			opts:           []Option{WithSymbolsBarBlock3(), WithWriter(&buf)},
			spinnerName:    SpinnerBlockbar3,
			wantDoneSymbol: "✓",
			wantFailSymbol: "✗",
			wantMessage:    "Loading...",
			wantDelimiter:  NBSP,
			wantActive:     false,
			wantWriter:     &buf,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := New(tt.opts...)

			if r == nil {
				t.Fatal("New() returned nil")
			}

			if r.isActive != tt.wantActive {
				t.Errorf("isActive = %v; want %v", r.isActive, tt.wantActive)
			}
			if r.message != tt.wantMessage {
				t.Errorf("message = %q; want %q", r.message, tt.wantMessage)
			}
			if r.delimiter != tt.wantDelimiter {
				t.Errorf("delimiter = %q; want %q", r.delimiter, tt.wantDelimiter)
			}
			if r.doneSymbol != tt.wantDoneSymbol {
				t.Errorf("doneSymbol = %q; want %q", r.doneSymbol, tt.wantDoneSymbol)
			}
			if r.failSymbol != tt.wantFailSymbol {
				t.Errorf("failSymbol = %q; want %q", r.failSymbol, tt.wantFailSymbol)
			}
			if r.doneChan == nil {
				t.Fatal("doneChan is nil; want initialized buffered channel")
			}
			if cap(r.doneChan) != 1 {
				t.Errorf("cap(doneChan) = %d; want 1", cap(r.doneChan))
			}

			switch {
			case tt.name == "with_nil_writer_overrides_default":
				if r.writer != nil {
					t.Errorf("writer = %v; want nil (explicitly set)", r.writer)
				}
			case tt.wantWriter != nil:
				if r.writer != tt.wantWriter {
					t.Errorf("writer = %v; want %v", r.writer, tt.wantWriter)
				}
			default:
				if r.writer != io.Writer(os.Stdout) {
					t.Errorf("writer = %v; want os.Stdout (default)", r.writer)
				}
			}

			wantFreq := 100 * time.Millisecond
			wantSymbols := defaultSymbols
			if tt.spinnerName != "" {
				if sp, ok := ByName(tt.spinnerName); ok {
					wantFreq = sp.Frequency
					wantSymbols = sp.Frames
				} else {
					t.Fatalf("test setup error: ByName(%q) not found", tt.spinnerName)
				}
			}

			if r.frequency != wantFreq {
				t.Errorf("frequency = %v; want %v", r.frequency, wantFreq)
			}
			if !reflect.DeepEqual(r.symbols, wantSymbols) {
				t.Errorf("symbols = %v; want %v", r.symbols, wantSymbols)
			}
		})
	}
}

func TestBuildPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		opts   []Option
		symbol string
		want   string
	}{
		{
			name: "normal_prefix_with_symbol",
			opts: []Option{
				WithPrefix("loading"),
				WithDelimiter(":"),
			},
			symbol: "*",
			want:   "loading:* ",
		},
		{
			name:   "empty_prefix_and_symbol",
			opts:   nil,
			symbol: "",
			want:   NBSP,
		},
		{
			name: "empty_symbol_no_trailing_space",
			opts: []Option{
				WithPrefix("wait"),
				WithDelimiter("-"),
			},
			symbol: "",
			want:   "wait-",
		},
		{
			name: "prefix_color_wraps_prefix",
			opts: []Option{
				WithPrefix("go"),
				WithPrefixColor(FgGreen),
				WithDelimiter(":"),
			},
			symbol: "",
			want:   FgGreen.Sprint("go") + ":",
		},
		{
			name: "delimiter_color_wraps_delimiter",
			opts: []Option{
				WithPrefix("go"),
				WithDelimiter(":"),
				WithDelimiterColor(FgRed),
			},
			symbol: "",
			want:   "go" + FgRed.Sprint(":"),
		},
		{
			name:   "default_delimiter_nbsp",
			opts:   []Option{WithPrefix("x")},
			symbol: "",
			want:   "x" + NBSP,
		},
		{
			name: "multi_char_symbol",
			opts: []Option{
				WithPrefix("p"),
				WithDelimiter(">"),
			},
			symbol: "==>",
			want:   "p>==> ",
		},
		{
			name: "prefix_and_delimiter_colors_combined",
			opts: []Option{
				WithPrefix("run"),
				WithPrefixColor(FgBlue),
				WithDelimiter(":"),
				WithDelimiterColor(FgYellow),
			},
			symbol: "!",
			want:   FgBlue.Sprint("run") + FgYellow.Sprint(":") + "! ",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := New(tt.opts...)

			var sb strings.Builder
			r.buildPrefix(&sb, tt.symbol)

			if got := sb.String(); got != tt.want {
				t.Fatalf("buildPrefix() = %q; want %q", got, tt.want)
			}
		})
	}
}
