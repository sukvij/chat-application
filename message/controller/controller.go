package controller

import (
	"fmt"
	"strconv"

	messageModel "rojgaarkaro-backend/message/model"
	messageService "rojgaarkaro-backend/message/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func MessagedApis(app *gin.Engine, DB *gorm.DB) {
	db = DB
	AllMessageApis := app.Group("/message")
	{
		AllMessageApis.GET("/:from/:to", getMessages)
		AllMessageApis.POST("/", postMessage)
	}
}

func getMessages(ctx *gin.Context) {
	user1 := ctx.Param("from")
	user2 := ctx.Param("to")
	from, _ := strconv.Atoi(user1)
	to, _ := strconv.Atoi(user2)
	message := &messageModel.Message{FromUser: int64(from), ToUser: int64(to)}
	service := &messageService.Service{Db: db, Message: message}
	result, _ := service.GetMessages()
	ctx.JSON(200, result)
}

func postMessage(ctx *gin.Context) {
	message := &messageModel.Message{}
	ctx.BindJSON(message)
	fmt.Println(&message)
	service := &messageService.Service{Db: db, Message: message}
	result, _ := service.PostMessage()
	ctx.JSON(200, result)
}
