# dev
FROM golang:1.17.6-alpine as development

ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ${ROOT}
RUN go mod tidy

COPY . ${ROOT}/
EXPOSE 8080

CMD ["go", "run", "./service/api/cmd/server/"]

# build
FROM golang:1.17.6-alpine as builder

ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum .${ROOT}
RUN go mod tidy

COPY . .${ROOT}
RUN GOOS=linux go build -o ./bin/server/main ./service/api/cmd/server/main.go

# prod
FROM alpine:3.15 as prod
ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk add --update --no-cache ca-certificates
COPY --from=builder ${ROOT}/bin ${ROOT}/bin
CMD ["/enva/bin/server/main"]
