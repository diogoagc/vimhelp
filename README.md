# vimhelp

Simple Vim bindings cheat sheet for the terminal.

## Install

```bash
go install github.com/diogoagc/vimhelp@latest
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

## Custom bindings

By default, `vimhelp` uses the bindings included in the package.

Create your own editable bindings file:

```bash
vimhelp init
```

Show where your bindings file is located:

```bash
vimhelp config
```

After running `vimhelp init`, edit the generated `bindings.json` file.

Once the custom file exists, `vimhelp` automatically uses it instead of the embedded default bindings.

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

## Update

Install latest version:

```bash
go install github.com/diogoagc/vimhelp@latest
```

Install a specific version:

```bash
go install github.com/diogoagc/vimhelp@v0.1.1
```

## License

MIT