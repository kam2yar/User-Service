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
pb:
	docker-compose exec app bash -c 'protoc -I ./api --go_out ./api/pb --go_opt paths=source_relative --go-grpc_out ./api/pb --go-grpc_opt paths=source_relative --grpc-gateway_out ./api/pb --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out ./api/swagger ./api/management.proto'