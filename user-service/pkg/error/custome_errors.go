package error

import "fmt"

var ERR_USER_EXIST = fmt.Errorf("user already exists")
var ERR_INCORRECT_CODE = fmt.Errorf("code doesn't match")
var ERR_USER_NOT_EXIST = fmt.Errorf("user doesn't exist")
var ERR_INCORRECT_PASSWORD = fmt.Errorf("incorrect password")