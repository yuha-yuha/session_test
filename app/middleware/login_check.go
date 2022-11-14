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
		idString, _ := Id.(string)
		_, err := database.GetUserData(idString)

		if err != nil {
			ctx.Status(401)
			ctx.Abort()
		}
	} else {
		ctx.Abort()
	}

	ctx.Next()
}
