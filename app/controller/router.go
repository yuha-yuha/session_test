package controller

import (
	"app/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ServerSetup() *gin.Engine {
	//サーバー、クッキーを使ったセッションのセットアップ
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//セッションクッキーの例
	//
	store.Options(sessions.Options{
		Path:   "/",
		Domain: "localhost",
	})
	//

	//永続クッキーの例
	/*store.Options(sessions.Options{
		Path: "/",
		Domain: "localhost",
		MaxAge: 3600, //一時間有効
	})*/

	r.LoadHTMLGlob("view/*.html")
	r.Static("/assets", "./assets")

	r.GET("/login", LoginPage)
	r.POST("/login", Login)
	r.GET("/", RootRouting)
	//r.GET("/signup", SignupPage)
	//r.POST("/signup", Signup)

	auth := r.Group("/auth")
	auth.Static("/assets", "./assets")
	auth.Use(middleware.IsLoggedIn)
	{
		auth.GET("/user", UserPage)
		auth.POST("/logout", Logout)
	}

	return r

}
