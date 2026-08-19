# 実DBテストでデータ汚染を防ぐ

実DBを使うテストで、更新系と取得系が同じデータを共有すると壊れやすい、という話。

- 方針1: テストごとにDBを分離する
- 方針2: トランザクションを張って最後にロールバックする
- どちらを選ぶかは、独立性を取るか速度を取るかのトレードオフ

## 出力

```bash
marp lt-talks/2026-05-go-db-testing/slide.md \
  --html \
  --theme-set ./themes/lt-layout.css \
  --allow-local-files \
  -o lt-talks/2026-05-go-db-testing/slide.html
```
