FROM golang:1.12.6-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/movie_service
COPY . .

RUN apk add --no-cache git tzdata
#COPY /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN go get github.com/pilu/fresh

ENV API_TOKEN=""

CMD ["fresh"]
EXPOSE 9090