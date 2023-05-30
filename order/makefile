run: ## runs app localy
	@go run cmd/app/main.go
.PHONY: run

up: ## testing infra
	@docker-compose -f deployments/docker-compose.yaml up -d
.PHONY: up

commit: ## build and push container to docker hub| provide tag=(...)
	@docker build -t dmitryavdonin/profile:latest . && docker push dmitryavdonin/profile:latest
.PHONY: commit

migrations-up:
	@migrate -source file://migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up
.PHONY: migrations-up

gen-migrations-file: ##creates new migrations files
	@migrate create -ext sql -dir migrations -seq $(name)
.PHONY: gen-migrations-file
