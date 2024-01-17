package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var srv *http.Server

func startServer() {
	r := gin.Default()

	r.GET("/.well-known/webfinger", webfinger)

	srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
func webfinger(c *gin.Context) {
	log.Println("webfinger request received")
	defer log.Println("finished webfinger request")

	// read the json payload we return
	payload, err := readPayload()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not read payload",
		})
		return
	}

	// marshal it to json
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not marshal payload",
		})
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Add("content-type", "application/jrd+json")
	c.Writer.Write(payloadJson)
	return
}
