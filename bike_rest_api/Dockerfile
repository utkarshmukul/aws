FROM golang:1.26.3

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy remaining source code
COPY . .

# Build application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
