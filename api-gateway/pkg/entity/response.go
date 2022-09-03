package entity

// RespUser is used to to sign in
type RespUser struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

// RespSentCode is used to totp
type RespSentCode struct {
	Sid string `json:"sid,omitempty"`
}

// RespTodo todo's body
type RespTodo struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Deadline  string `json:"deadline"`
	IsDone    bool   `json:"is_done"`
}
