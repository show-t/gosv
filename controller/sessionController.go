package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/show-t/gosvr/controller/validation"
	"github.com/show-t/gosvr/model"
)

func SigninController(c echo.Context) (err error) {

	return
}

func SignupController(c echo.Context) (err error) {

	user, ok := validation.ToUser(c)
	if !ok {
		return
	}
	err = model.CreateUser(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"err": "データベースエラー",
		})
		return
	}

	c.JSON(http.StatusCreated, echo.Map{
		"success": true,
	})

	return
}
