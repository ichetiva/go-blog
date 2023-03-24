FROM golang:1.20

WORKDIR /code

COPY go.* /code/
COPY ./cmd /code/
COPY ./internal /code/
COPY ./pkg /code/
COPY ./config /code/
