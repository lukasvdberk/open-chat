FROM debian:stretch
FROM golang:latest


EXPOSE 4000

WORKDIR /app

COPY go.mod go.sum ./

CMD apt-get install entr

RUN go mod download

COPY . .