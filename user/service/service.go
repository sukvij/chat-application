package service

import (
	"mime/multipart"

	// awsDownload "rojgaarkaro-backend/aws/download"
	// awsUpload "rojgaarkaro-backend/aws/upload"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	userModel "rojgaarkaro-backend/user/model"
	userRepo "rojgaarkaro-backend/user/repository"

	"gorm.io/gorm"
)

type Service struct {
	Db   *gorm.DB
	User *userModel.User
}

func NewService(db *gorm.DB, user *userModel.User) *Service {
	return &Service{
		Db:   db,
		User: user,
	}
}

type ServiceMethod interface {
	GetAllUsers() (*[]userModel.User, errorWithDetails.ErrorWithDetails)
	GetUser() (*userModel.User, errorWithDetails.ErrorWithDetails)
	CreateUser() errorWithDetails.ErrorWithDetails
	DeleteUser() errorWithDetails.ErrorWithDetails
	UpdateUser() errorWithDetails.ErrorWithDetails
	GetUserByEmail() (*userModel.User, errorWithDetails.ErrorWithDetails)
	UploadFile(*multipart.FileHeader) errorWithDetails.ErrorWithDetails
	DownloadFile() errorWithDetails.ErrorWithDetails
}

// func (service *Service) UploadFile(fileheader *multipart.FileHeader) errorWithDetails.ErrorWithDetails {
// 	location, err := awsUpload.Upload(fileheader, service.User.ID)
// 	if err != nil {
// 		return errorWithDetails.ErrorWithDetails{Status: 10011010, Detail: "aws upload file error"}
// 	}

// 	fmt.Println("2")
// 	// first we will get the user
// 	service.User.Image = location
// 	service.UpdateUser()
// 	repository := userRepo.NewRepository(service.Db, service.User)
// 	err1 := repository.UpdateUser()
// 	if err1.Detail != "" {
// 		return err1
// 	}
// 	return errorWithDetails.ErrorWithDetails{}
// }

// func (service *Service) DownloadFile() errorWithDetails.ErrorWithDetails {
// 	user, err := service.GetUser()
// 	if err.Detail != "" {
// 		return err
// 	}
// 	awsDownload.Download(user.Image)
// 	return errorWithDetails.ErrorWithDetails{}
// }

func (service *Service) GetAllUsers() (*[]userModel.User, errorWithDetails.ErrorWithDetails) {
	// var result *[]userModel.User

	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetAllUsers()
	return result, err
}

func (service *Service) GetUser() (*userModel.User, errorWithDetails.ErrorWithDetails) {
	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetUser()
	return result, err
}

func (service *Service) GetUserByEmail() (*userModel.User, errorWithDetails.ErrorWithDetails) {
	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetUserByEmail()
	return result, err
}

func (service *Service) CreateUser() errorWithDetails.ErrorWithDetails {
	repository := userRepo.NewRepository(service.Db, service.User)
	// first we will checck whether gmail exist or not
	_, err := service.GetUserByEmail()
	if err.Detail != "" {
		if err.Status == 404 {
			err = repository.CreateUser()
			return err
		}
		return err
	}
	return errorWithDetails.ErrorWithDetails{Status: 409, Detail: "email already exist"}
}

func (service *Service) DeleteUser() errorWithDetails.ErrorWithDetails {
	repository := userRepo.NewRepository(service.Db, service.User)
	err := repository.DeleteUser()
	return err
}

func (service *Service) UpdateUser() errorWithDetails.ErrorWithDetails {
	user, err := service.GetUser()
	if err.Detail != "" {
		return err
	}
	updatedUser := service.User

	if updatedUser.FirstName != "" {
		user.FirstName = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		user.LastName = updatedUser.LastName
	}

	if updatedUser.Contact != "" {
		user.Contact = updatedUser.Contact
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}
	if updatedUser.Image != "" {
		user.Image = updatedUser.Image
	}
	user.UpdatedAt = updatedUser.UpdatedAt

	repository := userRepo.NewRepository(service.Db, user)
	err = repository.UpdateUser()
	if err.Detail != "" {
		return err
	}
	return errorWithDetails.ErrorWithDetails{}

	// we will update all the fields which are not null

	// types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i))
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i).Name)
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(values.Field(i).Type())
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(values.Field(i))
	// }

	// valuesForUpdation := reflect.ValueOf(updatedUser)
	// types := valuesForUpdation.Type()
	// valueUser := reflect.ValueOf(user)
	// for i := 0; i < valuesForUpdation.NumField(); i++ {
	// 	// if values.Field(i).Type().String() == "int" {

	// 	// }
	// 	switch valuesForUpdation.Field(i).Type().String() {
	// 	case "int":
	// 		// we will edit this later
	// 	case "float":
	// 		// we will edit this later
	// 	case "bool":
	// 		// we will edit this later
	// 	case "string":
	// 		fieldName := types.Field(i).Name
	// 		fmt.Println(fieldName, valuesForUpdation.Field(i))
	// 		valueUser = valuesForUpdation.Field(i)
	// 	}
	// }

}
