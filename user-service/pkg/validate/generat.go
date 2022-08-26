package validate

import (
	"crypto/rand"
	"io"

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

func GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
