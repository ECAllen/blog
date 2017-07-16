package lib

import (
	"fmt"
	"os"
)

var Prefix string

func PrintFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	href := path[len(Prefix):]
	fmt.Println(href)
	return nil
}
