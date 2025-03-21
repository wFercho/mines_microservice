FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o mines_microservice main.go

FROM debian:bullseye

WORKDIR /root/

COPY --from=builder /app/mines_microservice .

CMD ["./mines_microservice"]
