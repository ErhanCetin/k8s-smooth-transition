## News Tracker Application 
- This simple project is created to use in k8s article. 

 Components 
---
#### News Tracker Producer 
     - It is written by Python to fetch the tech news from newsapi and write them down to Activemq queue.
 
###### Build image
 > docker build -t k8s-news-tracker-job:1.0.0 .
 
###### Run standalone container in local :
> docker run --name="news-python" -e NEWSAPI-AUTHORIZATION='<your-authorization-key-(OPTIONAL)>' -e ACTIVEMQ-HOST=activemq -e ACTIVEMQ-PORT=61613  k8s-news-tracker-job:1.0.0

#### ActiveMQ Service
- This service is used by news tracker producer and news consumer services. 

###### Test standalone activemq.
> docker run --name='activemq' -d   -p  61617:61616 -p 61613:61613 -p 8162:8161 -e 'ACTIVEMQ_CONFIG_MINMEMORY=512' -e 'ACTIVEMQ_CONFIG_MAXMEMORY=1024'\ -P webcenter/activemq:latest
> - in local : http://localhost:8162/admin

#### News Consumer Service
     - It is written by golang to fetch the news from activemq and write them down to mongo db.

#### Mongodb Service
     - It is used by new consumer to put news in it.

#### mongo-express 
   - It is used to check mongo db table via user interface. 
   - for local : http://localhost:8089/
   
### Steps between services
 * Fetch news from newsapi and write them to activemq ( by news tracker producer service)  
 * Consume news from active mq and put them into mongo db
 
 
 
