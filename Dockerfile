FROM golang:1.23 AS builder

WORKDIR /app

ENV GONOSUMDB=github.com/basputtipong/*

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o goapp main.go

FROM alpine:latest

RUN apk update && apk add --no-cache \
        ca-certificates \
        tzdata

ENV TZ=Asia/Bangkok

WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/goapp ./goapp

EXPOSE 1300

CMD ["./goapp"]