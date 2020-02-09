FROM golang:1.13.1

WORKDIR /go/auth_api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
