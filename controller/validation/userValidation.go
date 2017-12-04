package validation

import (
	"github.com/labstack/echo"
)

type (
	User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
)

func ToUser(c echo.Context) (*User, bool) {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return nil, false
	}
	return user, true
}
