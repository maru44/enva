# build
FROM golang:1.17.6-alpine as build

ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ${ROOT}/
RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/bin/server/main ${ROOT}/service/api/cmd/server/main.go

# prod
FROM alpine:latest as image

# RUN apk add --update --no-cache ca-certificates

ENV ROOT=/enva
WORKDIR ${ROOT}

COPY --from=build ${ROOT}/bin/server/main ${ROOT}

EXPOSE 8080
CMD ["/enva/main"]

