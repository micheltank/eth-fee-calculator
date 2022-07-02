deps:
	go mod tidy

run:
	set -o allexport && source env.example && go run cmd/main.go

first_run: deps run

lint:
	golangci-lint run

lint_fix:
	golangci-lint run --fix

docker_build:
	docker build -t micheltank/eth-fee-calculator .

test:
	go test -v -cover -race ./...

swagger:
	swag init -g cmd/rest/main.go

docker_compose:
	docker compose up -d

all: deps lint docker_build test swagger docker_compose