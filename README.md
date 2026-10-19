# codex-unfunny

A minimal Go CLI that prints **UNFUNNY** in a violet → blue → cyan terminal
gradient. Built with the Go standard library, with no third-party dependencies.

```text
  █   █  █   █  █████  █   █  █   █  █   █  █   █
  █   █  ██  █  █      █   █  ██  █  ██  █   █ █
  █   █  █ █ █  ████   █   █  █ █ █  █ █ █    █
  █   █  █  ██  █      █   █  █  ██  █  ██    █
   ███   █   █  █       ███   █   █  █   █    █
```

The preview above is plain text; the default terminal output is colored.

## Features

- A compact, five-row Unicode banner with a horizontal ANSI true-color gradient.
- Explicit `--no-color` and non-empty `NO_COLOR` environment variable support.
- Plain output when `TERM=dumb`.
- macOS, Linux, and Windows support; Windows console colors are enabled when
  available, with a plain-text fallback otherwise.
- No network access, configuration files, or external dependencies at runtime.

For the intended appearance, use a dark terminal with true-color support,
a monospace font containing `█`, and at least 49 columns. The banner has a fixed
width and does not resize. Color is enabled by default on macOS and Linux,
including redirected output; use `--no-color` when saving plain text.

## Install

Requires Go 1.22 or newer:

```sh
go install github.com/unfunnydeus/codex-unfunny@v0.1.0
```

Go places the executable in `GOBIN`, or in `$(go env GOPATH)/bin` when `GOBIN`
is unset. Add that directory to your `PATH` to run it by name.

## Usage

```sh
codex-unfunny
codex-unfunny --no-color
codex-unfunny --help
codex-unfunny --version
```

| Argument | Effect |
| --- | --- |
| `--no-color` | Print the banner without ANSI escape sequences. |
| `--help` | Print usage information. |
| `--version` | Print `codex-unfunny 0.1.0`. |

No positional arguments are accepted. Invalid arguments return exit code 2;
output failures return 1; successful invocations return 0. When both `--help`
and `--version` are supplied, help takes precedence.

Disable colors with the environment on macOS or Linux:

```sh
NO_COLOR=1 codex-unfunny
```

In Windows PowerShell:

```powershell
$env:NO_COLOR = '1'
codex-unfunny
```

Any non-empty `NO_COLOR` value disables color, including `0`. An empty value
does not disable color.

## Build from source

```sh
git clone https://github.com/unfunnydeus/codex-unfunny.git
cd codex-unfunny
go build -o codex-unfunny .
./codex-unfunny
```

On Windows, build with `go build -o codex-unfunny.exe .` and run
`.\codex-unfunny.exe`. During development, `go run .` works on all three platforms.

## Development and testing

```sh
go fmt ./...
go test ./...
go vet ./...
go build ./...
```

Tests cover the exact banner layout, gradient stops, ANSI resets, command-line
options, environment settings, invalid input, and output failures. GitHub Actions
runs formatting checks, tests, vet, and builds on macOS, Linux, and Windows.
The Windows console appearance still depends on the terminal and font in use.

## Project structure

```text
.
├── .github/
│   └── workflows/
│       └── ci.yml         # Cross-platform checks
├── .gitignore
├── LICENSE
├── README.md
├── ansi_unix.go          # ANSI support on non-Windows platforms
├── ansi_windows.go       # Windows virtual terminal support
├── banner.go             # Glyphs and gradient rendering
├── banner_test.go        # Layout and gradient tests
├── go.mod
├── main.go               # CLI options and process entry point
└── main_test.go          # CLI, environment, and error tests
```

One `main` package keeps the project easy to read. The renderer returns a string;
the CLI writes it to a supplied writer, allowing tests without starting a process.
Platform-specific files keep Windows console setup separate from rendering.
There is no `go.sum` because the module has no external dependencies.

## AI Development Experiment

This project was created as an experiment in autonomous AI agents for software
development.

**Unfunny ([@unfunnydeus](https://github.com/unfunnydeus))** originated the idea,
defined the requirements, directed the experiment, and owns the project.

**OpenAI Codex (AI Coding Agent)** implemented the code, organized the project,
prepared the documentation, and ran local tests and development checks using
real filesystem and terminal tools. The code is AI-authored; it is not presented
as code written manually by Unfunny.

The goal is to explore the capabilities of modern AI agents and their interaction
with real development tools on a small, understandable project.

## License

[MIT](LICENSE) — Copyright (c) 2026 Unfunny (@unfunnydeus).
