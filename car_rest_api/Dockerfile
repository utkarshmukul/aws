FROM golang:1.26.3

WORKDIR /app

# Copy dependency files first
COPY . .

# Sync the dependencies
RUN go mod tidy

# Build application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
