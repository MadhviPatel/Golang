#!/bin/bash

# Remove images in case we already have them to start clean
echo "Removing old containers and images in case they exist"
docker rm -f mongo
docker rmi -f mongo:4.2
docker rm -f graylog
docker rmi -f graylog/graylog:4.0
docker rm -f elasticsearch
docker rmi -f docker.elastic.co/elasticsearch/elasticsearch-oss:7.10.0

# Deletes the elastic search data folder if it exists and creates a new one
rm -rf ./data
mkdir ./data

# Download and run the containers
# https://docs.graylog.org/en/4.0/pages/installation/docker.html
echo "Starting setup"
docker run --name mongo -d mongo:4.2
docker run --name elasticsearch \
  -v $(pwd)/data:/usr/share/elasticsearch/data \
  -p 9200:9200 \
  -e "http.host=0.0.0.0" \
  -e "discovery.type=single-node" \
  -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
  -d docker.elastic.co/elasticsearch/elasticsearch-oss:7.10.0
docker run --name graylog --link mongo --link elasticsearch \
  -p 9000:9000 -p 12201:12201 -p 1514:1514 -p 5555:5555 \
  -e GRAYLOG_HTTP_EXTERNAL_URI="http://127.0.0.1:9000/" \
  -d graylog/graylog:4.0
