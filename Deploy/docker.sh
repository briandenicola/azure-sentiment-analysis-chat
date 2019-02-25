#!/bin/bash 

export RG=$1
export NAME=$2
export CONTAINER=$3

#key=$(az cognitiveservices account keys list -g $1 -n $2 --query 'key1' -o tsv)

docker run --rm -d -p 5000:5000 --name cogs --memory 4g --cpus 1 mcr.microsoft.com/azure-cognitive-services/sentiment \
    Eula=accept \
    Billing=https://southcentralus.api.cognitive.microsoft.com/ \
    ApiKey=c34790..
    
docker run --rm -d -p 8081:8081 --link cogs:cogs $CONTAINER
