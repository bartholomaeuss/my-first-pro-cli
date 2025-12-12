# my-first-pro-cli

A tiny [Cobra](https://github.com/spf13/cobra)-based CLI to get comfortable with building my own command-line tools in Go. The project keeps things intentionally simple: one namespace (`file`) and two subcommands (`create`, `delete`) so you can focus on the workflow around Cobra and the Go toolchain.

## Requirements

- Go 1.22 or newer (`go env GOVERSION` shows your version)
- Any modern terminal (macOS, Linux, Windows, or WSL)

## Quick start

```bash
git clone https://github.com/bartholomaeuss/my-first-pro-cli.git
cd my-first-pro-cli
go run . --help
```

## Build & installation

| Target | Command |
| --- | --- |
| Local binary for your current OS | `go build -o bin/my-first-pro-cli .` |
| Install into `$GOBIN` | `go install github.com/bartholomaeuss/my-first-pro-cli@latest` |
| Cross-compile for Windows x64 | `GOOS=windows GOARCH=amd64 go build -buildvcs=false -o my-first-pro-cli.exe .` |

> Tip: Once you add a `version` variable to `main.go`, append `-ldflags "-X main.version=<tag>"` so the build metadata is baked into the binary.

## Usage

All commands live under the `file` namespace:

```bash
$ my-first-pro-cli --help
CLI Tool for getting to know the magic
Usage:
  my-first-pro-cli [command]
```

### Command overview

| Command | Description | Example |
| --- | --- | --- |
| `my-first-pro-cli file create <path>` | Creates a file named `testfile` below the provided directory and writes `test` into it. | `my-first-pro-cli file create ./playground` |
| `my-first-pro-cli file delete <path>` | Removes a file named `testfile` relative to the current working directory. The argument currently exists only because Cobra expects one. | `my-first-pro-cli file delete .` |

```bash
# Create the file
mkdir -p ./playground
my-first-pro-cli file create ./playground
cat ./playground/testfile

# Delete it again
my-first-pro-cli file delete .
```

### Known quirks

- `file delete` always targets `testfile`; adjust `cmd/delete.go` if you need flexible paths.
- `os.RemoveAll` does not error when the file is missing, so the command reports success even if nothing was deleted.

## Development

- Format & lint: `go fmt ./...`
- Quick manual test: `go run . file create .` and `go run . file delete .`
- Cobra commands live under `cmd/`; new subcommands are wired up via `cmd/root.go`.

## License

See `LICENSE` for details.
