package middleware

import (
	"app/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(ctx *gin.Context) {
	session := sessions.Default(ctx)

	Id := session.Get("id")

	if Id != nil {
		idInt, _ := Id.(int)
		_, err := database.GetUserData(idInt)

		if err != nil {
			ctx.Status(401)
			ctx.Abort()
		}
	} else {
		ctx.Abort()
	}

	ctx.Next()
}
