FROM golang:latest as builder
RUN mkdir /build
ADD main.go /build/
RUN go get -u github.com/tatsushid/go-fastping
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM golang:alpine
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main", "1.1.1.1"]