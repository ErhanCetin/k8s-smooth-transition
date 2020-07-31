import datetime
import json
import os
from time import sleep

import requests
import stomp

# reference : https://newsapi.org/docs/client-libraries/python

fromDate = datetime.date.today() - datetime.timedelta(2)
toDate = datetime.date.today()
# connection for newsapi.org
news_url_for_everthing = 'https://newsapi.org/v2/everything'

# parameter reference https://newsapi.org/docs/endpoints/everything
payload_for_everything = {'q': 'technology', 'language': 'en', 'sortBy': 'publishedAt', 'from': fromDate, 'to': toDate}

# check if null then use sample authorization
authorization = os.environ.get('NEWSAPI_AUTHORIZATION')
if authorization is None:
    authorization = 'db87162d00af4d1bb4c8031ad1cf22f5'

headers = {'Authorization': authorization}

response = requests.get(url=news_url_for_everthing, headers=headers, params=payload_for_everything)
pretty_json_output = json.dumps(response.json(), indent=4)
print(pretty_json_output)

response_json_string = json.dumps(response.json())
response_dict = json.loads(response_json_string)
print(response_dict)

sleep(5)

## put news data to activemq
activemqHostName = os.environ.get('ACTIVEMQ_HOST')
activemqPort = os.environ.get('ACTIVEMQ_PORT')
activemqQueueName = os.environ.get('ACTIVEMQ_QUEUE_NAME')
activemqUser = os.environ.get('ACTIVEMQ_USER_LOGIN')
activemqPassword = os.environ.get('ACTIVEMQ_USER_PASSWORD')
activemqDestination = "/queue/" + activemqQueueName
conn = stomp.Connection([(activemqHostName, activemqPort)])
conn.connect(activemqUser, activemqPassword, wait=False)
conn.send(body=response_json_string, destination=activemqDestination)
conn.disconnect()
