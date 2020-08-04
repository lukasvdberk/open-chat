FROM golang:latest

EXPOSE 4000

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

COPY . .

# starth.sh actually starts the server this is just for setup.