FROM golang:1.13-alpine AS build

WORKDIR /app

RUN go build -o out/konference cmd/konference/main.go