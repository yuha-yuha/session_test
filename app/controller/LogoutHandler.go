package controller

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()

	log.Println("logout")

	ctx.Redirect(302, "/login")
}
