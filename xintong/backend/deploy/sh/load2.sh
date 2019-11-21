#!/bin/bash
echo $0
name=`docker load -i $2`
echo $name
OLD_IFS="$IFS"
IFS=":"
arr=($name)
IFS="$OLD_IFS"
str=${arr[0]}
echo $str
if [ "$str"  == "Loaded image" ]
then
   echo $3
   imageName=${arr[1]// /}
   harborImage=$3/$imageName:${arr[2]}
   docker tag $imageName:${arr[2]}  ${harborImage}
   docker push  ${harborImage}
fi
