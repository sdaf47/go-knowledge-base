FROM golang:alpine

RUN apk update && apk add git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

CMD ["app"]
