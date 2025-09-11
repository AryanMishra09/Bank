#Build stage
FROM golang:1.25-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go 

##Run stage
FROM alpine:3.22

EXPOSE 8080

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
