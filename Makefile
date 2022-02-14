.PHONY: build
build:
		go mod tidy
		go build -tags migrate -o app ./cmd

.PHONY: gql-generate
gql-generate:
		go get -u github.com/99designs/gqlgen/cmd
		go run github.com/99designs/gqlgen generate

.DEFAULT_GOAL := build