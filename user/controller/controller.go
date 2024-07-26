package controller

import (
	"fmt"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	userModel "rojgaarkaro-backend/user/model"
	userService "rojgaarkaro-backend/user/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func UserApis(app *gin.Engine, DB *gorm.DB) {
	db = DB
	AllUserApis := app.Group("/user")
	{
		AllUserApis.POST("/signin/userEmail/:userEmail/userPassword/:userPassword", checkUserExistance, getUserByEmail)
		AllUserApis.GET("/", getAllUsers)
		AllUserApis.GET("/id/:userId", getUser)
		AllUserApis.GET("/userEmail/:userEmail", getUserByEmail)
		// AllUserApis.Post("/signin/{userEmail}/{userPassword}", checkUserExistance, authentication.GenerateToken, getUserByEmail)
		// AllUserApis.Get("/", authentication.VerifyMiddleware, getAllUsers)
		// AllUserApis.Get("/{userId}", authentication.VerifyMiddleware, getUser)
		// // AllUserApis.Get("/{userEmail}", authentication.VerifyMiddleware, getUserByEmail)
		// AllUserApis.Post("/{userEmail}", authentication.VerifyMiddleware, getUserByEmail, resetPassword)
		AllUserApis.POST("/signUp", createUser)
		// AllUserApis.Delete("/{userId}", authentication.VerifyMiddleware, deleteUser)
		// AllUserApis.Put("/{userId}", authentication.VerifyMiddleware, updateUser)
		// AllUserApis.Post("/image/upload", authentication.VerifyMiddleware, uploadFile)
		// AllUserApis.Get("/download/image", downloadImage)
		// AllUserApis.Get("/logout", authentication.VerifyMiddleware, authentication.Logout)
	}
}

func checkUserExistance(ctx *gin.Context) {
	userEmail := ctx.Param("userEmail")
	userPassword := ctx.Param("userPassword")
	fmt.Println(userEmail, userPassword)
	user := &userModel.User{Email: userEmail, Password: userPassword}
	service := &userService.Service{Db: db, User: user}
	userExist, err := service.GetUserByEmail()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}

	if userExist.Password != userPassword {
		err1 := errorWithDetails.ErrorWithDetails{Status: 401, Detail: "wrong password"}
		ctx.AbortWithStatusJSON(err1.Status, err1)
		return
	}
	ctx.Next()
}

func getAllUsers(ctx *gin.Context) {
	service := &userService.Service{Db: db, User: &userModel.User{}}
	result, err := service.GetAllUsers()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, result)
}

func getUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user := &userModel.User{}
	user.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: user}
	result, err := service.GetUser()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, result)
}

// func resetPassword(ctx *gin.Context) {
// 	// enter gmail address to check whether email exist or not.
// 	userEmail := ctx.Params().Get("userEmail")
// 	userPassword := ctx.Params().Get("userPassword")
// 	user := &userModel.User{Email: userEmail, Password: userPassword}
// 	service := &userService.Service{Db: db, User: user}
// 	updatedUser, err := service.GetUserByEmail()
// 	if err.Detail != "" {
// 		ctx.StopWithJSON(err.Status, err)
// 		return
// 	}
// 	service.User = updatedUser
// 	errs := service.UpdateUser()
// 	if errs.Detail != "" {
// 		ctx.StopWithJSON(errs.Status, errs)
// 		return
// 	}
// 	ctx.JSON("password reset successfully")
// }

func getUserByEmail(ctx *gin.Context) {
	userEmail := ctx.Param("userEmail")
	user := &userModel.User{}
	user.Email = userEmail
	service := &userService.Service{Db: db, User: user}
	result, err := service.GetUserByEmail()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, result)
	ctx.Next()
}

func createUser(ctx *gin.Context) {
	var user userModel.User
	ctx.BindJSON(&user)
	service := &userService.Service{Db: db, User: &user}
	err := service.CreateUser()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, "successfully created")
}

func deleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user := &userModel.User{}
	user.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: user}
	err := service.DeleteUser()
	if err.Detail != "" {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(200, "user deleted successfully")
}

func updateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	updatedUser := &userModel.User{}
	ctx.BindJSON(&updatedUser)
	updatedUser.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: updatedUser}
	err := service.UpdateUser()
	if err.Detail == "" {
		user, errs := service.GetUser()
		if errs.Detail != "" {
			ctx.JSON(err.Status, err)
			return
		}
		ctx.JSON(200, user)
	} else {
		ctx.JSON(err.Status, err)
	}
}

// func uploadFile(ctx *gin.Context) {

// 	userId := ctx.FormValue("id")
// 	file, fileHeader, err := ctx.FormFile("file")
// 	if err != nil {
// 		ctx.StopWithError(iris.StatusBadRequest, err)
// 		return
// 	}
// 	user := &userModel.User{}
// 	user.ID, _ = strconv.ParseInt(userId, 10, 64)
// 	service := &userService.Service{Db: db, User: user}
// 	err1 := service.UploadFile(fileHeader)
// 	if err1.Detail != "" {
// 		fmt.Print("hi")
// 		ctx.StopWithJSON(err1.Status, err1)
// 	}
// 	ctx.JSON("file upload successfully")
// 	defer file.Close()
// }

// func downloadImage(ctx *gin.Context) {
// 	userId := ctx.FormValue("id")
// 	user := &userModel.User{}
// 	user.ID, _ = strconv.ParseInt(userId, 10, 64)
// 	service := &userService.Service{Db: db, User: user}
// 	err := service.DownloadFile()
// 	ctx.JSON(err)
// }
