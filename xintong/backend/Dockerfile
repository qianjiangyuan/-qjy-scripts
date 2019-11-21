FROM golang
MAINTAINER like "like@oneflow.org"
WORKDIR /home
ADD ./ofmain /home/
RUN mkdir config
ENV  MONGO_SERVICE_HOST=10.5.24.113
ENV  MONGO_SERVICE_PORT=27017
ADD ./config /home/config
