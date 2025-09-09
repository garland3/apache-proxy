# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dos2unix to fix line endings
RUN apk add --no-cache dos2unix

# Copy go mod files (go.sum may not exist initially)
COPY go.mod ./
COPY go.sum* ./

# Fix line endings for Windows files
RUN dos2unix go.mod
RUN test -f go.sum && dos2unix go.sum || true

# Download dependencies
RUN go mod download && go mod verify

# Copy source code and templates (but don't overwrite go.mod/go.sum)
COPY main.go ./
COPY templates/ ./templates/

# Fix line endings for all Go files
RUN find . -name "*.go" -exec dos2unix {} \;

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary and templates from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./main"]
