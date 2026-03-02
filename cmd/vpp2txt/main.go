package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const usageJa = `usage: vpp2txt [options] <input.vpp> [input2.vpp ...]

Voice Peak (.vpp) ファイルをテキストに変換します。

options:
  -o <file>       出力先ファイル (デフォルト: 標準出力)
  -format <type>  出力形式 (デフォルト: script)
                    script  : "話者名: テキスト" 形式
                    plain   : テキストのみ
  -h              ヘルプを表示
`

const usageEn = `usage: vpp2txt [options] <input.vpp> [input2.vpp ...]

Convert Voice Peak (.vpp) files to text.

options:
  -o <file>       output file (default: stdout)
  -format <type>  output format (default: script)
                    script  : "speaker: text" format
                    plain   : text only
  -h              show help
`

func isJapanese() bool {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if v := os.Getenv(key); v != "" {
			return strings.HasPrefix(strings.ToLower(v), "ja")
		}
	}
	return false
}

func usageText() string {
	if isJapanese() {
		return usageJa
	}
	return usageEn
}

func main() {
	outputFile := flag.String("o", "", "output file")
	format := flag.String("format", "script", "output format: script | plain")
	flag.Usage = func() { fmt.Fprint(os.Stderr, usageText()) }
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, usageText())
		os.Exit(1)
	}

	w := io.Writer(os.Stdout)
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}

	for _, path := range args {
		if err := convert(w, path, *format); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", path, err)
			os.Exit(1)
		}
	}
}
