FROM golang:1.13.14-alpine
RUN mkdir -p /news-tracker-consumer
ADD . /news-tracker-consumer
WORKDIR /news-tracker-consumer

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /news-tracker-consumer
RUN chmod +x /news-tracker-consumer/wait
COPY start.sh /news-tracker-consumer
RUN chmod +x /news-tracker-consumer/start.sh

RUN apk update -qq && apk add git
RUN go get github.com/go-stomp/stomp
RUN go get gopkg.in/mgo.v2
RUN go build -o main .

ENTRYPOINT  ["/news-tracker-consumer/start.sh"]
