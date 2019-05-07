package gaialisp

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(fileName string) string {
	buff, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("file [%s] not found", fileName))
	}

	str := string(buff)

	return str
}
