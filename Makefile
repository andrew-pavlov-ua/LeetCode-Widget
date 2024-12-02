POSTGRES_DSN="postgresql://lc_user:lc_pass@lc_badge_postgres:5432/lc_badge?sslmode=disable"


env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

restart:
	docker restart lc_badge_app

logs:
	docker logs lc_badge_app

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

app-build:
	docker exec lc_badge_app go build -o /bin/lc-redirect-server ./cmd/v1/main.go

app-start:
	docker exec lc_badge_app lc-redirect-server

app-stop:
	docker exec lc_badge_app pkill lc-redirect-server || echo "lc_redirect-server already stopped"

app-restart: app-build app-stop app-start

migrate-pgsql-goose-install:
	docker exec lc_badge_app go install github.com/pressly/goose/v3/cmd/goose@latest

#	Migrating postgresql db with goose
migrate-pgsql-up: migrate-pgsql-goose-install
	docker exec -e GOOSE_DRIVER=postgres \
                -e GOOSE_DBSTRING=$(POSTGRES_DSN) \
                lc_badge_app goose -dir internal/storage/migrations -table schema_migrations up

migrate-pgsql-create:
	# mkdir -p ./internal/storage/migrations
	$(eval NAME ?= todo)
	goose -dir internal/storage/migrations postgres $(POSTGRES_DSN) create init sql

migrate-pgsql-down:
	docker exec lc_badge_app goose -dir ./internal/storage/migrations -table schema_migrations postgres down

generate-sqlc:
	sqlc generate

go-mod-update:
	go mod tidy
	go mod vendor

qtc-gen:
	qtc -dir=internal/templates/v1

run-tests:
	docker exec lc_badge_app go test -v ./...

connect-psql-cont:
	docker exec -it lc_badge_postgres bash