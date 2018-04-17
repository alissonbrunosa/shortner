FROM golang:latest

COPY . $GOPATH/src/github.com/alissonbrunosa/shortner
WORKDIR $GOPATH/src/github.com/alissonbrunosa/shortner

RUN go get github.com/go-redis/redis

EXPOSE 8080
