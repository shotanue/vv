FROM golang:1.12.6-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/bff
COPY . .

RUN apk add --no-cache git tzdata

RUN go get github.com/pilu/fresh

CMD ["fresh"]
EXPOSE 9090