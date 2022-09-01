package entity

import (
	"fmt"

	customErr "github.com/NeverlandMJ/ToDo/api-gateway/pkg/error"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/utilities"
)

// CheckReqPhone checks if users input valid
func (rp ReqPhone) CheckReqPhone() error {
	if err := utilities.PhoneNumber(rp.Phone); err != nil {
		return fmt.Errorf("%w: %s", customErr.ERR_INVALID_INPUT, err.Error())
	}
	return nil
}

// CheckReqCode checks if the users inputs for sign up is valid
func (rc ReqCode) CheckReqCode() error {
	if rc.Code == "" || rc.Phone == "" {
		return fmt.Errorf("%w: empty input", customErr.ERR_INVALID_INPUT)
	}
	if err := utilities.PhoneNumber(rc.Phone); err != nil {
		return fmt.Errorf("%w: %s", customErr.ERR_INVALID_INPUT, err.Error())
	}

	return nil
}

// CheckReqSignIn is used to check if user's input for sign in is valid
func (rsi ReqSignIn) CheckReqSignIn() error {
	if rsi.UserName == "" || rsi.Password == "" {
		return fmt.Errorf("%w: empty input", customErr.ERR_INVALID_INPUT)
	}
	return nil
}

// CheckReqCreateTodo is used to check if user's input for creating todo is valid
func (rct ReqCreateTodo) CheckReqCreateTodo() error {
	if rct.Body == "" {
		return fmt.Errorf("%w: empty body", customErr.ERR_INVALID_INPUT)
	}
	if rct.Deadline == "" {
		return fmt.Errorf("%w: deadline", customErr.ERR_INVALID_INPUT)
	}

	return nil
}