# go_houdini
<a href="https://github.com/luizfiuzaa/go_houdini/blob/main/doc/README_pt-br.md">[pt-br]</a>

A CLI tool to quickly scaffold `.go` files with a defined structure — package, imports, and function stubs — without manual boilerplate.

![Gopher wizard](doc/assets/gopher.png)

## Features

- Interactive prompts to scaffold a new Go file
- Configurable package name (defaults to `main`)
- Imports and function stubs generated from space-separated input
- Prevents overwriting existing files
- Zero external dependencies

## Requirements

- [Go 1.23+](https://golang.org/dl/)

## Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/luizfiuzaa/go_houdini.git
cd go_houdini
go build -o houdini .
```

To use `houdini` from any directory, move the binary to a folder in your `PATH`:

```sh
# Linux / macOS
mv houdini /usr/local/bin/

# Windows — move houdini.exe to a folder already in your PATH, e.g. C:\tools\
```

## Usage

### Create a new Go file

```sh
houdini create
```

You can also pass the filename directly to skip the name prompt:

```sh
houdini create hello_world
```

**Example session:**

```
📦 Package name (default 'main'): main
📚 Imports (space separated): fmt time
🔧 Functions (space separated): initApp startServer

✅ File created successfully! Enjoy!
```

This generates `hello_world.go`:

```go
package main

import(
  "fmt"
  "time"
)

func initApp() {}

func startServer() {}
```

### Help

```sh
houdini --help
houdini -h
```

## Contributing

Pull requests are welcome. For significant changes, please open an issue first to discuss what you'd like to change.

## License

[MIT](LICENSE)
