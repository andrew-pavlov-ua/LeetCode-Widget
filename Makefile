
env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

restart:
	docker restart lc_redirect_app

logs:
	docker logs lc_redirect_app

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

app-build:
	docker exec lc_redirect_app go build -o /bin/lc-redirect-server ./cmd/v1/main.go

app-start:
	docker exec lc_redirect_app lc-redirect-server

app-stop:
	docker exec lc_redirect_app pkill lc-redirect-server || echo "lc_redirect-server already stopped"

app-restart: app-build app-stop app-start

generate-sqlc:
	sqlc generate

go-mod-update:
	go mod tidy
	go mod vendor

local-go-app-run:
	POSTGRES_DSN=$(POSTGRES_DSN) PORT=:8080 go run ./cmd/v1/main.go