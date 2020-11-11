FROM maven:3-jdk-8-alpine
ENV MONGODB-HOST=localhost
ENV MONGODB-NAME=news-db
ENV MONGODB-PORT=27017
WORKDIR /usr/src/java-code
COPY . /usr/src/java-code/
RUN mvn package
WORKDIR /usr/src/java-app
RUN cp /usr/src/java-code/target/*.jar ./app.jar
CMD ["java","-jar", "app.jar"]
