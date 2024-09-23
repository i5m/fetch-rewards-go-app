# syntax=docker/dockerfile:1

FROM golang:1.21-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy all the source code files.
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /reward-app

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

# Run the app
CMD ["/reward-app"]
