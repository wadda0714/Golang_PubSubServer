package util

import (
	"os"
)

func CreateFile(path string) (*os.File, error) {
	if IsFileExists(path) {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0664)
		if err != nil {
			return nil, err
		}
		return f, nil

	} else {
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		return f, nil

	}

}
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

}
