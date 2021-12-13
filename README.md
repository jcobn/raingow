# rainGOw
Simple program that makes the input rainbow
![demo photo](https://i.imgur.com/A2swqCZ_d.webp?maxwidth=760&fidelity=grand)

# Installation
1. Go to the [releases tab](https://github.com/jcobn/raingow/releases)
2. Download the `raingow.exe` binary from the latest release
3. (Optional) put the binary into your Path so you can run it from the command line

# Usage
```bash
echo "Hello World" | raingow (-c)

raingow -i "Hello World" (-c)
```
### Flags:
`-c` - Clears the terminal before printing the text (tested on Windows)

`-i` - Input string

# Building
1. Make sure you have [Go](https://go.dev/dl/) installed
2. `go build -o raingow.exe ./cmd/main.go`