package repository

import (
	"fmt"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	userModel "rojgaarkaro-backend/user/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db   *gorm.DB
	User *userModel.User
}

func NewRepository(Db *gorm.DB, User *userModel.User) *Repository {
	return &Repository{Db: Db, User: User}
}

type RepositoryMethod interface {
	GetAllUsers() (*[]userModel.User, errorWithDetails.ErrorWithDetails)
	GetUser() (*userModel.User, errorWithDetails.ErrorWithDetails)
	GetUserByEmail() (*userModel.User, errorWithDetails.ErrorWithDetails)
	CreateUser() errorWithDetails.ErrorWithDetails
	DeleteUser() errorWithDetails.ErrorWithDetails
	UpdateUser() errorWithDetails.ErrorWithDetails
}

func (repository *Repository) GetAllUsers() (*[]userModel.User, errorWithDetails.ErrorWithDetails) {
	// var result *[]userModel.User
	result := &[]userModel.User{}
	err := repository.Db.Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorWithDetails.ErrorWithDetails{Status: 404, Detail: gorm.ErrRecordNotFound.Error()}
		}
		return nil, errorWithDetails.ErrorWithDetails{}
	}
	return result, errorWithDetails.ErrorWithDetails{}
}

func (repository *Repository) GetUser() (*userModel.User, errorWithDetails.ErrorWithDetails) {
	result := &userModel.User{}
	userId := repository.User.ID
	err := repository.Db.Where("id = ?", userId).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorWithDetails.ErrorWithDetails{Status: 404, Detail: gorm.ErrRecordNotFound.Error()}
		}
		return nil, errorWithDetails.ErrorWithDetails{}
	}
	return result, errorWithDetails.ErrorWithDetails{}
}

func (repository *Repository) GetUserByEmail() (*userModel.User, errorWithDetails.ErrorWithDetails) {
	result := &userModel.User{}
	userEmail := repository.User.Email
	err := repository.Db.Where("email = ?", userEmail).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorWithDetails.ErrorWithDetails{Status: 404, Detail: gorm.ErrRecordNotFound.Error()}
		}
		return nil, errorWithDetails.ErrorWithDetails{}
	}
	return result, errorWithDetails.ErrorWithDetails{}
}

func (repository *Repository) CreateUser() errorWithDetails.ErrorWithDetails {
	err := repository.Db.Create(&repository.User).Error
	fmt.Println(err)
	return errorWithDetails.ErrorWithDetails{}
}

func (repository *Repository) DeleteUser() errorWithDetails.ErrorWithDetails {
	result := &userModel.User{}
	err := repository.Db.Where("id = ?", repository.User.ID).Delete(&result).Error
	fmt.Println(err)
	return errorWithDetails.ErrorWithDetails{}
}

func (repository *Repository) UpdateUser() errorWithDetails.ErrorWithDetails {
	err := repository.Db.Save(&repository.User).Error
	if err != nil {
		return errorWithDetails.ErrorWithDetails{Status: 404, Detail: "not updated"}
	}
	return errorWithDetails.ErrorWithDetails{}
}
