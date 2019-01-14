FROM golang:1.7-alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/golnkr
WORKDIR $GOPATH/src/golnkr
# Fetch dependencies.
RUN go get -d -v github.com/go-redis/redis
# Build the binary.
RUN go build -o /go/bin/main
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]