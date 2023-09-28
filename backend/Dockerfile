FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /rest ./cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /rest /rest
CMD [ "/rest" ]
