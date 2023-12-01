FROM golang:1.21.4

WORKDIR /app

# Download Go modules
# add go.sum in future
COPY go.mod ./
RUN go mod download

COPY / ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Run
CMD ["/main"]