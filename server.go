package main

import (
	"github.com/show-t/gosvr/controller"
	"github.com/show-t/gosvr/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//var templates map[string]*template.Template

func main() {
	// DBマイグレーション
	model.DB.AutoMigrate(&model.User{})

	// Echoインスタンス生成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 静的ファイルのパス設定
	e.Static("/public/css/", "./public/css/")
	e.Static("/public/js/", "./public/js/")
	e.Static("/public/img/", "./public/img/")

	// ルーティング設定
	e.GET("/", controller.IndexViewController)

	u := e.Group("/user")
	u.GET("/signin", controller.SigninViewController)
	u.POST("/signin", controller.SigninController)
	u.GET("/signup", controller.SignupViewController)
	u.POST("/signup", controller.SignupController)

	// サーバー起動
	e.Logger.Fatal(e.Start(":3000"))

}

func init() {
	loadTemplates()
}

func loadTemplates() {

}
