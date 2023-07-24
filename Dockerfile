# Stage 1
FROM golang:1.20.3-alpine3.16 AS builder
WORKDIR /build
COPY . .
RUN go build -o main .

# Step 2
FROM alpine:3.18
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]