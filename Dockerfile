FROM golang:1.12.3-alpine

WORKDIR /go/src/github.com/kind84/cacoo
COPY . .

RUN apk update && apk add git gcc libc-dev
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cacoo"]