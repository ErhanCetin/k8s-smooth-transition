#### News Consumer Service
     - It is written by golang to fetch the news from activemq and write them down to mongo db.

#### Mongodb Service
     - It is used by new consumer to put news in it.

#### mongo-express 
   - It is used to check mongo db table via user interface. 
   - for local : http://localhost:8089/
## to push docker image to docker hub 
   - > docker build -t k8s-news-tracker-consumer .
   - > docker images
   - > docker tag 56c1dbad99a2 erhancetin/k8s-news-tracker-consumer:latest
   - > docker push erhancetin/k8s-news-tracker-consumer:latest
