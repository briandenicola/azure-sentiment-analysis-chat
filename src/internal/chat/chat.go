package chat

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type newChatHandler struct { }
func (eh *newChatHandler) optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (eh *newChatHandler) handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func (eh *newChatHandler) handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (eh *newChatHandler) InitHttpServer(port string) {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("../public"))) 
    router.HandleFunc("/ws", eh.handleConnections) 
	go eh.handleMessages()

	server := cors.Default().Handler(router)
	fmt.Print("Listening on ", port)
	log.Fatal(http.ListenAndServe( port , server))
}
func NewChatHandler() *newChatHandler {return &newChatHandler{}}