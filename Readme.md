# Overview 

This is a demo application written in Go to showcase Azure Cognitive Services running in Docker Containers. 
It is a websocket chat web application.  It parses all messages using Azure Cognitive Services Sentimnet Analysis.
If the sentiment is below a set threshold (ie. the message is too mean) then the AI Chat Moderator will broadcast to everyone to be nice. 

This demo uses the Azure Cognitive Services [container deployment](https://docs.microsoft.com/en-us/azure/cognitive-services/cognitive-services-container-support). This means that all analysis is done locally and not in Azure. A Cognitive Services deployment in Azure is still required for billing purposes but all the analysis is done locally.

# Azure Setup
## Prerequisites
* AKS Cluster Created. See [kubernetes-cluster-setup](https://github.com/briandenicola/kubernetes-cluster-setup)
* Azure Storage Account for Terraform state storage
* Azure Container Repository for Helm Chart
* Docker Hub Repository for the docker container (could re-use Azur Container Repository)

## GitHub Actions
* Excute .github/workflows/infrastructure.yaml

## Manual
1. cd infrastructure
2. terraform init -backend=true -backend-config="access_key=WY4ruq........" -backend-config="key=app.terraform.tfstate"
3. terraform plan -out="uat.plan" -var "resource_group_name=DevSub02_ChatApp_RG" -var-file="uat.tfvars"
4. terraform apply "uat.plan"

# Code Build
## Local
1. Set your GOPATH 
2. go get github.com/gorilla/websocket
3. go get github.com/gin-contrib/cors
4. go get github.com/gin-gonic/gin
5. go get github.com/gin-gonic/contrib/static
6. go get gopkg.in/resty.v1
7. cd ./src
8. docker build -t chatw
9. cd ./chart
10. helm package .
11. az acr helm push -n ${ACR_NAME} sentimentchat-0.2.${VERAION}.tgz --force

## GitHub Actions
* Excute .github/workflows/build.yaml

# Deploy to Kubernetes 

## GitOps
1. The [repository](https://github.com/briandenicola/kubernetes-cluster-setup) is configured to deploy this application via flux
2. Update [this] https://github.com/briandenicola/kubernetes-cluster-setup/blob/master/cluster-manifests/uat/app-81e86b.yaml) yaml to the correct ${VERSION}
