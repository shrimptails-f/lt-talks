---
marp: true
theme: default
paginate: true
html: true
---

# HTML埋め込みサンプル

Markdownの中にHTMLを混在できる

---

## 普通のMarkdown

- 箇条書き
- **太字** や _斜体_
- `コード`

---

## HTML埋め込み: 2カラムレイアウト

<div style="display: flex; gap: 2rem;">
<div style="flex: 1; background: #e8f4fd; padding: 1rem; border-radius: 8px;">

### 左カラム

Markdownも**普通に**書ける

</div>
<div style="flex: 1; background: #fdecea; padding: 1rem; border-radius: 8px;">

### 右カラム

- リストも使える
- こんな感じで

</div>
</div>

---

## HTML埋め込み: バッジ・ラベル

リリース状況:
<span style="background: #28a745; color: white; padding: 2px 10px; border-radius: 12px; font-size: 0.8em;">完了</span>
<span style="background: #ffc107; color: black; padding: 2px 10px; border-radius: 12px; font-size: 0.8em;">進行中</span>
<span style="background: #dc3545; color: white; padding: 2px 10px; border-radius: 12px; font-size: 0.8em;">未着手</span>

---

## HTML埋め込み: テーブルをスタイリング

<table style="width: 100%; border-collapse: collapse; font-size: 0.9em;">
  <thead>
    <tr style="background: #4a90d9; color: white;">
      <th style="padding: 8px;">機能</th>
      <th style="padding: 8px;">担当</th>
      <th style="padding: 8px;">状態</th>
    </tr>
  </thead>
  <tbody>
    <tr style="background: #f0f7ff;">
      <td style="padding: 8px;">ログイン</td>
      <td style="padding: 8px;">田中</td>
      <td style="padding: 8px;">✅ 完了</td>
    </tr>
    <tr>
      <td style="padding: 8px;">決済</td>
      <td style="padding: 8px;">鈴木</td>
      <td style="padding: 8px;">🚧 進行中</td>
    </tr>
  </tbody>
</table>

---

## HTML埋め込み: グラデーション背景ブロック

<div style="background: linear-gradient(135deg, #667eea, #764ba2); color: white; padding: 2rem; border-radius: 12px; text-align: center;">

## まとめ

HTMLを使えばデザインの幅が広がる

</div>

---

## sample テーマ紹介

<div style="display: flex; gap: 1.5rem; margin-top: 1rem;">
<div style="flex: 1; border-left: 5px solid #FFBE00; padding: 1rem; background: #FFFDF0; border-radius: 0 8px 8px 0;">

### カラーパレット

- メイン: <span style="background:#FFBE00; color:#1a1a1a; padding:2px 10px; border-radius:4px; font-size:0.8em;">#FFBE00</span>
- アクセント: <span style="background:#D93025; color:white; padding:2px 10px; border-radius:4px; font-size:0.8em;">#D93025</span>

</div>
<div style="flex: 1; border-left: 5px solid #D93025; padding: 1rem; background: #FFF5F5; border-radius: 0 8px 8px 0;">

### 使い方

```markdown
---
marp: true
theme: sample
---
```

</div>
</div>
