package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"./functions"
	"./handles"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", handles.IndexHandle)
	handler.HandleFunc("/add", handles.AddHandle)
	handler.HandleFunc("/save", handles.SaveHandle)
	handler.HandleFunc("/response", handles.ResponseHandle)
	//handler.HandleFunc("/manage/", handleInfo)*/


	functions.OpenDatabase()


	srv := http.Server{
		Addr:              ":8000",
		Handler:           handler,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()
	log.Print("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Printf("Got an error...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")

}