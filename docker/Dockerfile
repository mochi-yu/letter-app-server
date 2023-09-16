FROM golang:latest

ENV APP_ROOT=/app

RUN apt update

WORKDIR ${APP_ROOT}

RUN go install github.com/cosmtrek/air@latest
