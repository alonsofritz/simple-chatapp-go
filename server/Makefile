postgres_init:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root chatapp-go

dropdb:
	docker exec -it postgres15 dropdb chatapp-go

migrate_up:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/chatapp-go?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/chatapp-go?sslmode=disable" -verbose down

.PHONY: postgres_init postgres createdb dropdb migrate_up migrate_down