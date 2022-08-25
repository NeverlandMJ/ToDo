package entity

type ReqPhone struct {
	Phone string `json:"phone"`
}

type ReqCode struct {
	Code string `json:"code"`
	Phone string `json:"phone"`
}
