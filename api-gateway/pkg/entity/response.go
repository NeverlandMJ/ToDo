package entity

type RespUser struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type SendCodeID struct {
	Sid string `json:"sid,omitempty"`
}

type RespTodo struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Deadline  string `json:"deadline"`
	IsDone    bool   `json:"is_done"`
}
