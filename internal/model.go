package internal

import "time"

type BaseModel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// DeletedAt *time.Time
}

type Error struct {
	StatusCode int
	Message    string
}

type Response struct {
	Succeed bool
	Error   *Error
	Result  []interface{}
}

func NewResponse(succeed bool, err string, result []interface{}) *Response {
	response := &Response{}
	if succeed {
		response.Succeed = true
		response.Error = nil
		response.Result = result
	} else {
		response.Succeed = false
		response.Result = nil
		response.Error = &Error{StatusCode: 0, Message: err}
	}
	return response
}
