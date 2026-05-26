FROM golang:1.26.3-alpine

WORKDIR /app

# Copy dependency files first
COPY . .

# Download the dependencies
RUN go mod download

# Build application
RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
