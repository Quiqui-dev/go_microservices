
FROM alpine:3.17.0 as build_stage


WORKDIR /migrations

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

FROM alpine:latest

COPY --from=build_stage /bin/goose /bin/goose

