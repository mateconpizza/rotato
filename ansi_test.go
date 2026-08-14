package rotato

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Unsetenv("NO_COLOR")
	code := m.Run()
	os.Exit(code)
}

func TestPaletteFormat(t *testing.T) {
	tests := []struct {
		name    string
		c       Color
		text    string
		want    string
		enabled bool
	}{
		{name: "enabled_with_color_and_text", c: FgGreen, text: "hello", want: FgGreen.Sprint("hello"), enabled: true},
		{name: "disabled_returns_text_unchanged", c: FgGreen, text: "hello", want: "hello"},
		{name: "empty_text_returns_empty", c: FgGreen, text: "", want: "", enabled: true},
		{name: "empty_color_returns_text_unchanged", c: "", text: "hello", want: "hello"},
		{name: "disabled_and_empty_color_and_text", c: "", text: "", want: ""},
		{name: "enabled_with_multi_char_text", c: FgRed, text: "error occurred", want: FgRed.Sprint("error occurred"), enabled: true},
		{name: "enabled_but_color_and_text_both_empty", c: "", text: "", want: "", enabled: true},
		{name: "disabled_with_color_set_and_multi_char_text", c: FgBlue, text: "no color please", want: "no color please"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newPalette()
			p.spinner = tt.c
			p.Enabled = tt.enabled

			got := p.Format(tt.c, tt.text)
			if got != tt.want {
				t.Fatalf("Format(%q, %q) = %q; want %q", tt.c, tt.text, got, tt.want)
			}
		})
	}
}

func TestNewPalette(t *testing.T) {
	tests := []struct {
		name        string
		noColorSet  bool
		wantEnabled bool
	}{
		{
			name:        "no_color_unset_enabled",
			noColorSet:  false,
			wantEnabled: true,
		},
		{
			name:        "no_color_set_disabled",
			noColorSet:  true,
			wantEnabled: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.noColorSet {
				t.Setenv("NO_COLOR", "1")
			}

			p := newPalette()
			if p.Enabled != tt.wantEnabled {
				t.Fatalf("newPalette().Enabled = %v; want %v", p.Enabled, tt.wantEnabled)
			}
		})
	}
}

func TestPalette_Formatters(t *testing.T) {
	t.Parallel()

	p := Palette{
		Enabled:     true,
		spinner:     FgBlue,
		message:     Color("c2"),
		prefixMesg:  Color("c3"),
		delimiter:   Color("c4"),
		doneMessage: Color("c5"),
		doneSymbol:  Color("c6"),
		failMessage: Color("c7"),
		failSymbol:  Color("c8"),
	}

	text := "test_string"

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"spinner", p.Spinner(text), p.spinner.Sprint(text)},
		{"message", p.Message(text), p.message.Sprint(text)},
		{"prefix", p.Prefix(text), p.prefixMesg.Sprint(text)},
		{"delimiter", p.Delimiter(text), p.delimiter.Sprint(text)},
		{"done_msg", p.DoneMsg(text), p.doneMessage.Sprint(text)},
		{"done_symbol", p.DoneSymbol(text), p.doneSymbol.Sprint(text)},
		{"fail_msg", p.FailMsg(text), p.failMessage.Sprint(text)},
		{"fail_symbol", p.FailSymbol(text), p.failSymbol.Sprint(text)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.got != tt.want {
				t.Fatalf("%s() = %q; want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
