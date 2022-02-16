FROM golang:latest

RUN mkdir -p /go/src/github.com/44taka/golang-gin
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air
RUN go get -u gorm.io/driver/mysql
RUN go get -u gorm.io/gorm
RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/google/uuid \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls \
  github.com/dgrijalva/jwt-go \ 
  github.com/appleboy/gin-jwt/v2 \
  github.com/stretchr/testify
RUN go install github.com/golang/mock/mockgen@latest

WORKDIR /go/src/github.com/44taka/golang-gin
