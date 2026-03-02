package vpp

import (
	"encoding/json"
	"os"
	"strings"
)

// Parse はVPPファイルを読み込んでパースする
func Parse(path string) (*VPP, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// ファイル末尾のヌルバイトを除去する
	content := strings.TrimRight(string(data), "\x00")

	var v VPP
	if err := json.Unmarshal([]byte(content), &v); err != nil {
		return nil, err
	}
	return &v, nil
}
