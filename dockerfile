FROM golang:alpine AS builder
LABEL project-name = "braincome"
WORKDIR /app
COPY . .
RUN apk add --no-cache build-base
RUN go build -o braincome cmd/api/main.go
FROM alpine
WORKDIR /app 
COPY --from=builder /app .
CMD ["./forum"]
EXPOSE 8080