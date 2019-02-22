#!/bin/bash 

docker run --rm -d -p 5000:5000 --name cogs --memory 4g --cpus 1 mcr.microsoft.com/azure-cognitive-services/sentiment \
    Eula=accept \
    Billing=https://southcentralus.api.cognitive.microsoft.com/ 
    ApiKey=c34790d25d.......................
docker run --rm -d -p 8081:8081 --link cogs:cogs bjd145/chatws:0.2.5
