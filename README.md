# vimhelp

Simple Vim bindings cheat sheet for the terminal.

## Install

```bash
go install github.com/YOUR_USERNAME/vimhelp@latest
```

Make sure Go's bin folder is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

Show all bindings:

```bash
vimhelp
```

Show available topics:

```bash
vimhelp topics
```

Show bindings for a topic:

```bash
vimhelp movement
vimhelp editing
```

Show help for a specific key:

```bash
vimhelp dd
vimhelp 'ci"'
```

## Development

Run locally:

```bash
go run .
```

Build:

```bash
go build -o vimhelp
./vimhelp
```

## License

MIT