FROM golang:latest as builder
RUN mkdir /build
ADD ping.go /build/
RUN go get -u github.com/tatsushid/go-fastping
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ping.

FROM golang:alpine
COPY --from=builder /build/ping /app/
WORKDIR /app
CMD ["./ping", "1.1.1.1"]