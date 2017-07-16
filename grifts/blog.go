package grifts

import (
	"fmt"
	"github.com/ECAllen/blog/lib"
	. "github.com/markbates/grift/grift"
	"path/filepath"
)

var _ = Add("blog", func(c *Context) error {
	dir := "templates/blog-posts/posts"
	lib.Prefix = "templates/"
	err := filepath.Walk(dir, lib.PrintFile)
	if err != nil {
		fmt.Println(err)
	}
	return nil
})
