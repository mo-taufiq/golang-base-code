# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine
RUN apk add --no-cache bash

WORKDIR /app

COPY bin ./bin
COPY migrations ./migrations
COPY env-production.env ./
COPY go-app-binary ./

ENTRYPOINT ["/app/go-app-binary"]