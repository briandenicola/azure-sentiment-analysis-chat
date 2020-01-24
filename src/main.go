package main

import (
	"os"
	"log"
	"github.com/bjd145/go/chat"
)

//Variables 
var	port    = ":8081"
var cogsUrl = "http://cogs:5000/text/analytics/v2.0/sentiment"

func main() {
	log.Print("Starting Server . . ")

	if os.Getenv("GOPORT") != "" {
		port = os.Getenv("GOPORT")
	} 
	if os.Getenv("COGSURL") != ""  {
		cogsUrl = os.Getenv("COGSURL")
	}

	c := chat.CreateServer(port, cogsUrl)
	c.RunServer()
}