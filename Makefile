up:
	docker-compose up -d
down:
	docker-compose down
logs:
	docker-compose logs -f app
bash:
	docker-compose exec app bash
build:
	docker-compose build app
restart:
	docker-compose restart
recreate:
	docker-compose up -d --force-recreate
executable:
	docker-compose exec app go build -o ./bin/run ./cmd/run.go
	docker-compose exec app go build -o ./bin/migrate ./cmd/migrate.go