# dev
FROM golang:1.17.6-alpine as development

ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ${ROOT}/
EXPOSE 8080

CMD ["go", "run", "./service/api/cmd/server/"]
