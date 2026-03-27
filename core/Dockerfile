FROM golang:alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd ./cmd/core && go build -o app

RUN mv ./cmd/core/app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY --from=builder /app/migrations migrations

EXPOSE 8081

CMD ["./app"]