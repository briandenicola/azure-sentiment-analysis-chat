package chat
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"log"
)

type HttpServer struct {}

func (h *HttpServer) optionsHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}

//func (h *HttpServer) handleConnections(hub *Hub, cogs string, w http.ResponseWriter, r *http.Request) {
func (h *HttpServer) handleConnections(hub *Hub, cogs string, c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), cogsUrl: cogs}
	client.hub.register <- client

	go client.readMessages()
	go client.writeMessages()
}

func (h *HttpServer) InitHttpServer(port string, cogs string) {

	hub := newHub()
	go hub.run()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {
		h.handleConnections(hub, cogs, c)
	})
	
	//router.StaticFS("/", http.Dir("../public"))
	router.Use(static.Serve("/", static.LocalFile("../public", true)))
	router.OPTIONS("/", h.optionsHandler)

	log.Println("Server listener started on %s", port)
	router.Run(port) 	
}

func NewChatServer() *HttpServer {
	return &HttpServer{}
}