# vpp2txt

[![CI](https://github.com/ieee0824/vpp2txt/actions/workflows/ci.yml/badge.svg)](https://github.com/ieee0824/vpp2txt/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/ieee0824/vpp2txt/graph/badge.svg)](https://codecov.io/gh/ieee0824/vpp2txt)

[Voice Peak](https://www.ah-soft.com/voice/) のプロジェクトファイル (`.vpp`) をテキストファイルに変換するコマンドラインツールです。

[English](README_en.md)

## インストール

[Releases ページ](https://github.com/ieee0824/vpp2txt/releases/latest)からお使いの OS に合ったバイナリをダウンロードできます。

または `go install` でインストール:

```bash
go install github.com/ieee0824/vpp2txt/cmd/vpp2txt@latest
```

リポジトリをクローンしてビルドすることもできます:

```bash
git clone https://github.com/ieee0824/vpp2txt.git
cd vpp2txt
go build -o vpp2txt ./cmd/vpp2txt/
```

## 使い方

```
usage: vpp2txt [options] <input.vpp> [input2.vpp ...]

options:
  -o <file>       出力先ファイル (デフォルト: 標準出力)
  -format <type>  出力形式 (デフォルト: script)
                    script  : "話者名: テキスト" 形式
                    plain   : テキストのみ
  -h              ヘルプを表示
```

### 例

```bash
# 標準出力に出力 (script 形式)
vpp2txt project.vpp

# plain 形式でテキストのみ出力
vpp2txt -format plain project.vpp

# ファイルに保存
vpp2txt -o output.txt project.vpp

# 複数ファイルをまとめて変換
vpp2txt -o output.txt scene1.vpp scene2.vpp
```

### 出力例

**script 形式 (デフォルト)**

```
重音テト: こんにちは
小春六花: よろしくね
重音テト: こちらこそ
```

**plain 形式**

```
こんにちは
よろしくね
こちらこそ
```

## VPP ファイルについて

`.vpp` ファイルは JSON 形式のテキストファイルです。`project.blocks` 内の各ブロックに話者名 (`narrator.key`) とセリフ (`sentence-list[].text`) が格納されており、本ツールはそれを順に抽出します。

## ライセンス

MIT
