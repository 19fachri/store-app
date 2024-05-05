package main

import "github.com/19fachri/store-app/internal/store_server"

func main() {
	storeServer := store_server.NewServer()
	storeServer.Start()
}
