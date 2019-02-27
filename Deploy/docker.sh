#!/bin/bash 

export CONTAINER=$1
export RG=$2
export NAME=$3


#key=$(az cognitiveservices account keys list -g $1 -n $2 --query 'key1' -o tsv)

docker run --rm -d -p 5000:5000 --name cogs --memory 4g --cpus 1 mcr.microsoft.com/azure-cognitive-services/sentiment \
    Eula=accept \
    Billing=https://southcentralus.api.cognitive.microsoft.com/ \
    ApiKey=c34790d25d7641898cc6e66c2869cff4
    
docker run --rm -d -p 8081:8081 --link cogs:cogs $CONTAINER
