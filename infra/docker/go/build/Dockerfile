# build
FROM golang:1.17.6-alpine

ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ${ROOT}/
RUN GOOS=linux go build -o ./bin/server/main ./service/api/cmd/server/main.go
