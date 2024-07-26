package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type FriendID struct {
	Id int64 `json:"id"`
}

type FriendIDs []*FriendID

// Scan implements the sql.Scanner interface for FriendIDs
func (f *FriendIDs) Scan(value interface{}) error {
	if value == nil {
		*f = FriendIDs{}
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, f)
	case string:
		return json.Unmarshal([]byte(v), f)
	default:
		return errors.New("unsupported data type")
	}
}

// Value implements the driver.Valuer interface for FriendIDs
func (f FriendIDs) Value() (driver.Value, error) {
	return json.Marshal(f)
}

type Friend struct {
	UserId      int64     `json:"user_id"`
	FriendsList FriendIDs `json:"friends_list" gorm:"type:jsonb"`
}
