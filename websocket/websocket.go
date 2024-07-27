package websocket

import (
	"fmt"
	"net/http"
	"strconv"

	messageModel "rojgaarkaro-backend/message/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatUserDetail struct {
	From int64
	To   int64
}

type Client struct {
	From int64
	To   int64
	Conn *websocket.Conn
}

var clients = make(map[ChatUserDetail]*Client)

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade connection: %v\n", err)
		return
	}

	from, _ := strconv.Atoi(c.Param("from"))
	to, _ := strconv.Atoi(c.Param("to"))
	chatUserDetail := ChatUserDetail{From: int64(from), To: int64(to)}
	client := &Client{From: int64(from), To: int64(to), Conn: conn}
	clients[chatUserDetail] = client

	for {
		msg := messageModel.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			delete(clients, chatUserDetail)
			conn.Close()
			break
		}
		fmt.Println(msg.FromUser, msg.ToUser, msg.Detail)
		if receiver, ok := clients[ChatUserDetail{From: chatUserDetail.To, To: chatUserDetail.From}]; ok {
			if err := receiver.Conn.WriteJSON(msg); err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			}
		}
	}
}
