postgres:
	docker run --name tester -p 5432:5432 POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2

createdb:
	docker exec -it tester createdb --username=root --owner=root testerdb

dropdb:
	docker exec -it tester dropdb testerdb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testerdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testerdb?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -destination db/mock/store.go github.com/tester/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test