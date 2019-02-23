# Overview 

This is a silly little chat application that demostrates Azure Cognitive Services 
It is written in Go and uses Web Sockets
If a 

# Azure 
This code uses Azure Cognitive Services for sentiment analysis.  
We will use their [container option](https://docs.microsoft.com/en-us/azure/cognitive-services/cognitive-services-container-support)
This requires Azure Cognitive Services created in Azure for billing purposes but all analysis will be done locally

## Azure Setup
1. Login to Azure Cloud Shell
2. az cognitiveservices account create -g DevSub01_Sentiment_RG -n bjdai002 --sku S0 --location southcentralus --kind CognitiveServices
3. az cognitiveservices account keys list -g DevSub01_Sentiment_RG -n bjdai002
   * Copy key1 which will be used in the deploy.yaml file 

# Build
1. Set your GOPATH 
   * Yes I am not using modules yet for this project. Maybe in a bit
2. go get github.com/gorilla/websocket
3. go get github.com/rs/cors
4. go get gopkg.in/resty.v1
5. cd src
6. docker build -t chatws .

# Deploy 
1. Update the ApiKey in the deploy.yaml with the Key you obtained from Step 3 in the Azure Setup. 
2. kubectl apply -f ./deploy.yaml

