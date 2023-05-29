# Modules caching
FROM golang:buster AS modules

WORKDIR /modules

COPY go.mod go.sum ./

RUN go mod download

# Builder
FROM golang:buster AS builder

COPY --from=modules /go/pkg /go/pkg

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -tags migrate -o /bin/app ./cmd/app

# App
FROM scratch

COPY --from=builder /app/config /config
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app

EXPOSE 8080

CMD ["/app"]