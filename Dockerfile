FROM golang:1.9

COPY .  /go/src/github.com/pcieslar/goforge

WORKDIR /go/src/github.com/pcieslar/goforge

RUN go get -d -v
RUN go install -v

CMD ["goforge"]
EXPOSE 80