# Go Webfinger Server
This is a small Go server built with [Gin](https://github.com/gin-gonic/gins) to serve Webfinger content for Mastodon users. 

## Webfinger
To learn more about Webfinger and it's place in ActivityPub, read more [in the Mastodon documentation](https://docs.joinmastodon.org/spec/webfinger/). 

## Usage
This application contains a `payload-example.json` file with Webfinger template data. Replace the templated data `domain.com` and `username` with the your own instance and username data. This will be served at port `8080` on the server. 

## Running Go Webfinger
I provide several ways to run Go Webfinger: 
1. Compiling the application using `go build .` and then running it directly with `./webfinger`. 
2. Building the Docker container using `docker build- -t webfinger .` and running it through the Docker engine using `docker run -p 8080:8080 webfinger`. This will forward your host port `8080` to the container port `8080`. 
3. Run the application using `docker compose up --build` to build and run the docker container, exposing port `8080` as above. 

Whichever way, the payload data will be available at `http://localhost:8080/.well-known/webfinger`. 

## Kubernetes
If you're a madlad, you can deploy the application to a Kubernetes cluster. 
1. Change the Dockerfile and `deployment/kubernetes/deployment.yaml` to reference your own registry image. 
2. `kubectl apply -f deployment/kubernetes/deployment.yaml` 
3. `kubectl apply -f deployment/kubernetes/service.yaml`
You'll have to work out the cloud gateway yourself to serve the traffic to your cluster. But this will get it started. 