# prod
FROM alpine:3.15
ENV ROOT=/enva
WORKDIR ${ROOT}

RUN apk add --update --no-cache ca-certificates
# COPY --from=builder ${ROOT}/bin ${ROOT}/bin/
CMD ["/enva/bin/server/main"]
