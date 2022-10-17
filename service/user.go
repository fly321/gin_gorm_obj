package service

import (
	"demoProject/helper"
	"demoProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// UserLogin
// @Summary 用户登录
// @Tags 公共方法
// @Description 用户登录
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} string "OK"
// @Router /userLogin [post]
func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "username or password can not be empty",
		})
		return
	}

	data := new(models.UserBasic)
	err := models.Db.Where("username = ?", username).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "user not found",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error: " + err.Error(),
		})
		return
	}

	if data.Password != helper.GetMd5(password) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "password error",
		})
		return
	}
	token, err := helper.GenerateToken(username, data.Identity, data.IsAdmin)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error: " + err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": map[string]interface{}{
			"token":         token,
			"userModelData": data,
		},
	})
}
