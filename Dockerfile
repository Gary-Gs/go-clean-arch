# Builder
FROM golang:1.17-alpine3.16 as builder
RUN apk update && apk upgrade && \
    apk --update add git make
WORKDIR /app
COPY . .
RUN go build -o /app/engine /app/*.go
RUN go get github.com/swaggo/swag/cmd/swag@v1.8.3 && make swag

# Distribution
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata bash && \
    mkdir /app
WORKDIR /app
EXPOSE 9090
COPY --from=builder /app /app
CMD ["/app/engine"]
