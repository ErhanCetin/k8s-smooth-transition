## News Tracker APPLICATION 
- This project is created to use in k8s article. 

 Components 
---
#### News Tracker Producer 
   - It is implemented with Python to fetch the tech news from newsapi and write them down to Activemq queue.
 
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
   - It is written with golang to fetch the news from activemq and write them down to mongo db.

#### Mongodb Service
   - It is used by new consumer to put news in it.

#### Mongo-Express 
   - It is used to check mongo db table via user interface. 
   - for local : http://localhost:8089/

#### NewsApi Service 
   - It is implemented with Spring Boot and SpringData MongoTemplate and is used to serve the news data to 3rd application and  
   - for local : http://localhost:8085/news/getAll

#### Frontend Service 
   - It is implemented with React  
   - for local : http://localhost:3001
         
### Steps between services
 * Fetch news from newsapi and write them to activemq ( by news tracker producer service)  
 * Consumes news from active mq and put them into mongo db ( by consumer service)
 * Fetch news data from Mongo Db and expose them by newsapi service for 3rd applications
 * Frontend service will show the news in formatted web page.   
 
 
 
