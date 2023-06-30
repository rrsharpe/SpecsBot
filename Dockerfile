# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app/go-ssd-bot

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go ./
COPY ssd ./ssd
COPY dispatcher ./dispatcher

RUN go build -o . ./...

CMD [ "./Go-SSD-Bot" ]
