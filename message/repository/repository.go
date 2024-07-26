package repository

import (
	errorWithDetails "rojgaarkaro-backend/baseThing"
	messageModel "rojgaarkaro-backend/message/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db      *gorm.DB
	Message *messageModel.Message
}

func NewRepository(Db *gorm.DB, message *messageModel.Message) *Repository {
	return &Repository{Db: Db, Message: message}
}
func (repository *Repository) GetMessages() ([]*messageModel.Message, errorWithDetails.ErrorWithDetails) {
	message := repository.Message
	result := []*messageModel.Message{}
	repository.Db.Where("from_user = ? and to_user = ? or from_user = ? and to_user = ?", message.FromUser, message.ToUser, message.ToUser, message.FromUser).Find(&result)
	return result, errorWithDetails.ErrorWithDetails{Detail: ""}
}

func (repository *Repository) PostMessage() (*messageModel.Message, errorWithDetails.ErrorWithDetails) {
	message := repository.Message
	repository.Db.Create(message)
	return message, errorWithDetails.ErrorWithDetails{Detail: ""}
}
