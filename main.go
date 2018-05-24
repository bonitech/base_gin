package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/bonitech/base_gin/injection"
	"github.com/bonitech/base_gin/routes/api"
	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Load HTML templates
	tmpl := template.Must(template.ParseGlob("templates/*"))

	// Load Vue.js html
	const VueIndex = "frontend/dist/index.html"
	if _, err := os.Stat(VueIndex); err == nil {
		tmpl.ParseFiles(VueIndex)
	}
	r.SetHTMLTemplate(tmpl)

	// Set route group
	api.Routes(r)

	// Websocket Settings
	wsUser := 0
	m := melody.New()
	m.HandleConnect(func(s *melody.Session) {
		wsUser++
		fmt.Printf("User: %d\n", wsUser)
	})
	m.HandleDisconnect(func(s *melody.Session) {
		wsUser--
		fmt.Printf("User: %d\n", wsUser)
	})
	// Handle Requested message
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	r.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{"temp": "This is Value"})
	})

	// Dependency injection example
	r.GET("/hey/:mode", func(c *gin.Context) {
		mode := c.Param("mode")
		modeSelect := (map[bool]injection.StoreMode{true: injection.PROD, false: injection.MOCK})[mode == "prod"]
		store := injection.GetStore(modeSelect)
		m := &injection.Injection{Store: store}
		text, _ := m.Store.GetSomething()
		c.JSON(http.StatusOK, gin.H{
			"message": text,
		})
	})

	// Serve Vue.js
	r.Static("/static", "./frontend/dist/static")
	r.GET("/", func(c *gin.Context) {
		if _, err := os.Stat(VueIndex); err == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}
		c.String(http.StatusNotFound, "404 page not found")
	})

	// Fallback to vue.js route
	r.NoRoute(gin.HandlerFunc(func(c *gin.Context) {
		if _, err := os.Stat(VueIndex); err == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}
		c.String(http.StatusNotFound, "404 page not found")
	}))

	r.Run(":" + port)
}
