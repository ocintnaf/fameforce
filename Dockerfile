FROM golang:1.20.4-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o fameforce ./cmd/app

WORKDIR /dist

RUN cp /build/fameforce .

FROM alpine:3.14.2

RUN apk add --no-cache ca-certificates

COPY --from=builder /dist/fameforce /fameforce
COPY --from=builder /build/config /config

EXPOSE 8080

CMD ["/fameforce"]
