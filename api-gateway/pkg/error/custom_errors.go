package error

import "fmt"

var ERR_USER_EXIST = fmt.Errorf("user already exists")
var ERR_CODE_HAS_EXPIRED = fmt.Errorf("code has been expired")
var ERR_INCORRECT_CODE = fmt.Errorf("code doesn't match")
var ERR_INTERNAL = fmt.Errorf("internal server error")
var ERR_USER_BLOCKED = fmt.Errorf("user is blocked")
var ERR_USER_NOT_EXIST = fmt.Errorf("user doesn't exist")
var ERR_INCORRECT_PASSWORD = fmt.Errorf("incorrect password")
