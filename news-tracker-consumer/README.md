#### News Consumer Service
 - It is implemented with golang to fetch the news from activemq and write them down to Mongo db.

- For pushing the docker image to the docker hub 
  > docker build -t k8s-news-tracker-consumer .
  > docker images
  > docker tag 56c1dbad99a2 erhancetin/k8s-news-tracker-consumer:latest
  > docker push erhancetin/k8s-news-tracker-consumer:latest


- For exec
  > kubectl exec -it newsconsumer-pod-name  /bin/sh   
