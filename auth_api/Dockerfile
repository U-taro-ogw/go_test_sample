FROM golang:1.13.1

RUN apt-get update -qq

ENV GO111MODULE=on

WORKDIR /go/src/auth_api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/pilu/fresh
CMD ["fresh"]
