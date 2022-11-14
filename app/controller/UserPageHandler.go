package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserPage(ctx *gin.Context) {
	session := sessions.Default(ctx)
}
