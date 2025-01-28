package main

import (
	"goBackend/internal/server"
)

func main() {
	server := server.NewAPIServer(":3000")
	server.Run()
}
