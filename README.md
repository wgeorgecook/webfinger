# Go Webfinger Server
This is a small Go server built with [Gin](https://github.com/gin-gonic/gins) to serve Webfinger content for Mastodon users. 

## Webfinger
To learn more about Webfinger and it's place in ActivityPub, read more [in the Mastodon documentation](https://docs.joinmastodon.org/spec/webfinger/). 

## Usage
This application contains a `payload-example.json` file with Webfinger template data. Replace the templated data `domain.com` and `username` with the your own instance and username data. This will be served at port `8080` on the server. 

## Running Go Webfinger
I provide two ways to run Go Webfinger: 
1. Compiling the application using `go build .` and then running it directly with `./webfinger`. 
2. Building the Docker container using `docker build- -t webfinger .` and running it through the Docker engine using `docker run -p 8080:8080 webfinger`. This will forward your host port `8080` to the container port `8080`. 
Either way, the payload data will be available at `http://localhost:8080/.well-known/webfinger`. 