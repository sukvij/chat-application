package main

import (
	"fmt"
	"net/http"
	database "rojgaarkaro-backend/database"
	friendController "rojgaarkaro-backend/friend/controller"
	messageController "rojgaarkaro-backend/message/controller"
	messageModel "rojgaarkaro-backend/message/model"
	userController "rojgaarkaro-backend/user/controller"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {

	// awsDownload.Download()
	db, err := database.Connection()
	if err != nil {
		return
	}
	fmt.Println("database connection success")
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// app.Post("/signin", authentication.GenerateToken)
	// app.Get("/logout", authentication.VerifyMiddleware, authentication.Logout)
	app.GET("/ws/:id", handleConnections)
	userController.UserApis(app, db)
	friendController.FriendApis(app, db)
	messageController.MessagedApis(app, db)
	app.Run(":8080")

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatUser struct {
	From int64
	To   int64
}

type Client struct {
	ID   int64
	Conn *websocket.Conn
}

var clients = make(map[int64]*Client)

func handleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade connection: %v\n", err)
		return
	}

	clientID, _ := strconv.Atoi(c.Param("id"))
	client := &Client{ID: int64(clientID), Conn: conn}
	clients[client.ID] = client

	for {
		msg := messageModel.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			delete(clients, client.ID)
			conn.Close()
			break
		}
		fmt.Println(msg.FromUser, msg.ToUser, msg.Detail)

		receiverID := msg.ToUser
		if receiver, ok := clients[receiverID]; ok {
			if err := receiver.Conn.WriteJSON(msg); err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			}
		}
	}
}
