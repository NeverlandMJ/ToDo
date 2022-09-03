package entity

// ReqPhone is used to send TOTP code to the user
type ReqPhone struct {
	Phone string `json:"phone" example:"+998937776655"`
}

//ReqSignUp is used for user to enter sent TOTP and sign up
type ReqSignUp struct {
	Code  string `json:"code" example:"183003"`
	Phone string `json:"phone" example:"+998937776655"`
}

// ReqSignIn is used to sign up
type ReqSignIn struct {
	UserName string `json:"user_name" example:"mono_liza_doggy"`
	Password string `json:"password" example:"jsahbdshdaa"`
}

// ReqCreateTodo is used to create Todo
type ReqCreateTodo struct {
	Body     string `json:"body" example:"wake up early"`
	Deadline string `json:"deadline" example:"Mon Sep 12 6:30:00 UTC 2022"`
}

// ReqUpdateBody is used to update Todo's body
type ReqUpdateBody struct {
	TodoID string `json:"todo_id" example:"eeebcf44-593c-4b19-9dd9-bd83d30d4681"`
	Body   string `json:"body" example:"make a cake"`
}

// ReqUpdateDeadline is used to update Todo's set deadline
type ReqUpdateDeadline struct {
	TodoID   string `json:"todo_id" example:"eeebcf44-593c-4b19-9dd9-bd83d30d4681"`
	Deadline string `json:"deadline" example:"Mon Sep 12 6:30:00 UTC 2022"`
}

// ReqChangePassword is used to change user's password
type ReqChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// ReqChangeUsername is used to change user's password
type ReqChangeUsername struct {
	UserName string `json:"user_name"`
}
