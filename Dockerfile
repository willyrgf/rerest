FROM golang:1.13.4-alpine3.10
WORKDIR /rerest
RUN apk update
COPY . .

RUN go build

