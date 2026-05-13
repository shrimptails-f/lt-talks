# AGENTS.md

このリポジトリは [Marp](https://marp.app/) を使ってMarkdownからスライドを作成するための環境です。

## ディレクトリ構成

```text
themes/       # 共通テーマ管理ディレクトリ（CSSファイルを置く）
sample/       # サンプルスライドプロジェクト
  slide.md    # スライドのソース
  slide.html  # 生成されるHTML（コミット不要）
  slide.pdf   # 生成されるPDF（コミット不要）
```

## テーマ管理

テーマCSSは `themes/` ディレクトリで一元管理します。

- ファイル先頭に `/* @theme <テーマ名> */` が必要（Marpの認識に必須）
- `@import "default";` でベーステーマを継承してから上書きする

現在のテーマ: `themes/sample.css`

## スライドのビルド

Marpサーバーを起動してホットリロードプレビュー＆HTML出力

```bash
 marp --server sample --html --theme ./themes/sample.css
```
