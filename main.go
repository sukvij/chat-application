package main

import (
	"fmt"
	database "rojgaarkaro-backend/database"
	friendController "rojgaarkaro-backend/friend/controller"
	messageController "rojgaarkaro-backend/message/controller"
	userController "rojgaarkaro-backend/user/controller"
	websocket "rojgaarkaro-backend/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	app.GET("/ws/:from/:to", websocket.HandleConnections)
	userController.UserApis(app, db)
	friendController.FriendApis(app, db)
	messageController.MessagedApis(app, db)
	app.Run(":8080")

}
