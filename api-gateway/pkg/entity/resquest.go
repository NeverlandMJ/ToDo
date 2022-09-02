package entity

// ReqPhone is used to send TOTP code to the user
type ReqPhone struct {
	Phone string `json:"phone"`
}

//ReqCode is used for user to enter sent TOTP and sign up
type ReqCode struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

// ReqSignIn is used to sign up
type ReqSignIn struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ReqCreateTodo is used to create Todo
type ReqCreateTodo struct {
	Body     string `json:"body"`
	Deadline string `json:"deadline"`
}


