package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestBanner(t *testing.T) {
	want := "\n" +
		"  █   █  █   █  █████  █   █  █   █  █   █  █   █\n" +
		"  █   █  ██  █  █      █   █  ██  █  ██  █   █ █ \n" +
		"  █   █  █ █ █  ████   █   █  █ █ █  █ █ █    █  \n" +
		"  █   █  █  ██  █      █   █  █  ██  █  ██    █  \n" +
		"   ███   █   █  █       ███   █   █  █   █    █  \n\n"
	if got := renderBanner(false); got != want {
		t.Fatalf("banner mismatch:\n%s", got)
	}
	colored := renderBanner(true)
	ansi := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	if ansi.ReplaceAllString(colored, "") != want {
		t.Fatal("color changes banner layout")
	}
	if strings.Count(colored, "\x1b[0m") != 5 {
		t.Fatal("each row must reset color")
	}
}

func TestGradient(t *testing.T) {
	for _, tc := range []struct {
		x    int
		want [3]int
	}{
		{0, [3]int{180, 100, 255}}, {23, [3]int{90, 130, 255}}, {46, [3]int{50, 220, 255}},
	} {
		r, g, b := gradient(tc.x, 46)
		if got := [3]int{r, g, b}; got != tc.want {
			t.Errorf("gradient(%d) = %v, want %v", tc.x, got, tc.want)
		}
	}
}
