# build
FROM golang:1.17.6-alpine as build

ENV ROOT /enva/migration
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ${ROOT}/
RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/main ${ROOT}/service/admin/internal/main.go

# prod
# FROM scratch as image # if scratch
FROM alpine:latest as image

ARG ENVIRONMENT
ARG POSTGRES_URL
ARG POSTGRES_MAX_CONNECTIONS
ARG POSTGRES_MAX_IDLE_CONNECTIONS

ENV ENVIRONMENT ${ENVIRONMENT}
ENV POSTGRES_URL ${POSTGRES_URL}
ENV POSTGRES_MAX_CONNECTIONS ${POSTGRES_MAX_CONNECTIONS}
ENV POSTGRES_MAX_IDLE_CONNECTIONS ${POSTGRES_MAX_IDLE_CONNECTIONS}

RUN apk add --no-cache ca-certificates && update-ca-certificates

ENV ROOT /enva/migration
WORKDIR ${ROOT}

# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ # if scratch
COPY --from=build ${ROOT}/main ${ROOT}

CMD ["/enva/migration/main", "migrate"]
