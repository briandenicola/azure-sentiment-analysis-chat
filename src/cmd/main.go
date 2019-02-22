package main

import (
	"os"
	"log"
	"internal/chat"
)

func main() {
	log.Print("Starting Server . . ")

	port := ":8081"
	if os.Getenv("GOPORT") != "" {
		port = os.Getenv("GOPORT")
	} 

	c := chat.NewChatServer()
	c.InitHttpServer(port)
}