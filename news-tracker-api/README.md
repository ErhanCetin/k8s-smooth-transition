#### News Tracker API 
 - It is implemented with spring boot to fetch the tech news from Mongo db to serve the FE.

- Build a image
  > docker build -t k8s-news-tracker-api .
                                       
- To push docker image to docker hub 
  > docker build -t k8s-news-tracker-api .
  > docker images
  > docker tag ced9054b2925 erhancetin/k8s-news-tracker-api:latest
  > docker push erhancetin/k8s-news-tracker-api:latestt