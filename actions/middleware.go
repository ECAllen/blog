package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
)

func SetVars(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		year := fmt.Sprintf("%d", time.Now().Year())
		c.Set("year", year)
		err := next(c)
		return err
	}
}
