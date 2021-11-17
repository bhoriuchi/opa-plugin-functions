package utils

import (
	"os"
	"strconv"
)

func ParseFilemode(str string) (os.FileMode, error) {
	fileMode, err := strconv.ParseInt(str, 8, 64)
	if err != nil {
		return 0, err
	}

	return os.FileMode(fileMode), nil
}
