# このリポジトリについて

## 環境構築手順

### ソースをクローン

### .envをコピー

```bash
cp .devcontainer/.env.sample .devcontainer/.env
```

### VsCodeでプロジェクトフォルダーを開く

### Reopen in Containerを押下

Ctrl Shift P → Reopen in containerと入力して実行

以上で環境構築完了です。

## スライド出力

[Marp](https://marp.app/) を使ってMarkdownからHTMLスライドを生成します。

### ファイル構成

| ファイル     | 説明                     |
| ------------ | ------------------------ |
| `slide.md`   | スライドのソースファイル |
| `slide.html` | 生成されるHTMLファイル   |

### 出力コマンド

```bash
task slide
```

`slide.md` と同じディレクトリに `slide.html` が生成されます。
