package main

import (
	"log"

	"github.com/kenshin579/analyzing-go-web-server-tips/server"
)

func main() {
	srv := server.Server()
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
