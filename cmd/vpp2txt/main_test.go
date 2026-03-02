package main

import "testing"

func TestIsJapanese(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want bool
	}{
		{
			name: "LANG=ja_JP.UTF-8",
			env:  map[string]string{"LANG": "ja_JP.UTF-8"},
			want: true,
		},
		{
			name: "LANG=en_US.UTF-8",
			env:  map[string]string{"LANG": "en_US.UTF-8"},
			want: false,
		},
		{
			name: "LC_ALL overrides LANG",
			env:  map[string]string{"LANG": "ja_JP.UTF-8", "LC_ALL": "en_US.UTF-8"},
			want: false,
		},
		{
			name: "LC_MESSAGES overrides LANG",
			env:  map[string]string{"LANG": "en_US.UTF-8", "LC_MESSAGES": "ja_JP.UTF-8"},
			want: true,
		},
		{
			name: "LC_ALL overrides LC_MESSAGES",
			env:  map[string]string{"LC_MESSAGES": "ja_JP.UTF-8", "LC_ALL": "en_US.UTF-8"},
			want: false,
		},
		{
			name: "no env vars defaults to English",
			env:  map[string]string{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
				t.Setenv(key, "")
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			if got := isJapanese(); got != tt.want {
				t.Errorf("isJapanese() = %v, want %v", got, tt.want)
			}
		})
	}
}
