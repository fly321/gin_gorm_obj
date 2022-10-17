package middleware

import (
	"demoProject/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "auth can not be empty",
			})
			return
		}

		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "error: " + err.Error(),
			})
			return
		}

		if userClaim.IsAdmin != 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "no permission",
			})
			return
		}

		ctx.Next()
	}
}
