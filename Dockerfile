# STAGE: BUILD
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download

RUN CGO_ENABLED=0 go build -o daas

# STAGE: TARGET

FROM alpine:latest
RUN apk add ffmpeg

RUN addgroup -S user && adduser -S user -G user

USER user

WORKDIR /app
COPY --from=builder /app/daas /app/daas

ENTRYPOINT ["/app/daas"]
