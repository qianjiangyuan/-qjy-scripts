#!/bin/bash
cd $1
echo "build image $2"
docker build . -t $2
#docker push $2
