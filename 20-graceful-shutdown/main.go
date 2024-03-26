package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(5 * time.Second)
		writer.Write([]byte("Hello World"))
	})

	go func() {
		log.Println("server is listening on port 3000")
		err := server.ListenAndServe()
		if err != nil && http.ErrServerClosed != err {
			log.Fatalf("could not listen on %s:%v", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shutting down server")
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server closed")
}
