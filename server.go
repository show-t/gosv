package main

import (
	"github.com/show-t/gosvr/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//var templates map[string]*template.Template

func main() {
	// Echoインスタンス生成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 静的ファイルのパス設定
	e.Static("/public/css/", "./public/css")
	e.Static("/public/js/", "./public/js/")
	e.Static("/public/img/", "./public/img/")

	// ルーティング設定
	e.GET("/", controllers.HandleIndexView)

	u := e.Group("/user")
	u.GET("/signin", controllers.HandleSigninView)
	u.POST("/signin", controllers.HandleSignin)
	u.GET("/signup", controllers.HandleSignupView)
	u.POST("/signup", controllers.HandleSignup)

	// サーバー起動
	e.Logger.Fatal(e.Start(":3000"))

}

func init() {
	loadTemplates()
}

func loadTemplates() {

}
