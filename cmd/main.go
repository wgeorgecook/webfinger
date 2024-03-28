package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpserver "github.com/wgeorgecook/webfinger/internal/http"
)

func main() {
	log.Println("Hello, world!")
	defer log.Println("Goodbye!")

	go httpserver.Start()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpserver.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("done!")
	}
	log.Println("Server exiting")

}
