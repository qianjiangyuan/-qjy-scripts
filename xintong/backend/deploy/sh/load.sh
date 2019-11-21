#!/bin/bash
name=`docker load -i $1`
echo "name="$name
OLD_IFS="$IFS"
IFS=":"
arr=($name)
IFS="$OLD_IFS"
str=${arr[0]}
if [ "$str"  == "Loaded image" ]
then
   imageName=${arr[1]// /}
   harborImage=$2/$3:$4
   docker tag $imageName:${arr[2]}  ${harborImage}
   docker push  ${harborImage}
fi
