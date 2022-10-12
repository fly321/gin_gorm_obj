package service

import (
	"demoProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail
// @Summary 查询用户
// @Tags 公共方法
// @Description 查询用户 identity
// @Param identity query string false "identity"
// @Success 200 {string} string "OK"
// @Router /userDetail [get]
func GetUserDetail(ctx *gin.Context) {
	identity := ctx.Query("identity")
	if identity == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "identity can not be empty",
		})
		return
	}

	data := new(models.UserBasic)

	err := models.Db.Omit("password").Where("identity = ?", identity).First(&data).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": data,
	})

}
