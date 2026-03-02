package vpp

import (
	"testing"
)

func TestLines_order(t *testing.T) {
	v := &VPP{
		Project: Project{
			Blocks: []Block{
				{
					Narrator: Narrator{Key: "テト"},
					SentenceList: []Sentence{
						{Text: "こんにちは", HasEOS: true},
						{Text: "さようなら", HasEOS: true},
					},
				},
				{
					Narrator: Narrator{Key: "六花"},
					SentenceList: []Sentence{
						{Text: "おはよう", HasEOS: true},
					},
				},
			},
		},
	}

	lines := v.Lines()
	if len(lines) != 3 {
		t.Fatalf("Lines() count = %d, want 3", len(lines))
	}

	wants := []Line{
		{Narrator: "テト", Text: "こんにちは"},
		{Narrator: "テト", Text: "さようなら"},
		{Narrator: "六花", Text: "おはよう"},
	}
	for i, want := range wants {
		if lines[i] != want {
			t.Errorf("Lines()[%d] = %+v, want %+v", i, lines[i], want)
		}
	}
}

func TestLines_skipsEmptyText(t *testing.T) {
	v := &VPP{
		Project: Project{
			Blocks: []Block{
				{
					Narrator: Narrator{Key: "テト"},
					SentenceList: []Sentence{
						{Text: "", HasEOS: false},
						{Text: "こんにちは", HasEOS: true},
					},
				},
			},
		},
	}

	lines := v.Lines()
	if len(lines) != 1 {
		t.Fatalf("Lines() count = %d, want 1 (empty text should be skipped)", len(lines))
	}
	if lines[0].Text != "こんにちは" {
		t.Errorf("Lines()[0].Text = %q, want %q", lines[0].Text, "こんにちは")
	}
}

func TestLines_emptyBlocks(t *testing.T) {
	v := &VPP{}
	if lines := v.Lines(); lines != nil {
		t.Errorf("Lines() on empty VPP = %v, want nil", lines)
	}
}
