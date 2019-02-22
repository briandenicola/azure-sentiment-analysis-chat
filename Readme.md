docker pull mcr.microsoft.com/azure-cognitive-services/sentiment:latest

docker run --rm -d -p 5000:5000 --memory 4g --cpus 1 mcr.microsoft.com/azure-cognitive-services/sentiment 
Eula=accept \
Billing=https://southcentralus.api.cognitive.microsoft.com/ \
ApiKey=c34790d25d7641898cc6e66c2869cff4

http://localhost:5000/swaggerx