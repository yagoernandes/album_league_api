db:
	docker-compose up -d

db-stop:
	docker-compose down

db-logs:
	docker-compose logs -f

server:
	make db
	go run ./...

.PHONY: db db-stop db-logs server