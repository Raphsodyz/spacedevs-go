FROM golang:1.26-alpine AS builder

RUN apk add --no-cache \
    tzdata \
    git \
    ca-certificates

ENV TZ=UTC

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o spacedevs-api ./cmd/api

FROM alpine:3.20 AS runtime

RUN apk add --no-cache \
    tzdata \
    ca-certificates

ENV TZ=UTC
RUN cp /usr/share/zoneinfo/UTC /etc/localtime \
    && echo "UTC" > /etc/timezone

RUN addgroup -S spacedevsgroup && adduser -S spacedevsuser -G spacedevsgroup

WORKDIR /app

COPY --from=builder /app/spacedevs-api ./spacedevs-api
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs

RUN chown -R spacedevsuser:spacedevsgroup /app
USER spacedevsuser

EXPOSE 5000

ENTRYPOINT ["./spacedevs-api"]