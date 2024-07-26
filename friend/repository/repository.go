package repository

import (
	"fmt"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	friendModel "rojgaarkaro-backend/friend/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db     *gorm.DB
	Friend *friendModel.Friend
}

func NewRepository(Db *gorm.DB, friend *friendModel.Friend) *Repository {
	return &Repository{Db: Db, Friend: friend}
}

// type RepositoryMethod interface {
// 	GetAllUsers() (*[]userModel.User, errorWithDetails.ErrorWithDetails)
// 	GetUser() (*userModel.User, errorWithDetails.ErrorWithDetails)
// 	GetUserByEmail() (*userModel.User, errorWithDetails.ErrorWithDetails)
// 	CreateUser() errorWithDetails.ErrorWithDetails
// 	DeleteUser() errorWithDetails.ErrorWithDetails
// 	UpdateUser() errorWithDetails.ErrorWithDetails
// }

func (repository *Repository) GetFriendsById() (*friendModel.Friend, errorWithDetails.ErrorWithDetails) {
	result := friendModel.Friend{UserId: repository.Friend.UserId, FriendsList: []*friendModel.FriendID{}}
	repository.Db.Where("user_id = ?", result.UserId).First(&result)
	fmt.Println("Result -- ", result.FriendsList)
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, errorWithDetails.ErrorWithDetails{Status: 404, Detail: gorm.ErrRecordNotFound.Error()}
	// 	}
	// 	return nil, errorWithDetails.ErrorWithDetails{}
	// }
	return &result, errorWithDetails.ErrorWithDetails{Detail: ""}
}

func (repository *Repository) ListOfNonFriendById() (*friendModel.Friend, errorWithDetails.ErrorWithDetails) {
	result := friendModel.Friend{UserId: repository.Friend.UserId, FriendsList: []*friendModel.FriendID{}}
	repository.Db.Where("user_id = ?", result.UserId).First(&result)
	fmt.Println("Result -- ", result.FriendsList)
	return &result, errorWithDetails.ErrorWithDetails{Detail: ""}
}

func (repository *Repository) MakeFriend() errorWithDetails.ErrorWithDetails {
	userid := repository.Friend.UserId
	friendUserId := repository.Friend.FriendsList[0].Id
	fmt.Println("ids - ", userid, friendUserId)

	result1 := &friendModel.Friend{}
	result2 := &friendModel.Friend{}
	err1 := repository.Db.Where("user_id = ?", userid).First(&result1).Error
	err2 := repository.Db.Where("user_id = ?", friendUserId).First(&result2).Error
	fmt.Println("err1", err1, &result1.FriendsList)
	fmt.Println("err2", err2, &result2.FriendsList)
	result1.FriendsList = append(result1.FriendsList, &friendModel.FriendID{Id: friendUserId})
	result2.FriendsList = append(result2.FriendsList, &friendModel.FriendID{Id: userid})
	repository.Db.Where("user_id = ?", userid).Save(&result1)
	repository.Db.Where("user_id = ?", friendUserId).Save(&result2)
	return errorWithDetails.ErrorWithDetails{}
}
