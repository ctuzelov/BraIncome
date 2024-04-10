# First stage: Build the Go application
FROM golang:alpine AS builder
LABEL project-name="braincome"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache build-base
RUN go build -o braincome cmd/api/main.go

# Second stage: Create a minimal production image
FROM alpine
WORKDIR /app
COPY --from=builder /app/braincome .
CMD ["./braincome"]
EXPOSE 8080
