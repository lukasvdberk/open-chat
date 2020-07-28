FROM golang:latest

EXPOSE 4000

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .