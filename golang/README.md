# Go + PostgreSQL CRUD API

Go 1.25の標準ライブラリ（net/http）を使ったシンプルなユーザー管理CRUD APIです。

## 技術スタック

- ✅ **Go 1.25** - 標準ライブラリのみ使用
- ✅ **PostgreSQL 15** - Docker Composeで起動
- ✅ **net/http** - Go標準のHTTPサーバー（フレームワークなし）
- ✅ **Cobra** - CLIツール構築フレームワーク
- ✅ **pgx/v5** - PostgreSQLドライバ
- ✅ **internal/dbmgr** - DB接続とCRUD操作を管理

## プロジェクト構成

```
golang/
├── cmd/
│   ├── api/
│   │   └── main.go          # HTTPサーバー
│   └── cli/
│       ├── main.go          # CLIツールのエントリーポイント
│       └── cmd/
│           ├── root.go      # Cobraルートコマンド
│           └── users.go     # usersコマンド（list/get/create）
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

### 3-A. APIサーバー起動

```bash
go run cmd/api/main.go
```

サーバーは `http://localhost:8080` で起動します。

### 3-B. CLIツール（オプション）

```bash
# ビルド
go build -o golang-cli cmd/cli/main.go

# または直接実行
go run cmd/cli/main.go [command]
```

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

---

## CLI コマンド

CobraベースのCLIツールでDB操作が可能です。

### ヘルプ表示
```bash
go run cmd/cli/main.go --help
```

### 全ユーザー一覧 (GET)
```bash
go run cmd/cli/main.go users list
```

**出力例:**
```json
[
  {
    "id": 1,
    "name": "山田太郎",
    "email": "yamada@example.com",
    "created_at": "2026-01-11T07:48:02.570287Z",
    "updated_at": "2026-01-11T07:48:02.570287Z"
  }
]
```

### 特定ユーザー取得 (GET)
```bash
go run cmd/cli/main.go users get 1
```

### ユーザー作成 (POST)
```bash
go run cmd/cli/main.go users create --name "中村花子" --email "nakamura@example.com"

# 省略形
go run cmd/cli/main.go users create -n "中村花子" -e "nakamura@example.com"
```

### カスタムDB接続文字列
```bash
go run cmd/cli/main.go --db-url "postgres://user:pass@localhost:5432/mydb" users list
```

---

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

1. **プレゼンテーション層**
   - **HTTPハンドラー** (`cmd/api/main.go`) - REST API
   - **CLIツール** (`cmd/cli/`) - Cobraコマンド

2. **DB管理層** (`internal/dbmgr`)
   - DB接続プール管理
   - CRUD操作の実装
   - エラーハンドリング

### 特徴

- **標準ライブラリのみ**: Ginなどのフレームワーク不使用（HTTP API部分）
- **Go 1.25の新機能**: `http.NewServeMux()`のメソッドベースルーティング
- **パスパラメータ**: `r.PathValue("id")`で取得
- **Cobra CLI**: コマンドライン操作も可能
- **依存性注入**: DBManagerを `internal/dbmgr`パッケージで分離
- **マルチエントリーポイント**: API/CLIの両方に対応

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
