package service

import (

	// awsDownload "rojgaarkaro-backend/aws/download"
	// awsUpload "rojgaarkaro-backend/aws/upload"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	messageModel "rojgaarkaro-backend/message/model"
	messageRepo "rojgaarkaro-backend/message/repository"

	"gorm.io/gorm"
)

type Service struct {
	Db      *gorm.DB
	Message *messageModel.Message
}

func NewService(db *gorm.DB, message *messageModel.Message) *Service {
	return &Service{
		Db:      db,
		Message: message,
	}
}

func (service *Service) GetMessages() ([]*messageModel.Message, errorWithDetails.ErrorWithDetails) {
	repository := messageRepo.NewRepository(service.Db, service.Message)
	result, err := repository.GetMessages()
	return result, err
}

func (service *Service) PostMessage() (*messageModel.Message, errorWithDetails.ErrorWithDetails) {
	repository := messageRepo.NewRepository(service.Db, service.Message)
	result, err := repository.PostMessage()
	return result, err
}
