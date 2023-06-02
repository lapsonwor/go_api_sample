container:
	docker run --name polka_game -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it polka_game createdb --username=root --owner=root game_backend_db

dropdb:
	docker exec -it polka_game dropdb game_backend_db

migrateGameUp:
	migrate -path datastore/migration/game -database "postgresql://root:password@localhost:5432/game_backend_db?sslmode=disable" -verbose up

migrateGameDown:
	migrate -path datastore/migration/game -database "postgresql://root:password@localhost:5432/game_backend_db?sslmode=disable" -verbose down

migrateBuildingUp:
	migrate -path datastore/migration/building -database "" -verbose up

migrateBuildingDown:
	migrate -path datastore/migration/building -database "" -verbose down

debugBuilding:
	API_ENV=DEBUG GIN_MODE=debug go run cmd/building/main.go --config cmd/building/config.json

debugGame:
	API_ENV=DEBUG GIN_MODE=debug go run cmd/game/main.go --config cmd/game/config.json

sqlc:
	sqlc generate

.PHONY: container createdb dropdb migrateGameUp migrateGameDown debugGame sqlc


