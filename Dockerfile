FROM golang:latest

RUN mkdir -p /go/src/github.com/44taka/golang-gin
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air

WORKDIR /go/src/github.com/44taka/golang-gin
