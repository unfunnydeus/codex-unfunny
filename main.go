package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "0.1.0"

const help = `codex-unfunny — a minimal terminal banner

Usage: codex-unfunny [options]

Options:
  --no-color  Disable ANSI colors
  --help      Show this help
  --version   Show the version

Set NO_COLOR to a non-empty value to disable colors.
`

func main() {
	color := colorsAllowed(os.Getenv) && enableANSI()
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, color))
}

func colorsAllowed(getenv func(string) string) bool {
	return getenv("NO_COLOR") == "" && getenv("TERM") != "dumb"
}

func run(args []string, out, errOut io.Writer, color bool) int {
	flags := flag.NewFlagSet("codex-unfunny", flag.ContinueOnError)
	flags.SetOutput(errOut)
	noColor := flags.Bool("no-color", false, "Disable ANSI colors")
	showHelp := flags.Bool("help", false, "Show help")
	showVersion := flags.Bool("version", false, "Show version")
	flags.Usage = func() { fmt.Fprint(errOut, help) }
	if err := flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}
	if flags.NArg() != 0 {
		fmt.Fprintln(errOut, "codex-unfunny: unexpected positional arguments")
		return 2
	}
	var output string
	switch {
	case *showHelp:
		output = help
	case *showVersion:
		output = "codex-unfunny " + version + "\n"
	default:
		output = renderBanner(color && !*noColor)
	}
	if _, err := io.WriteString(out, output); err != nil {
		fmt.Fprintf(errOut, "codex-unfunny: %v\n", err)
		return 1
	}
	return 0
}
