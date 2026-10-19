package main

import (
	"fmt"
	"strings"
)

var glyphs = map[rune][5]string{
	'U': {"#   #", "#   #", "#   #", "#   #", " ### "},
	'N': {"#   #", "##  #", "# # #", "#  ##", "#   #"},
	'F': {"#####", "#    ", "#### ", "#    ", "#    "},
	'Y': {"#   #", " # # ", "  #  ", "  #  ", "  #  "},
}

func renderBanner(color bool) string {
	const word = "UNFUNNY"
	const width = len(word)*7 - 2
	var out strings.Builder
	out.WriteByte('\n')
	for row := 0; row < 5; row++ {
		var line strings.Builder
		for i, letter := range word {
			if i > 0 {
				line.WriteString("  ")
			}
			line.WriteString(glyphs[letter][row])
		}
		out.WriteString("  ")
		for x, pixel := range line.String() {
			if pixel == '#' {
				if color {
					r, g, b := gradient(x, width-1)
					fmt.Fprintf(&out, "\x1b[38;2;%d;%d;%dm", r, g, b)
				}
				out.WriteRune('█')
			} else {
				out.WriteByte(' ')
			}
		}
		if color {
			out.WriteString("\x1b[0m")
		}
		out.WriteByte('\n')
	}
	out.WriteByte('\n')
	return out.String()
}

// Interpolate through violet, blue, and cyan across the banner.
func gradient(x, max int) (int, int, int) {
	stops := [3][3]int{{180, 100, 255}, {90, 130, 255}, {50, 220, 255}}
	segment, offset := 0, x*2
	if offset > max {
		segment, offset = 1, offset-max
	}
	a, b := stops[segment], stops[segment+1]
	return a[0] + (b[0]-a[0])*offset/max,
		a[1] + (b[1]-a[1])*offset/max,
		a[2] + (b[2]-a[2])*offset/max
}
