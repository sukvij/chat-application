package model

import (
	model "rojgaarkaro-backend/basemodel"
)

type User struct {
	model.BaseModel
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsMember  bool   `json:"is_member"`
	Priority  int    `json:"priority"`
	Verified  bool   `json:"verified"`
	Image     string `json:"image"`
}
