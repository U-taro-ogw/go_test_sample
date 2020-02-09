FROM golang:1.13.1

WORKDIR /go/src/auth_api
COPY . ./go/src/auth_api
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
