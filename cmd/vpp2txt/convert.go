package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ieee0824/vpp2txt/vpp"
)

func convert(w io.Writer, path string, format string) error {
	v, err := vpp.Parse(path)
	if err != nil {
		return err
	}

	for _, line := range v.Lines() {
		var out string
		switch strings.ToLower(format) {
		case "plain":
			out = line.Text
		default: // script
			out = line.Narrator + ": " + line.Text
		}
		fmt.Fprintln(w, out)
	}
	return nil
}
