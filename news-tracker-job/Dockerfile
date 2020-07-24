FROM python:3.7-alpine
COPY ./org/news/tracker /app
WORKDIR /app
RUN pip3 install -r requirements.txt
CMD ["python", "NewsApi.py"]
