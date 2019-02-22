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

	cogsUrl := "http://cogs:5000/text/analytics/v2.0/sentiment"
	if os.Getenv("COGSURL") != ""  {
		os.Getenv("COGSURL")
	}

	c := chat.NewChatServer()
	c.InitHttpServer(port, cogsUrl)
}