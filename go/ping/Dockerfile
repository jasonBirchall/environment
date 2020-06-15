FROM golang:latest

USER root
RUN apt-get update

ADD main.go .

RUN go get -u github.com/tatsushid/go-fastping
RUN go build main.go

CMD ["main", "1.1.1.1"]