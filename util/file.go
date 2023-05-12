package util

import (
	"bufio"
	"io"
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

func ReadFile(path string) ([]string, error) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, 64)

	var nameList []string
	var tmp []byte

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tmp = append(tmp, line...)
		nameList = append(nameList, string(tmp))
		tmp = []byte{}

	}

	return nameList, nil

}
