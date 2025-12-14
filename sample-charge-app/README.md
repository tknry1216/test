## servicechargeservice

Go とクリーンアーキテクチャ構成のベース実装です。標準 `net/http` とメモリリポジトリを用いた最小構成。

### ディレクトリ構成

```
.
├── cmd/
│   └── servicechargeservice/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── domain/
│   │   ├── repository.go
│   │   └── servicecharge.go
│   ├── usecase/
│   │   └── servicecharge.go
│   ├── interface/
│   │   └── http/
│   │       └── handler.go
│   └── infrastructure/
│       └── repository/
│           └── memory/
│               └── servicecharge_memory.go
├── env.sample
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod
```

### 使い方

- ローカル:

```bash
cp env.sample .env
make run
```

- Docker:

```bash
cp env.sample .env
make docker-build
make docker-run
```

### エンドポイント

- GET `/health`
- GET `/service-charges`
- GET `/service-charges/{id}`
- POST `/service-charges` 例: `{ "name":"Shipping", "rate":0.1 }`
- DELETE `/service-charges/{id}`

## servicechargeservice

Go とクリーンアーキテクチャ構成のベース実装です。HTTP API（標準ライブラリのみ）とメモリリポジトリを含む最小構成で起動できます。

### ディレクトリ構成

```
.
├── cmd/
│   └── servicechargeservice/
│       └── main.go               # エントリポイント（DIとHTTPサーバ起動）
├── internal/
│   ├── config/
│   │   └── config.go             # 環境変数の読み込み
│   ├── domain/                   # エンティティ & リポジトリIF（ビジネスルール）
│   │   ├── repository.go
│   │   └── servicecharge.go
│   ├── usecase/                  # ユースケース（アプリケーションビジネスルール）
│   │   └── servicecharge.go
│   ├── interface/                # インターフェースアダプタ（入出力）
│   │   └── http/
│   │       └── handler.go
│   └── infrastructure/           # フレームワーク/ドライバ（実装詳細）
│       └── repository/
│           └── memory/
│               └── servicecharge_memory.go
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod
```

### 使い方

- ローカル起動:

```bash
cp .env.example .env
make run
```

- Docker 起動:

```bash
cp .env.example .env
make docker-build
make docker-run
```

### エンドポイント

- GET `/health` ヘルスチェック
- GET `/service-charges` 一覧
- GET `/service-charges/{id}` 取得
- POST `/service-charges` 作成（JSON: `{ "name": "X", "rate": 0.1 }`）
- DELETE `/service-charges/{id}` 削除


