package chat
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"log"
	"time"
)

type ChatServer struct {
	port         string
	sentimentUri string
}

func (h *ChatServer) optionsHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func (h *ChatServer) handleConnections(hub *Hub, c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), cogsUrl: h.sentimentUri}
	client.hub.register <- client

	go client.readMessages()
	go client.writeMessages()
}

func (h *ChatServer) RunServer()  {

	hub := newHub()
	go hub.run()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/ws", func(c *gin.Context) {
		h.handleConnections(hub, c)
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "i'm alive!"})
	})

	router.Use(static.Serve("/", static.LocalFile("../public", true)))
	router.OPTIONS("/", h.optionsHandler)

	log.Printf("Server listener started on %s", h.port)
	router.Run(h.port) 	
}

func CreateServer(port string, cogs string) *ChatServer {
	return &ChatServer{port: port, sentimentUri: cogs}
}