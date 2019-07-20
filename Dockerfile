FROM golang:alpine

# Install git. (required for fetching dependencies)
RUN apk update && apk add --no-cache git
# Create appuser
RUN adduser -D -g '' appuser
USER appuser
WORKDIR /go/go-project
COPY . .
# Fetch dependencies
# Using go get.
RUN go get -d -v
# Using go mod.
RUN go mod download && go mod verify
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" daemon.go
EXPOSE 8080
ENTRYPOINT ["./daemon"]
CMD ["console"]