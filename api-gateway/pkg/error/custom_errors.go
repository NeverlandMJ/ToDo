package error

import "fmt"

var ERR_USER_EXIST = fmt.Errorf("user already exists")
var ERR_CODE_HAS_EXPIRED = fmt.Errorf("code has been expired")
var ERR_INCORRECT_CODE = fmt.Errorf("code doesn't match")
