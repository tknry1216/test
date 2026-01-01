# Golang PostgreSQL Docker Compose

Go 1.25 と PostgreSQL を使用したサンプルアプリケーションです。

## 必要要件

- Docker
- Docker Compose

## セットアップ

1. コンテナを起動:
```bash
docker-compose up -d
```

2. ログを確認:
```bash
docker-compose logs -f
```

3. コンテナを停止:
```bash
docker-compose down
```

4. データベースも含めて完全に削除:
```bash
docker-compose down -v
```

## API エンドポイント

アプリケーションは `http://localhost:8080` で起動します。

### ホーム
```bash
curl http://localhost:8080/
```

### ヘルスチェック
```bash
curl http://localhost:8080/health
```

### ユーザー一覧取得
```bash
curl http://localhost:8080/users
```

### ユーザー作成
```bash
curl -X POST http://localhost:8080/users/create \
  -H "Content-Type: application/json" \
  -d '{"name":"山田太郎","email":"yamada@example.com"}'
```

## データベース接続

PostgreSQL に直接接続する場合:

```bash
docker-compose exec postgres psql -U postgres -d testdb
```

## 環境変数

`.env` ファイルで以下の設定を変更できます:

- `POSTGRES_USER`: データベースユーザー名
- `POSTGRES_PASSWORD`: データベースパスワード
- `POSTGRES_DB`: データベース名
- `POSTGRES_PORT`: PostgreSQLのポート
- `APP_PORT`: アプリケーションのポート

## 開発

ソースコードを変更した場合、コンテナを再起動してください:

```bash
docker-compose restart app
```

または、再ビルドする場合:

```bash
docker-compose up -d --build
```

