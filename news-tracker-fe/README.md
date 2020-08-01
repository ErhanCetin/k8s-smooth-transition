* docker build -t k8s-news-tracker-fe:1.0.0 .
* docker run -it --rm  -v ${PWD}:/app -v /app/node_modules -p 3001:3000 -e CHOKIDAR_USEPOLLING=true  k8s-news-tracker-fe:1.0.0
* docker-compose up --build