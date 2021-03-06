package grifts

import (
	"fmt"
	"github.com/ECAllen/blog/lib"
	. "github.com/markbates/grift/grift"
	"path/filepath"
)

var _ = Add("notes", func(c *Context) error {
	dir := "templates/notes/"
	lib.Prefix = dir
	err := filepath.Walk(dir, lib.PrintFile)
	if err != nil {
		fmt.Println(err)
	}
	return nil
})
