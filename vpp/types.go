package vpp

import "encoding/json"

// RawList はJSONの任意のリストを保持する (パース不要なフィールド用)
type RawList = json.RawMessage

// VPP はVoice PeakのプロジェクトファイルのトップレベルJSONを表す
type VPP struct {
	Version string           `json:"version"`
	Project Project          `json:"project"`
	Voices  map[string]Voice `json:"voices"`
}

// Project はプロジェクト全体の設定とブロックリストを保持する
type Project struct {
	Params         Params  `json:"params"`
	Emotions       RawList `json:"emotions"`
	GlobalEmotions RawList `json:"global-emotions"`
	GlobalSettings RawList `json:"global-settings"`
	Blocks         []Block `json:"blocks"`
}

// Params は音声パラメータを保持する
type Params struct {
	Speed  float64 `json:"speed"`
	Pitch  float64 `json:"pitch"`
	Pause  float64 `json:"pause"`
	Volume float64 `json:"volume"`
}

// Block はひとつのセリフブロックを表す
type Block struct {
	Narrator     Narrator   `json:"narrator"`
	SentenceList []Sentence `json:"sentence-list"`
}

// Narrator はセリフを話すキャラクター情報を保持する
type Narrator struct {
	Key      string `json:"key"`
	Language string `json:"language"`
}

// Sentence はひとつの文を表す
type Sentence struct {
	Text   string `json:"text"`
	HasEOS bool   `json:"has-eos"`
}

// Voice はキャラクターごとの設定を保持する
type Voice struct {
	Latest int    `json:"latest"`
	Nid    string `json:"nid"`
}

// Line はひとつのセリフ行を表す
type Line struct {
	Narrator string
	Text     string
}
