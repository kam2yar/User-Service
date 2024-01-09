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
recreate:
	docker-compose up -d --force-recreate