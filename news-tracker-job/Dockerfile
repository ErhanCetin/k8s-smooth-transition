FROM python:3.7-alpine
COPY ./org/news/tracker /app
WORKDIR /app
RUN pip3 install -r requirements.txt
# Add docker-compose-wait tool
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /app
RUN chmod +x /app/wait
COPY start.sh /app
RUN chmod +x /app/start.sh
ENTRYPOINT  ["/app/start.sh"]
