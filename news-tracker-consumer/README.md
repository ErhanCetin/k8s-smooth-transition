## in your local to use stomp ( details : https://github.com/go-stomp/stomp )
  > go get github.com/go-stomp/stomp

## sample : https://github.com/thomasmodeneis/go-amq-poc
## REFERENCE : simplified implementation in my project is come from https://github.com/go-stomp/stomp/blob/master/examples/client_test/main.go


---
https://www.golangprograms.com/go-language/struct.html
https://vsupalov.com/go-json-omitempty/



docker build -t k8s-news-tracker-consumer:1.0.0 .
docker run -d --name="news-tracker-consumer"  k8s-news-tracker-consumer:1.0.0 .

#to exec docker for alpine
- docker exec -it news-tracker-consumer /bin/ash
