FROM golang:alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd ./cmd/email && go build -o app

RUN mv ./cmd/email/app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8082

CMD ["./app"]