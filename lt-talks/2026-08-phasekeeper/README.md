# AI開発の進行管理を、自作ワークフローエンジンに任せてみた

AIとの会話の中で人間が担っていた進行管理を、Go製のワークフローエンジン「Phasekeeper」へ任せてみた話。

- Step FunctionsとAWS AI-DLCから得た着想
- AIとEngineの責務分離
- 会話内の暗黙的な状態を、manifestやPhase Resultへ外出し
- Human Gate、Retry、Rollbackによるループ制御

## 出力

```bash
marp lt-talks/2026-08-phasekeeper/phasekeeper-slide.md \
  --html \
  --theme-set ./themes/lt-layout.css \
  --allow-local-files \
  -o lt-talks/2026-08-phasekeeper/phasekeeper-slide.html
```
