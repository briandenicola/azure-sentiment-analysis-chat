# Overview 

This is a demo application written in Go to showcase Azure Cognitive Services running in Docker Containers. 
It is a websocket chat web application.  It parses all messages using Azure Cognitive Services Sentimnet Analysis.
If the sentiment is below a set threshold (ie. the message is too mean) then the AI Chat Moderator will broadcast to everyone to be nice. 

# Azure 
We will use the Azure Cognitive Services [container deployment](https://docs.microsoft.com/en-us/azure/cognitive-services/cognitive-services-container-support)
This means that all analysis is done locally and not in Azure. 
    * This still requires a Cognitive Services deployment in Azure for billing purposes 

## Azure Setup
1. Login to Azure Cloud Shell
2. cogs='bjdai001'
2. az cognitiveservices account create -g DevSub01_Sentiment_RG -n $cogs --sku S0 --location southcentralus --kind CognitiveServices
3. az cognitiveservices account keys list -g DevSub01_Sentiment_RG -n $cogs --query 'key1' -o tsv
   * Copy the key which will be used in the deploy.yaml file 

# Local Build
1. Set your GOPATH 
   * Not using modules yet for this project. Now that Go 1.12 is out, I might revisit this.
2. go get github.com/gorilla/websocket
3. go get github.com/gin-contrib/cors
4. go get github.com/gin-gonic/gin
5. go get github.com/gin-gonic/contrib/static
6. go get gopkg.in/resty.v1
7. cd src
8. docker build -t chatws .

# Local Deploy
1. az login
2. Deploy/docker.sh <ContainerName> <Resource Group for Sentiment Analysis> <Sentiment Analysis API Name>
3. Launch a Browser go to http://localhost:8081

# Deploy to Kubernetes Manually 
1. Update the Container Name/Versionin the deploy.yaml 
2. Update the with the Sentiment Analysis Key you obtained from Step 3 in the Azure Setup. 
3. kubectl apply -f Deploy/deploy.yaml

# Azure DevOps

## Build 
1. Create a new build pipeline
2. Select GitHub. 
    * Create a service connection to your GitHub repo
3. Select repo and branch
4. Select Yaml 
5. Select Hosted Ubuntu for build server
6. Select Build/pipeline.yaml for Yaml configuration 

## Deploy with Helm
1. Create a new release pipeline
2. Select Empty Job
3. Add the Build from above as the source Artifacts to deploy
4. Add Task - 'Helm Tool Installer'
5. Add Task - 'Pack and Deploy Helm Chart'
6. Configure Helm Chart Task
    * Specify your connection to your Kubernetes Cluster - either using a Azure Service Connection or Kubernetes Service Connection 
    * Command: Upgrade
    * Chart Type: Name
    * Chart Name: $(System.DefaultWorkingDirectory)/_Chat/drop/Helm
    * Release Name: sentiment-chat
    * Set Values: sentimentapi_key=$(cogsApiKey),chat_image_version=$(Build.BuildId) 
    * Value File: $(System.DefaultWorkingDirectory)/_Chat/drop/Helm/values.yaml
7. Create Release Variable named cogsApiKey and set it to the key for your Sentiment Analysis API. 


