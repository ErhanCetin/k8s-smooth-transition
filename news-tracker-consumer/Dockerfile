FROM golang:1.13.14-alpine
RUN mkdir -p /news-tracker-consumer
ADD . /news-tracker-consumer
WORKDIR /news-tracker-consumer
RUN apk update -qq && apk add git
RUN go get github.com/go-stomp/stomp
RUN go get gopkg.in/mgo.v2
RUN go build -o main .
ENTRYPOINT ["/news-tracker-consumer/main"]
