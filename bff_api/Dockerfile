FROM golang:1.13.1

RUN apt-get update -qq && apk add --no-cache alpine-sdk git

ENV GO111MODULE=on

WORKDIR /go/src/github.com/U-taro-ogw/go_test_sample

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
