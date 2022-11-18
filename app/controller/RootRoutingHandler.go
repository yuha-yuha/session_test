package controller

import (
	"app/model"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RootRouting(ctx *gin.Context) {
	session := sessions.Default(ctx)

	LoginFlag, _ := model.IsLoggedIn(session.Get("Id"))

	log.Println("Id:", session.Get("Id"))
	log.Println("LoginFlag:", LoginFlag)
	if LoginFlag {
		ctx.Redirect(302, "/auth/user")
	} else {
		ctx.Redirect(302, "/login")
	}
}
