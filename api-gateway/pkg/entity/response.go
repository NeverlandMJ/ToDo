package entity

type RespUser struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type SendCodeID struct {
	Sid string `json:"sid,omitempty"`
}
