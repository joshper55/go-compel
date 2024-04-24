FROM golang:1.21.1-bullseye

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY ./src .
RUN go mod tidy