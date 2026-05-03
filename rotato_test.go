package rotato

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"
)

func TestSpinnerOutput(t *testing.T) {
	var buf bytes.Buffer
	mesg := "Testing"
	r := New(
		WithWriter(&buf),
		WithMesg(mesg),
		WithDoneMesg("Done"),
		WithFrequency(10*time.Millisecond),
	)

	r.Start()
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
	var buf bytes.Buffer
	r := New(
		WithWriter(&buf),
		WithFrequency(10*time.Millisecond),
		WithSymbols([]string{"-", "\\", "|", "/"}...),
		WithDoneMesg("Stopped"),
	)
	r.Start()
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
	var buf bytes.Buffer

	r := New(
		WithWriter(&buf),
		WithFrequency(10*time.Millisecond),
		WithMesg("Initial"),
		WithDoneMesg("Done"),
	)
	r.Start()
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
	var buf bytes.Buffer
	r := New(WithWriter(&buf))
	r.Start()
	r.Fail("Failed")
	out := buf.String()
	if !strings.Contains(out, "Failed") {
		t.Errorf("expected spinner output to contain 'Failed', got %q", out)
	}
}

func TestRemoveANSI(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Simple ANSI codes",
			input: Colorize("Hello", ColorRed, ColorStyleBold),
			want:  "Hello",
		},
		{
			name:  "No ANSI codes",
			input: "No ANSI codes here",
			want:  "No ANSI codes here",
		},
		{
			name:  "Multiple ANSI sequences",
			input: "Text " + Colorize("Red", ColorRed) + " and " + Colorize("Green", ColorGreen),
			want:  "Text Red and Green",
		},
		{
			name:  "ANSI only",
			input: Colorize("Blue Bold Text", ColorBlue, ColorStyleBold),
			want:  "Blue Bold Text",
		},
		{
			name:  "Empty ANSI",
			input: string(ColorBlue),
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeANSI(tt.input)
			if got != tt.want {
				t.Errorf("removeANSI(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestContext(t *testing.T) {
	var buf bytes.Buffer
	mesg := "Testing"
	ctx, cancel := context.WithCancel(context.Background())
	r := New(
		WithWriter(&buf),
		WithMesg(mesg),
		WithDoneMesg("Done"),
		WithContext(ctx),
		WithForceInteractive(),
		WithFrequency(10*time.Millisecond),
	)

	r.Start()
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(50 * time.Millisecond)

	isActive := r.IsRunning()
	if isActive {
		t.Fatalf("spinner should be inactive after context cancellation: %v", isActive)
	}
}

func TestContextCancelThenDone(t *testing.T) {
	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())
	r := New(
		WithWriter(&buf),
		WithMesg("Running..."),
		WithDoneMesg("Completed"),
		WithContext(ctx),
		WithForceInteractive(),
		WithFrequency(10*time.Millisecond),
	)

	r.Start()
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
	var buf bytes.Buffer
	timeout := 10 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	r := New(
		WithWriter(&buf),
		WithMesg("Timing out..."),
		WithContext(ctx),
		WithForceInteractive(),
		WithFrequency(5*time.Millisecond),
	)

	r.Start()
	time.Sleep(timeout + 50*time.Millisecond)

	isActive := r.IsRunning()

	if isActive {
		t.Fatalf("spinner should be inactive after context timeout: %v", isActive)
	}
}

func TestStartWithPreCancelledContext(t *testing.T) {
	var buf bytes.Buffer
	mesg := "Test Message"
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	r := New(
		WithWriter(&buf),
		WithMesg(mesg),
		WithContext(ctx),
		WithFrequency(5*time.Millisecond),
		WithForceInteractive(),
	)

	r.Start()
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
	tests := []struct {
		name     string
		symbol   string
		color    string
		msg      []string
		prefix   string
		expected string
	}{
		{
			name:     "symbol + message",
			symbol:   "✓",
			color:    "",
			msg:      []string{"done"},
			expected: "✓ done\n",
		},
		{
			name:     "color applied",
			symbol:   "✓",
			color:    "\x1b[32m",
			msg:      []string{"ok"},
			expected: "✓ ok\n",
		},
		{
			name:     "no symbol",
			color:    "",
			msg:      []string{"hello"},
			expected: "hello\n",
		},
		{
			name:     "prefix overrides symbol position",
			symbol:   "✓",
			color:    "",
			msg:      []string{"done"},
			prefix:   "step1",
			expected: "step1 ✓ done\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}

			r := &Rotato{
				Writer:     buf,
				prefixMesg: tt.prefix,
			}

			r.displayMessage(tt.symbol, tt.color, tt.msg...)

			got := buf.String()
			if got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestMesgDecorator(t *testing.T) {
	prefix := "prefix::"
	suffix := "::suffix"

	r := New(
		WithMesgDecorator(func(mesg string) string {
			return prefix + mesg + suffix
		}),
	)

	result := r.decorateMessage("hello")
	expected := prefix + "hello" + suffix

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestContextCancelled(t *testing.T) {
	var buf bytes.Buffer
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	r := New(
		WithWriter(&buf),
		WithContext(ctx),
		WithForceInteractive(),
		WithContextDoneHandler(func(r *Rotato, err error) {
			r.Fail(err.Error())
		}),
	)
	r.Start()

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
	var buf bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	r := New(
		WithWriter(&buf),
		WithPrefix("CtxDeadlineExceeded"),
		WithContext(ctx),
		WithForceInteractive(),
		WithContextDoneHandler(func(r *Rotato, err error) {
			r.Fail(err.Error())
		}),
	)
	r.Start()

	time.Sleep(20 * time.Millisecond)

	want := context.DeadlineExceeded.Error()
	r.writerMu.Lock()
	got := buf.String()
	r.writerMu.Unlock()

	if !strings.Contains(got, want) {
		t.Fatalf("expected %s, got %s", want, got)
	}
}
