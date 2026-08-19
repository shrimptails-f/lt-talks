# QRコード生成

文字列から512×512のQRコードをPNG形式で生成するシンプルなGoツール。

## 実行

```bash
go -C tools run . \
  "https://example.com" \
  qrcode.png
```
