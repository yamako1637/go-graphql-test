FROM golang:tip-alpine3.22 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -ldflags '-s -w' -o main .

FROM alpine:3.22 AS runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]