---
marp: true
theme: lt-layout
paginate: true
html: true
---

# 実DBテストでデータ汚染を防ぐ

---

## 目次

<!--
  - 深く説明しない
  - すぐ飛ばしてよい
-->

1. 自己紹介
2. 何がつらいか
3. 解決策1: テストごとにDBを分離
4. トレードオフ
5. 解決策2: トランザクション
6. トレードオフ
7. 使い分け
8. まとめ

---

## 自己紹介

<!--
  - 深く説明しない
  - 名前
  - 業務アプリのバックエンドがメイン
  - Go / DB / AWS あたりに触れる
  - 最後によろしくお願いします
-->

- 名前 えび
- 経歴
  - 2021年9月 入社（現職）
  - エンジニア 5年目
- 技術
  - Open/Web系の業務アプリのバックエンドがメイン
  - 言語,F/W: Go / Laravel / Rails / React.ts
  - DB: MySQL,PostgreSQL
  - クラウド/インフラ: AWS
- 資格
  - 基本情報、AWS SAA、AWS SAP

---

## 実DBテストで何がつらいか

<!--
  - 早速本題
  - 取得系と更新系が同じデータを触る例
  - 実行順が不定だと壊れる
  - `users.id = 1` を共有しているのが問題
  - 共用DBだと並列実行でも不安定になりやすい

  - 次で解決策を2つ紹介する
 -->

<div style="display: flex; gap: 24px;">
<div style="flex: 1;">

### 取得系テスト

```go
func TestGetUserName(t *testing.T) {
    user := findUserByID(db, 1)
    if user.Name != "Alice" {
        t.Fatalf("got %s", user.Name)
    }
}
```

</div>

<div style="flex: 1;">

### 更新系テスト

```go
func TestUpdateUserName(t *testing.T) {
    db.Exec(`UPDATE users SET name = 'Bob' WHERE id = 1`)

    user := findUserByID(db, 1)
    if user.Name != "Bob" {
        t.Fatalf("got %s", user.Name)
    }
}
```

</div>
</div>

---

## 解決策1: テストごとにDBを分離

<!--
  - 解決策1はDB分離
  - ランダム名でDB作成
  - 呼び出し側でマイグレーション
  - `defer` で最後にDROP
  - テストごとに完全に独立した状態を作れて安定してテストができる

  - 次でトレードオフを説明
-->

テストごとに**専用のDBを作成・破棄**する

```go
func CreateNewTestDB() (*MySQL, func() error, error) {
    id, _ := generateUniqueID()            // nanoidでランダムID生成
    dbName := fmt.Sprintf("%s_test", id)

    createMySQLDatabase(dbName)            // CREATE DATABASE + GRANT

    db, _ := gorm.Open(mysql.Open(dsn(dbName)), &gorm.Config{})

    cleanUp := func() error {
        return deleteMySQLDatabase(dbName) // DROP DATABASE
    }

    return &MySQL{DB: db}, cleanUp, nil
}
```

---

## トレードオフ

<!--
  - DB分離のトレードオフ
  - 独立性が高い
  - マイグレーション漏れにも気づきやすい
  - ただし毎回 CREATE / DROP / migration で重い
  - 小規模なら選びやすい
  - 中規模以上はコストを見て判断

  - 次はもう1つの解決策
 -->

### メリット

- DBの共用データが壊れて他テストが落ちる事故がなくなる
- テストケースごとの独立性が高く、`t.Parallel()` しやすい
- マイグレーション自体が常にテストされる

### デメリット

- **速度**: テスト毎にCREATE&DROP&マイグレーションは重い
- **接続数**: 並列数だけ接続が増える（`max_connections` に注意）
- DB作成権限や初期化処理が必要になる

---

## 解決策2: トランザクション

<!--
  - 解決策2はトランザクション
  - Rails や Laravel でもよくある方法
  - テスト開始時にBEGIN
  - 最後にROLLBACK

  - 次でトレードオフを説明
 -->

テスト開始時に**トランザクションを張って、最後に破棄**する

```go
func TestUpdateUserName(t *testing.T) {
    tx := db.Begin()
    defer tx.Rollback()

    tx.Exec(`UPDATE users SET name = 'Bob' WHERE id = 1`)

    user := findUserByID(tx, 1)
    if user.Name != "Bob" {
        t.Fatalf("got %s", user.Name)
    }
}
```

---

## トレードオフ

<!--
  - トランザクションのトレードオフ
  - 速いのが最大のメリット
  - 既存DBを使い回せる
  - 行ロック次第で競合やデッドロックはありうる
  - 実装側で同じトランザクションを渡す設計も必要

  - 次で使い分け
 -->

### メリット

- `CREATE DATABASE` が不要なので速い
- 既存DBを使い回せる

### デメリット

- 同じ行を触る並列実行では、ロック待ちやデッドロックが起こりうる
- テスト対象コードを同じトランザクションに乗せる工夫がいる

---

## 使い分け

<!--
  - 使い分けの話
  - 基本はトランザクションを選ぶ
  - 理由は速度と開発体験
  - CI/CD 全体の速度にも効き、ユーザーへの価値提供が遅くなる
  - ただし独立性を最優先したいならDB分離
  - コミット後の状態確認が必要なときもDB分離

  - 次でまとめ
-->

### DB分離

- 独立性を最優先したい
- マイグレーションや初期化も含めて検証したい
- コミット後の状態変化まで含めて検証したい
- 実装都合でトランザクション管理が煩雑になる場合

### トランザクション

- テストを高速に回したい
- 既存基盤を活かしつつ改善したい
- コミット後の状態変化の検証が不要で、ロック競合を避けられる場合

---

## まとめ

- 実DBテストは、工夫しないとデータ汚染で不安定になりやすい
- テスト毎に**CREATE&DROP** する方法と、**トランザクション + ROLLBACK** する方法がある
- 独立性や検証範囲を重視するならDB分離、速度や既存基盤との整合を重視するならトランザクション
