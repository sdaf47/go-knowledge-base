FROM golang:1.11.2-alpine3.8 AS builder

WORKDIR /go/src/main
COPY . .

RUN apk update && apk add git
RUN go install -v ./...


FROM alpine

WORKDIR /go/src/main

COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/src/main /go/src/main

RUN apk update && apk add openssl

CMD /go/src/main/new_cert.sh ; /go/bin/main
