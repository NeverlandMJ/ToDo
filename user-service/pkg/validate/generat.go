package validate

import (
	"github.com/cip8/autoname"
	"github.com/sethvargo/go-password/password"
)

func GenerateUserName() string {
	name := autoname.Generate("_")

	return name
}

func GeneratePassword() (string, error) {
	res, err := password.Generate(10, 3, 0, false, false)
	if err != nil {
		return "", err
	}
	return res, nil
}
