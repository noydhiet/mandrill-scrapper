FROM golang:1.23-alpine AS builder

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN apk add --no-cache ca-certificates git gcc musl-dev

RUN go mod download -x
# copy all directory
COPY . .
# build the binary
RUN GOOS=linux GOARCH=amd64 \
    go build -v -o /tmp/main ./main.go

FROM alpine:latest AS production

# copy from builder stage to production stage
COPY --from=builder /tmp/main /app/main

# Use root user
USER root

# running the application
CMD ["/app/main"]