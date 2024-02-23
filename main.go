package main

import (
	"botwy/family-chat-server/models"
	"botwy/family-chat-server/storage"
	"botwy/family-chat-server/utils"
	"botwy/family-chat-server/wsEvents"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsConnections []*websocket.Conn

func main() {
	router := gin.Default()
	router.GET("/messages", getMessages)
	router.POST("/message", postMessage)

	router.GET("/close-all-ws", func(context *gin.Context) {
		for _, wsConnection := range wsConnections {
			if wsConnection != nil {
				wsConnection.Close()
			}
		}
		wsConnections = []*websocket.Conn{}
	})

	router.GET("/ws/subscribe/changes", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		fmt.Println("conn")
		wsConnections = append(wsConnections, conn)
		conn.SetCloseHandler(func(code int, text string) error {
			fmt.Printf("connection closed, code: %d, text: %q", code, text)
			connIndex := utils.FindConnIndex(wsConnections, conn)
			fmt.Printf("connection index: %d: ", connIndex)
			if connIndex > -1 {
				wsConnections = utils.RemoveConn(wsConnections, connIndex)
			}
			return nil
		})
	})
	router.Run(":8080")
}

func sendWsMulticast() {
	for _, wsConnection := range wsConnections {
		go sendWsEvent(wsConnection)
	}
}

func sendWsEvent(wsConnection *websocket.Conn) {
	if wsConnection == nil {
		return
	}
	wsConnection.WriteMessage(websocket.TextMessage, []byte(wsEvents.MESSAGES_DID_CHANGE))
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Messages)
}

func postMessage(c *gin.Context) {
	var newMessage models.Message

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	// Add the new album to the slice.
	storage.Messages = append(storage.Messages, newMessage)
	c.IndentedJSON(http.StatusOK, newMessage)
	go sendWsMulticast()
}
