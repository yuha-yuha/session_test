package middleware

import (
	"app/database"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(ctx *gin.Context) {
	session := sessions.Default(ctx)

	Id := session.Get("Id")
	log.Println("Id:", Id)

	if Id != nil {
		idString, _ := Id.(string)
		_, err := database.GetUserData(idString)

		if err != nil {
			log.Println(err)
			ctx.Status(401)
			ctx.Redirect(302, "/login")
			ctx.Abort()
		}
	} else {
		ctx.Status(401)
		ctx.Redirect(302, "/login")
		ctx.Abort()
	}
	log.Println("認証成功")
	ctx.Next()
}
