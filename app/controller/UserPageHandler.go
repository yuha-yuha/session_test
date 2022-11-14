package controller

import (
	"app/database"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserPage(ctx *gin.Context) {
	session := sessions.Default(ctx)

	Uid := session.Get("Id")
	user := database.User{}

	StrUid, _ := Uid.(string)
	user, _ = database.GetUserData(StrUid)

	log.Println("html")

	ctx.HTML(200, "user.html", gin.H{
		"UserName": user.UName,
		"UserId":   user.Uid,
	})

	log.Println("html")
}
