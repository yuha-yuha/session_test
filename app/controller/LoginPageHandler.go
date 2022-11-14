package controller

import (
	"github.com/gin-gonic/gin"
)

func LoginPage(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{})
}
