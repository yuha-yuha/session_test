package controller

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ServerSetup() *gin.Engine {
	//サーバー、クッキーを使ったセッションのセットアップ
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	//一時セッションと永続セッションの2つのセッション
	sessionNames := []string{"tempSession", "permaSession"}

	r.LoadHTMLGlob("/view/*/**")

	r.GET("/", OutputHTML)
	r.POST("/user", Login)
	r.DELETE("/user", Logout)

	return r

}
