package main

import (
	"fmt"
	"net/http"

	"github.com/bonitech/base_gin/injection"
	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

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

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/hey/:mode", func(c *gin.Context) {
		mode := c.Param("mode")
		modeSelect := (map[bool]injection.StoreMode{true: injection.PROD, false: injection.MOCK})[mode == "prod"]
		store := injection.GetStore(modeSelect)
		m := &injection.Injection{Store: store}
		text, _ := m.Store.GetSomething()
		c.JSON(200, gin.H{
			"message": text,
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
