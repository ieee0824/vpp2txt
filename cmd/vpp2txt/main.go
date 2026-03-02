package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const usage = `usage: vpp2txt [options] <input.vpp> [input2.vpp ...]

Voice Peak (.vpp) ファイルをテキストに変換します。

options:
  -o <file>       出力先ファイル (デフォルト: 標準出力)
  -format <type>  出力形式 (デフォルト: script)
                    script  : "話者名: テキスト" 形式
                    plain   : テキストのみ
  -h              ヘルプを表示
`

func main() {
	outputFile := flag.String("o", "", "出力先ファイル (デフォルト: 標準出力)")
	format := flag.String("format", "script", "出力形式: script | plain")
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, usage)
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
