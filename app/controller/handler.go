package controller

import (
	"app/database"
	"app/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func OutputHTML(ctx *gin.Context) {
	session := sessions.Default(ctx)

	Id := session.Get("Id").(int)

	isloggedIn := model.IsLoggedIn(Id)

	user := database.GetUserData(Id)

	ctx.HTML(200, "index.html", gin.H{
		"isLoggedIn": isloggedIn,
		"UserName":   user.UName,
		"UserId":     user.Uid,
	})

}
