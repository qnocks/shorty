.PHONY:

build:
	go build -o ./.bin/app cmd/app/main.go

run: build
	./.bin/app

lint:
	golangci-lint run

swagger:
	swag init -g cmd/app/main.go