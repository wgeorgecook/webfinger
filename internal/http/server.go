package httpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wgeorgecook/webfinger/internal/files"
)

var srv *http.Server

func Start() {
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

func Shutdown(ctx context.Context) error {
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
func webfinger(c *gin.Context) {
	log.Println("webfinger request received")
	defer log.Println("finished webfinger request")

	// read the json payload we return
	payload, err := files.ReadPayload()
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
}
