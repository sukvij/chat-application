package model

type Message struct {
	ID       int64  `json:"id"`
	FromUser int64  `json:"from_user"`
	ToUser   int64  `json:"to_user"`
	Detail   string `json:"detail"`
}
