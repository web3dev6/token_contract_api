MAIN_DB_URL=postgresql://root:secret@localhost:5432/token_transaction?sslmode=disable
TEST_DB_URL=postgresql://root:secret@localhost:5431/token_transaction_test?sslmode=disable

postgres:
	docker run --name token-contract-api-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
	docker run --name token-contract-api-db-test -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it token-contract-api-db createdb --username=root --owner=root token_transaction
	docker exec -it token-contract-api-db-test createdb --username=root --owner=root token_transaction_test
dropdb:
	docker exec -it token-contract-api-db dropdb token_transaction
	docker exec -it token-contract-api-db-test dropdb token_transaction_test

migrateup:
	migrate -path db/migration -database "$(MAIN_DB_URL)" -verbose up
	migrate -path db/migration -database "$(TEST_DB_URL)" -verbose up
migrateup1 :
	migrate -path db/migration -database "$(MAIN_DB_URL)" -verbose up 1
	migrate -path db/migration -database "$(TEST_DB_URL)" -verbose up 1
migratedown:
	migrate -path db/migration -database "$(MAIN_DB_URL)" -verbose down
	migrate -path db/migration -database "$(TEST_DB_URL)" -verbose down
migratedown1:
	migrate -path db/migration -database "$(MAIN_DB_URL)" -verbose down 1
	migrate -path db/migration -database "$(TEST_DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -short -cover ./...

test-verbose:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/web3dev6/token_transaction/db/sqlc Store

dbdocs:
	dbdocs build doc/db.dbml

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml   

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown sqlc test server mock dbdocs dbschema