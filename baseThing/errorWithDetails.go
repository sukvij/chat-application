package basething

type ErrorWithDetails struct {
	Status int
	Detail string
}

// func NewErrorWithDetails(code int, Type string, Detail string) *ErrorWithDetails {
// 	return &ErrorWithDetails{
// 		Type:   Type,                  //"type": "https://user",
// 		Status: code,                  // 404
// 		Title:  iris.StatusText(code), // status not found
// 		Detail: Detail,                // gmail doesnot exist
// 	}
// }
