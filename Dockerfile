FROM golang:1.14.3-alpine3.11 AS builder

LABEL maintainer="ondengew@gmail.com"

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git mercurial 

# Move to working directory /app
RUN mkdir /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build ./cmd/atsvr && go build ./cmd/atsvr-cli

# Build a small image
FROM scratch

COPY --from=builder /app /

EXPOSE 8000 8000

# Command to run
CMD ["./atsvr", "-host", "docker"]


