package controller

import (
	"app/database"
	"app/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func OutputHTML(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := database.User{}

	Id := session.Get("Id")

	isloggedIn, Idint := model.IsLoggedIn(Id)

	if isloggedIn {
		user = database.GetUserData(Idint)
	}

	ctx.HTML(200, "index.html", gin.H{
		"isLoggedIn": isloggedIn,
		"UserName":   user.UName,
		"UserId":     user.Uid,
	})

}
