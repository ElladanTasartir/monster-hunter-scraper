.PHONY: build
build:
	@go build -o ./monster-hunter-scraper ./cmd/monster-hunter-scraper-api/main.go

.PHONY: run
run:
	@air -c ./build/.air.toml

.PHONY: up
up:
	@docker-compose up -d
