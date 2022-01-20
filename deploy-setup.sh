#!/bin/sh

# export GIN_MODE=release
export GOOS=linux
go build -o go-app-binary

docker container stop go-app-1
docker container rm go-app-1
docker rmi go-app:latest

docker build --tag go-app .

docker compose up -d