* just sample about how to deploye react via kubernetes : 
   https://dev.to/rieckpil/deploy-a-react-application-to-kubernetes-in-5-easy-steps-516j
   https://blog.codecentric.de/en/2018/12/react-application-container-environment-aware-kubernetes-deployment/
   https://koala42.com/create-a-react-app-in-kubernetes/
   https://www.bogotobogo.com/DevOps/Docker/Docker-React-App.php
   https://dzone.com/articles/how-to-dockerize-reactjs-app
   https://dev.to/igmrrf/docker-react-exited-with-code-0-398n


* Reference for injecting env varible during the run time :https://www.freecodecamp.org/news/how-to-implement-runtime-environment-variables-with-create-react-app-docker-and-nginx-7f9d42a91d70/

* docker build -t k8s-news-tracker-fe:1.0.0 .
## If you are using "react-scripts": "3.4.1" , probably , you will get this error " "404s will fallback to react in docker container" .To perevent this you can downgrade "react-scripts": "3.4.0" like I did.If you keep 3.4.1 then : 
## -->   use "-it" for docker run command
## -->   put "stdin_open: true" to compose yaml. 
* docker run -it --rm  -v ${PWD}:/app -v /app/node_modules -p 3001:3000 -e CHOKIDAR_USEPOLLING=true  k8s-news-tracker-fe:1.0.0
* docker-compose up --build

 * docker run --rm --name  trackfe -e NEWS_API_HOST_NAME=127.0.0.1 -e NEWS_API_HOST_PORT=63864 -p 3000:3000  erhancetin/k8s-news-tracker-fe

* Usefull link 
-- https://mherman.org/blog/dockerizing-a-react-app/
-- https://dev.to/bmvantunes/next-js-environment-variables-and-runtime-config-50oa

## to push docker image to docker hub 
 > docker build -t erhancetin/k8s-news-tracker-fe .
 > docker images
 > docker tag 555f9b882d57 erhancetin/k8s-news-tracker-fe:latest
 > docker push erhancetin/k8s-news-tracker-fe:latest

 ## kubectl exec -it <react-pod-name> sh