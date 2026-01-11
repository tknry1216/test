# Go + PostgreSQL CRUD API

Go 1.25の標準ライブラリ（net/http）を使ったシンプルなユーザー管理CRUD APIです。

## 技術スタック

- ✅ **Go 1.25** - 標準ライブラリのみ使用
- ✅ **PostgreSQL 15** - Docker Composeで起動
- ✅ **net/http** - Go標準のHTTPサーバー（フレームワークなし）
- ✅ **pgx/v5** - PostgreSQLドライバ
- ✅ **internal/dbmgr** - DB接続とCRUD操作を管理

## プロジェクト構成

```
golang/
├── cmd/
│   └── main.go              # メインアプリケーション（HTTPハンドラー）
├── internal/
│   └── dbmgr/
│       └── dbmgr.go         # DB接続・CRUD操作
├── init/
│   └── 01_init.sql          # DB初期化スクリプト
├── docker-compose.yml       # PostgreSQL設定
├── go.mod                   # 依存関係管理
└── README.md
```

## セットアップ

### 1. データベース起動

```bash
docker-compose up -d
```

PostgreSQLコンテナが起動し、`init/01_init.sql`が自動実行されます。

### 2. 依存関係のインストール

```bash
go mod download
```

### 3. APIサーバー起動

```bash
go run cmd/main.go
```

サーバーは `http://localhost:8080` で起動します。

## API エンドポイント

### ヘルスチェック
```bash
curl http://localhost:8080/health
```

**レスポンス:**
```json
{"status": "ok"}
```

### 全ユーザー取得
```bash
curl http://localhost:8080/api/v1/users
```

**レスポンス例:**
```json
[
  {
    "id": 1,
    "name": "山田太郎",
    "email": "yamada@example.com",
    "created_at": "2026-01-11T10:00:00Z",
    "updated_at": "2026-01-11T10:00:00Z"
  }
]
```

### 特定ユーザー取得
```bash
curl http://localhost:8080/api/v1/users/1
```

### ユーザー作成
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "田中太郎",
    "email": "tanaka@example.com"
  }'
```

### ユーザー更新
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "田中次郎",
    "email": "tanaka2@example.com"
  }'
```

**注:** name、emailは両方とも任意（空でもOK）

### ユーザー削除
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## データベース構造

### users テーブル

| カラム名 | 型 | 説明 |
|---------|-----|------|
| id | SERIAL | 主キー |
| name | VARCHAR(100) | ユーザー名 |
| email | VARCHAR(100) | メールアドレス（ユニーク制約） |
| created_at | TIMESTAMP | 作成日時 |
| updated_at | TIMESTAMP | 更新日時 |

**初期データ:** サンプルユーザー3件が自動で挿入されます。

## 環境変数

| 変数名 | 説明 | デフォルト値 |
|-------|------|-------------|
| `DATABASE_URL` | PostgreSQL接続文字列 | `postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable` |
| `PORT` | APIサーバーのポート | `8080` |

## アーキテクチャ

### レイヤー構成

1. **HTTPハンドラー層** (`cmd/main.go`)
   - リクエストのルーティング
   - JSON パース/レスポンス
   - バリデーション

2. **DB管理層** (`internal/dbmgr`)
   - DB接続プール管理
   - CRUD操作の実装
   - エラーハンドリング

### 特徴

- **標準ライブラリのみ**: Ginなどのフレームワーク不使用
- **Go 1.25の新機能**: `http.NewServeMux()`のメソッドベースルーティング
- **パスパラメータ**: `r.PathValue("id")`で取得
- **依存性注入**: DBManagerを `internal/dbmgr`パッケージで分離

## 停止方法

```bash
# APIサーバー停止: Ctrl+C

# データベース停止
docker-compose down

# データベースとボリューム完全削除
docker-compose down -v
```

## 開発メモ

- Go 1.25から`http.NewServeMux`がメソッド（GET, POST等）とパスパラメータ`{id}`をネイティブサポート
- フレームワークなしでも十分実用的なREST APIが構築可能
- `internal/`パッケージで適切にレイヤー分離
