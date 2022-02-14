FROM golang:latest
RUN mkdir /golang-project-template
WORKDIR /golang-project-template
ADD . /golang-project-template
RUN go mod tidy | go mod download
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="make build;" --command=./app