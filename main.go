package main

import (
	"./functions"
	"./globalVars"
	"./handlers"
	"./middlewares"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var err error
	handler := http.NewServeMux()

	handler.HandleFunc("/show/", handleLink)
	handler.HandleFunc("/manage", middlewares.ContentChanger(handleInfo))

	globalVars.DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shortink")

	if err != nil {
		log.Fatal("Database is not open")
	}
	defer globalVars.DB.Close()

	err = globalVars.Links.FillFromDatabase(globalVars.DB)
	if err != nil {
		log.Fatal("Data not loaded")
	}

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

func handleLink(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handlers.ConnectHandler(w, r)
	} else {
		w = functions.GetResponse("Undefind method, please select another", fmt.Sprintf("Undefind method %s", r.Method), http.StatusOK, w)
	}
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handlers.GetHandler(w, r)
	} else if r.Method == http.MethodPost {
		handlers.PostHandler(w, r)
	} else if r.Method == http.MethodDelete {
		handlers.DeleteHandler(w, r)
	} else {
		w = functions.GetResponse("Undefind method, please select another", fmt.Sprintf("Undefind method %s", r.Method), http.StatusBadRequest, w)
	}
}
