.PHONY: up down logs restart build clean ps exec-db exec-app test

# コンテナを起動
up:
	docker-compose up -d

# コンテナを停止
down:
	docker-compose down

# ログを表示
logs:
	docker-compose logs -f

# コンテナを再起動
restart:
	docker-compose restart

# コンテナを再ビルドして起動
build:
	docker-compose up -d --build

# コンテナとボリュームを削除
clean:
	docker-compose down -v

# コンテナの状態を確認
ps:
	docker-compose ps

# PostgreSQLに接続
exec-db:
	docker-compose exec postgres psql -U postgres -d testdb

# アプリケーションコンテナに接続
exec-app:
	docker-compose exec app sh

# APIのテスト
test:
	@echo "ヘルスチェック:"
	@curl -s http://localhost:8080/health | jq .
	@echo "\nユーザー一覧:"
	@curl -s http://localhost:8080/users | jq .

