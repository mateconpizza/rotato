package rotato

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestSpinnerOutput(t *testing.T) {
	var buf bytes.Buffer
	mesg := "Testing"
	r := New(WithWriter(&buf), WithMesg(mesg), WithDoneMesg("Done"))
	r.frequency = 10 * time.Millisecond

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
		WithSpinnerFrequency(10*time.Millisecond),
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
		WithSpinnerFrequency(10*time.Millisecond),
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
