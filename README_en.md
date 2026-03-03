# vpp2txt

[![CI](https://github.com/ieee0824/vpp2txt/actions/workflows/ci.yml/badge.svg)](https://github.com/ieee0824/vpp2txt/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/ieee0824/vpp2txt/graph/badge.svg)](https://codecov.io/gh/ieee0824/vpp2txt)

A command-line tool that converts [Voice Peak](https://www.ah-soft.com/voice/) project files (`.vpp`) to text.

[日本語](README.md)

## Installation

### Homebrew (macOS / Linux)

```bash
brew install ieee0824/tap/vpp2txt
```

### Download binary

Download a prebuilt binary for your OS from the [Releases page](https://github.com/ieee0824/vpp2txt/releases/latest).

### go install

```bash
go install github.com/ieee0824/vpp2txt/cmd/vpp2txt@latest
```

### Build from source

```bash
git clone https://github.com/ieee0824/vpp2txt.git
cd vpp2txt
go build -o vpp2txt ./cmd/vpp2txt/
```

## Usage

```
usage: vpp2txt [options] <input.vpp> [input2.vpp ...]

options:
  -o <file>       output file (default: stdout)
  -format <type>  output format (default: script)
                    script  : "speaker: text" format
                    plain   : text only
  -h              show help
```

### Examples

```bash
# Output to stdout (script format)
vpp2txt project.vpp

# Output text only in plain format
vpp2txt -format plain project.vpp

# Save to file
vpp2txt -o output.txt project.vpp

# Convert multiple files
vpp2txt -o output.txt scene1.vpp scene2.vpp
```

### Output examples

**script format (default)**

```
重音テト: こんにちは
小春六花: よろしくね
重音テト: こちらこそ
```

**plain format**

```
こんにちは
よろしくね
こちらこそ
```

## About VPP files

`.vpp` files are JSON-formatted text files. Each block in `project.blocks` contains a speaker name (`narrator.key`) and dialogue lines (`sentence-list[].text`). This tool extracts them in order.

## License

MIT
