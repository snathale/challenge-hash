FROM golang:1.15-alpine as builder

RUN apk update && apk add ca-certificates

WORKDIR /opt/app
COPY ./ /opt/app/

RUN go mod tidy
RUN go build -o ./ ./

ENTRYPOINT ["/opt/app/calcultor", "run"]