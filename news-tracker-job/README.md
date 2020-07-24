# NEW TRACKER JOB 
- This simple project is created to use in k8s article. 

### There are two services in here : 
     - newsjob service : It is written by Python to fetch the tech news from newsapi and write them down to Activemq que.
     - activemq service : It is used for a que service.  
  
### Some Instructor for local : 
#### Build image
 > docker build -t k8s-news-tracker-job:1.0.0 .
 
### Run standalone container in local
> docker run --name="news-python" -e NEWSAPI-AUTHORIZATION='<your-authorization-key-(OPTIONAL)>' -e ACTIVEMQ-HOST=activemq -e ACTIVEMQ-PORT=61613  k8s-news-tracker-job:1.0.0

## Test standalone activemq.
##### browse in local : http://localhost:8162/admin
 > docker run --name='activemq' -d   -p  61617:61616 -p 61613:61613 -p 8162:8161 -e 'ACTIVEMQ_CONFIG_MINMEMORY=512' -e 'ACTIVEMQ_CONFIG_MAXMEMORY=1024'\ -P webcenter/activemq:latest

## Run docker-compose
- The newsJob service needs to activemq service to write news . You can look at docker-compose.yml configuration to up the service. 
- To run all service in compose yml file
   > docker-compose up -d   
