package entity

import (
	"fmt"

	customErr "github.com/NeverlandMJ/ToDo/api-gateway/pkg/error"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/utilities"
)

type ReqPhone struct {
	Phone string `json:"phone"`
}

func (rp ReqPhone) CheckReqPhone() error {
	if err := utilities.PhoneNumber(rp.Phone); err != nil {
		return fmt.Errorf("%w: %s", customErr.ERR_INVALID_INPUT, err.Error())
	}
	return nil
}

type ReqCode struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

func (rc ReqCode) CheckReqCode() error {
	if rc.Code == "" || rc.Phone == "" {
		return fmt.Errorf("%w: empty input", customErr.ERR_INVALID_INPUT)
	}
	if err := utilities.PhoneNumber(rc.Phone); err != nil {
		return fmt.Errorf("%w: %s", customErr.ERR_INVALID_INPUT, err.Error())
	}

	return nil
}

type ReqSignIn struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (rsi ReqSignIn) CheckReqSignIn() error {
	if rsi.UserName == "" || rsi.Password == "" {
		return fmt.Errorf("%w: empty input", customErr.ERR_INVALID_INPUT)
	}
	return nil
}
