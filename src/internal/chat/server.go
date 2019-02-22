package chat
import (
	"net/http"
	"fmt"
	"log"
	"github.com/rs/cors"
)

type HttpServer struct {}

func (h *HttpServer) optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *HttpServer) handleConnections(hub *Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.readMessages()
	go client.writeMessages()
}

func (h *HttpServer) InitHttpServer(port string) {

	hub := newHub()
	go hub.run()

	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("../public"))) 
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		h.handleConnections(hub, w, r)
	}) 

	server := cors.Default().Handler(router)
	fmt.Print("Listening on ", port)
	log.Fatal(http.ListenAndServe( port , server))
}

func NewChatServer() *HttpServer {
	return &HttpServer{}
}