package controller

import (
	"Golang/session/app/database"
	"Golang/session/app/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func OutputHTML(ctx *gin.Context) {
	session := sessions.Default(ctx)

	userId := session.Get("Id").(int)

	isloggedIn := model.IsLoggedIn(userId)

	user := database.GetUserData(userId)

	ctx.HTML(200, "index.html", gin.H{
		"isLoggedIn": isloggedIn,
		"UserName":   database.user.UName,
		"UserId":     database.user.Uid,
	})

}
