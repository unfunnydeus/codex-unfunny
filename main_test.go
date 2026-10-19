package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	for _, tc := range []struct {
		name     string
		args     []string
		color    bool
		code     int
		contains string
		ansi     bool
	}{
		{"banner", nil, true, 0, "█", true},
		{"plain", []string{"--no-color"}, true, 0, "█", false},
		{"environment disables color", nil, false, 0, "█", false},
		{"help", []string{"--help"}, true, 0, "Usage:", false},
		{"version", []string{"--version"}, true, 0, "codex-unfunny 0.1.0\n", false},
		{"unknown flag", []string{"--unknown"}, true, 2, "", false},
		{"positional argument", []string{"hello"}, true, 2, "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out, errOut bytes.Buffer
			if got := run(tc.args, &out, &errOut, tc.color); got != tc.code {
				t.Fatalf("exit = %d, want %d", got, tc.code)
			}
			if !strings.Contains(out.String(), tc.contains) {
				t.Fatalf("unexpected output: %q", out.String())
			}
			if strings.Contains(out.String(), "\x1b[") != tc.ansi {
				t.Fatal("unexpected ANSI color state")
			}
			if tc.code != 0 && errOut.Len() == 0 {
				t.Fatal("missing error diagnostic")
			}
		})
	}
}

func TestColorsAllowed(t *testing.T) {
	for _, tc := range []struct {
		noColor, term string
		want          bool
	}{
		{"", "", true}, {"1", "", false}, {"0", "", false}, {"", "dumb", false}, {"", "xterm-256color", true},
	} {
		get := func(key string) string {
			if key == "NO_COLOR" {
				return tc.noColor
			}
			return tc.term
		}
		if got := colorsAllowed(get); got != tc.want {
			t.Errorf("NO_COLOR=%q TERM=%q: got %v", tc.noColor, tc.term, got)
		}
	}
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }

func TestWriteFailure(t *testing.T) {
	var errOut bytes.Buffer
	if code := run(nil, failingWriter{}, &errOut, false); code != 1 {
		t.Fatalf("exit = %d", code)
	}
	if !strings.Contains(errOut.String(), "write failed") {
		t.Fatal("missing write error")
	}
}
