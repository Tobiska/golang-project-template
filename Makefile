.PHONY: build
build:
		go mod tidy
		go build -v -o golang-project-template ./cmd/main.go

.PHONY: gql-generate
gql-generate:
		go get -u github.com/99designs/gqlgen/cmd
		go run github.com/99designs/gqlgen generate

.DEFAULT_GOAL := build