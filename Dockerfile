FROM golang:1.13 as build
WORKDIR /go/src/github.com/willyrgf/rerest
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build . \
    && mv rerest /go/bin/rerest

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/willyrgf/rerest/config.toml /app/
COPY --from=build /go/bin/rerest /app/
CMD ["/app/rerest"]

