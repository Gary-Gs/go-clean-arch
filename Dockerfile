# Builder
FROM golang:1.17-alpine3.16 as builder
RUN apk update && apk upgrade && \
    apk --update add git make
WORKDIR /app
COPY . .
RUN go build -o /app/engine /app/*.go

# Distribution
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata bash && \
    mkdir /app
WORKDIR /app
EXPOSE 9090
COPY --from=builder /app /app
CMD ["/app/engine"]