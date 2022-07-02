# Build
FROM golang:1.18-alpine as builder

RUN apk update && apk add --no-cache git

COPY . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -ldflags '-s -w -extldflags "-static"' -o ./app cmd/rest/main.go

# Run
FROM gcr.io/distroless/base-debian11

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]