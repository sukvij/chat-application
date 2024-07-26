package controller

import (
	"fmt"
	"strconv"

	friendModel "rojgaarkaro-backend/friend/model"
	friendService "rojgaarkaro-backend/friend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func FriendApis(app *gin.Engine, DB *gorm.DB) {
	db = DB
	AllUserApis := app.Group("/friend")
	{
		AllUserApis.GET("/:userId", getFriendsById)
		AllUserApis.POST("/", makeFriend)
		AllUserApis.GET("/notFriends/:userId", listOfNonFriendById)
	}
}

func getFriendsById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	fmt.Println(userId)
	friend := &friendModel.Friend{}
	x, _ := strconv.Atoi(userId)
	friend.UserId = int64(x)
	service := &friendService.Service{Db: db, Friend: friend}
	result, _ := service.GetFriendsById()
	// if err.Detail != "" {
	// 	ctx.JSON(err.Status, err)
	// 	return
	// }
	ctx.JSON(200, result)
}

func listOfNonFriendById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	fmt.Println(userId)
	friend := &friendModel.Friend{}
	x, _ := strconv.Atoi(userId)
	friend.UserId = int64(x)
	service := &friendService.Service{Db: db, Friend: friend}
	result, _ := service.ListOfNonFriendById()
	// if err.Detail != "" {
	// 	ctx.JSON(err.Status, err)
	// 	return
	// }
	ctx.JSON(200, result)
}

func makeFriend(ctx *gin.Context) {
	friend := &friendModel.Friend{}
	ctx.BindJSON(friend)
	fmt.Println(friend.UserId, friend.FriendsList[0])
	service := &friendService.Service{Db: db, Friend: friend}
	err := service.MakeFriend()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, "friend successfully")
}
