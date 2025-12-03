.PHONY: run-dev
run-dev:
	docker compose -f compose.dev.yaml up -d

.PHONY: down-dev
down-dev:
	docker compose -f compose.dev.yaml down

.PHONY: install
install:
	go mod download && npm install

.PHONY: database-seed
database-seed:
	go tool goose -dir "./database/seeds" -no-versioning down
	go tool goose -dir "./database/seeds" -no-versioning up

.PHONY: tailwind-watch
tailwind-watch:
	npx @tailwindcss/cli -i ./website/assets/css/main.css -o ./website/dist/css/main.css --watch

.PHONY: tsc-watch
tsc-watch:
	npx tsc --watch
