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
	docker-compose exec app go build -race -o ./bin/run ./cmd/run.go
	docker-compose exec app go build -race -o ./bin/migrate ./cmd/migrate.go
pb:
	docker-compose exec app bash -c 'protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --openapiv2_out . ./api/management.proto'