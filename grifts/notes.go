package grifts

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/markbates/grift/grift"
)

func printFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	prefix := "templates/"
	href := path[len(prefix):]
	fmt.Println(href)
	return nil
}

var _ = Add("notes", func(c *Context) error {

	// dir := os.Args[1]
	dir := "templates/notes/"
	err := filepath.Walk(dir, printFile)
	if err != nil {
		fmt.Println(err)
	}
	return nil
})
