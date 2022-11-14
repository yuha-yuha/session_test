package controller

import (
	"app/model"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	userId := ctx.PostForm("userId")
	userPass := ctx.PostForm("userPass")

	log.Println(userId, ":", userPass)

	MatchPass := model.PassCheck(userId, userPass)

	if MatchPass {
		log.Println("logged in!!!!!!!!!!!")
		session.Set("Id", userId)
		session.Save()
		ctx.Redirect(302, "/auth/user")
	} else {
		ctx.Redirect(302, "/login")
	}

}
