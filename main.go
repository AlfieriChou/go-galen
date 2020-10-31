package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/AlfieriChou/go-galen/routes"
)

func main() {
	router := routes.CreateRouter()
	server := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}

	log.Println("Server exiting")
}
