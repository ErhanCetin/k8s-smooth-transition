## NEW TRACKER PRODUCER SERVICE 
 - This simple project is created to use in k8s article. 

Components 
---
#### News Tracker Producer 
     - It is written by Python to fetch the tech news from newsapi and write them down to Activemq queue.
 
###### Build image
 > docker build -t k8s-news-tracker-job:1.0.0 .
 
###### Run standalone container in local :
> docker run --name="news-python" -e NEWSAPI_AUTHORIZATION='<your-authorization-key-(OPTIONAL)>' -e ACTIVEMQ_HOST=activemq -e ACTIVEMQ_PORT=61613  k8s-news-tracker-job:1.0.0

#### ActiveMQ Service
- This service is used by news tracker producer and news consumer services. 

###### Test standalone activemq.
> docker run --name='activemq' -d   -p  61617:61616 -p 61613:61613 -p 8162:8161 -e 'ACTIVEMQ_CONFIG_MINMEMORY=512' -e 'ACTIVEMQ_CONFIG_MAXMEMORY=1024'\ -P webcenter/activemq:latest
> - in local : http://localhost:8162/ad

## Run docker-compose
- The newsJob service needs to activemq service to write news . You can look at docker-compose.yml configuration to up the service. 
- To run all service in compose yml file
   > docker-compose up -d   
                                           >
## to push docker image to docker hub 
> docker build -t k8s-news-tracker-job .
> docker images
> docker tag bb38976d03cf erhancetin/k8s-news-tracker-job:latest
> docker push erhancetin/k8s-news-tracker-job:latest

## to push activemq image to your repository in order to use in local minikube.
> docker images | grep "activemq"
> docker tag 3af156432993 erhancetin/activemq
> docker push  erhancetin/activemq 

## to exec 
> kubectl exec -it <newsproducer-pod-name>  /bin/sh