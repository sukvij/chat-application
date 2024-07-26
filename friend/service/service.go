package service

import (

	// awsDownload "rojgaarkaro-backend/aws/download"
	// awsUpload "rojgaarkaro-backend/aws/upload"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	friendModel "rojgaarkaro-backend/friend/model"
	friendRepo "rojgaarkaro-backend/friend/repository"
	userModel "rojgaarkaro-backend/user/model"
	userService "rojgaarkaro-backend/user/service"

	"gorm.io/gorm"
)

type Service struct {
	Db     *gorm.DB
	Friend *friendModel.Friend
}

func NewService(db *gorm.DB, friend *friendModel.Friend) *Service {
	return &Service{
		Db:     db,
		Friend: friend,
	}
}

func (service *Service) GetFriendsById() ([]*userModel.User, errorWithDetails.ErrorWithDetails) {
	repository := friendRepo.NewRepository(service.Db, service.Friend)
	result, err := repository.GetFriendsById()
	users := []*userModel.User{}
	for _, val := range result.FriendsList {
		id := val.Id
		user := &userModel.User{}
		user.ID = id
		userServiceMethod := userService.NewService(service.Db, user)
		x, _ := userServiceMethod.GetUser()
		users = append(users, x)
	}
	return users, err
}

func (service *Service) ListOfNonFriendById() ([]*userModel.User, errorWithDetails.ErrorWithDetails) {
	repository := friendRepo.NewRepository(service.Db, service.Friend)
	result, err := repository.GetFriendsById()
	users := []*userModel.User{}
	for _, val := range result.FriendsList {
		id := val.Id
		user := &userModel.User{}
		user.ID = id
		userServiceMethod := userService.NewService(service.Db, user)
		x, _ := userServiceMethod.GetUser()
		users = append(users, x)
	}
	return users, err
}

func (service *Service) MakeFriend() errorWithDetails.ErrorWithDetails {
	repository := friendRepo.NewRepository(service.Db, service.Friend)
	repository.MakeFriend()
	return errorWithDetails.ErrorWithDetails{}
}
