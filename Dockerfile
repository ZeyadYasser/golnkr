FROM golang:1.7-alpine
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/github.com/ZeyadYasser/golnkr
WORKDIR $GOPATH/src/github.com/ZeyadYasser/golnkr
# Fetch dependencies.
RUN go get -d -v github.com/go-redis/redis
# Install
RUN go install -v .
# Run the binary.
ENTRYPOINT ["golnkr"]