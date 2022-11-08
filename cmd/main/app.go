package main

import (
	"github.com/Zajroid/rest-api-go/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("Create router")
	router := httprouter.New()

	log.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("Start application")

	listener, err := net.Listen("tcp", "0.0.0.0:1235")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is listening port 0.0.0.0:1235")
	log.Fatal(server.Serve(listener))
}
