FROM golang:1.14 as builder

WORKDIR /go/src/lc
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/lc


FROM alpine

WORKDIR /bin/lc

COPY --from=builder /bin/lc .
COPY secrets.json .

CMD ["./lc"]
