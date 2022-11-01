.PHONY:

build:
	go build -o ./.bin/app cmd/app/main.go

run: build
	./.bin/app

test:
	go test ./cmd/... ./internal/... -race -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"

lint:
	golangci-lint run

swagger:
	swag init -g cmd/app/main.go

gen:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go

docker-image:
	docker build -t shorty .
