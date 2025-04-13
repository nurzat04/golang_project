FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy

RUN go build -o bin/shorten ./cmd/shorten

RUN go build -o bin/redirect ./cmd/redirect

RUN go build -o bin/stats ./cmd/stats

# 运行镜像阶段
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin ./bin
COPY .env .

CMD ["./bin/shorten"]
