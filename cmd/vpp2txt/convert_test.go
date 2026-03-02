package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

var testVPP = filepath.Join("..", "..", "vpp", "testdata", "minimal.vpp")

func TestConvert_scriptFormat(t *testing.T) {
	var buf bytes.Buffer
	if err := convert(&buf, testVPP, "script"); err != nil {
		t.Fatalf("convert() error = %v", err)
	}

	out := buf.String()
	wantLines := []string{
		"テト: こんにちは",
		"テト: さようなら",
		"六花: おはよう",
	}
	for _, want := range wantLines {
		if !strings.Contains(out, want) {
			t.Errorf("output does not contain %q\ngot:\n%s", want, out)
		}
	}
}

func TestConvert_plainFormat(t *testing.T) {
	var buf bytes.Buffer
	if err := convert(&buf, testVPP, "plain"); err != nil {
		t.Fatalf("convert() error = %v", err)
	}

	out := buf.String()
	// 話者名が含まれないことを確認
	if strings.Contains(out, "テト:") || strings.Contains(out, "六花:") {
		t.Errorf("plain format should not contain narrator names, got:\n%s", out)
	}
	// テキストは含まれることを確認
	for _, want := range []string{"こんにちは", "さようなら", "おはよう"} {
		if !strings.Contains(out, want) {
			t.Errorf("output does not contain %q\ngot:\n%s", want, out)
		}
	}
}

func TestConvert_formatCaseInsensitive(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	convert(&buf1, testVPP, "plain")
	convert(&buf2, testVPP, "PLAIN")
	if buf1.String() != buf2.String() {
		t.Error("format should be case-insensitive")
	}
}

func TestConvert_fileNotFound(t *testing.T) {
	var buf bytes.Buffer
	if err := convert(&buf, "nonexistent.vpp", "script"); err == nil {
		t.Error("convert() expected error for missing file, got nil")
	}
}
