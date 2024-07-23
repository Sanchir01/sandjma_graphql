FROM golang:1.22.5-alpine3.20

LABEL authors="sanchir"

COPY ./ ./

RUN go mod download

RUN sudo apt-get update && \
    sudo apt-get install -y make

RUN go build -o ./.bin/main ./cmd/main/main.go
CMD [""]
