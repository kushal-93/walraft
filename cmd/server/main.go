package main

import (
	"log"

	"github.com/kushal-93/walraft/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Println("server listening at port :8080")
	log.Fatal(srv.ListenAndServe())
}
