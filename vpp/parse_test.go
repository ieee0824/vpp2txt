package vpp

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_minimal(t *testing.T) {
	v, err := Parse(filepath.Join("testdata", "minimal.vpp"))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if v.Version != "1.0.0" {
		t.Errorf("Version = %q, want %q", v.Version, "1.0.0")
	}
	if len(v.Project.Blocks) != 2 {
		t.Errorf("Blocks count = %d, want 2", len(v.Project.Blocks))
	}
}

func TestParse_trailingNullByte(t *testing.T) {
	// ファイル末尾に \x00 が付いていてもパースできることを確認する
	src, err := os.ReadFile(filepath.Join("testdata", "minimal.vpp"))
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := os.CreateTemp(t.TempDir(), "*.vpp")
	if err != nil {
		t.Fatal(err)
	}
	tmp.Write(src)
	tmp.Write([]byte{0x00})
	tmp.Close()

	if _, err := Parse(tmp.Name()); err != nil {
		t.Fatalf("Parse() with trailing null byte error = %v", err)
	}
}

func TestParse_fileNotFound(t *testing.T) {
	if _, err := Parse("nonexistent.vpp"); err == nil {
		t.Error("Parse() expected error for missing file, got nil")
	}
}

func TestParse_invalidJSON(t *testing.T) {
	tmp, err := os.CreateTemp(t.TempDir(), "*.vpp")
	if err != nil {
		t.Fatal(err)
	}
	tmp.WriteString("not valid json")
	tmp.Close()

	if _, err := Parse(tmp.Name()); err == nil {
		t.Error("Parse() expected error for invalid JSON, got nil")
	}
}
