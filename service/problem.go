package service

import (
	"demoProject/define"
	"demoProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Summary ping example
// @Description 获取列表
// @Tags List
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param category_identity query string false "category_id"
// @Success 200 {string} string "OK"
// @Router /problemList [get]
func GetProblemList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(ctx.DefaultQuery("limit", define.DefaultSize))
	page = (page - 1) * size
	var count int64
	keyword := ctx.Query("keyword")
	categoryIdentity := ctx.Query("category_identity")
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	err := tx.Debug().Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList error:", err)
	}
	ctx.JSON(
		200,
		gin.H{
			"code":    200,
			"message": "pong",
			"data": map[string]interface{}{
				"list":  list,
				"count": count,
			},
		},
	)
}

// GetProblemDetail
// @Summary 问题详情
// @Tags 公共方法
// @Description 问题详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} string "OK"
// @Router /problemList [get]
func GetProblemDetail(ctx *gin.Context) {
	identity := ctx.Query("identity")
	if identity == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "identity is empty",
		})
		return
	}
	problemBasic := new(models.ProblemBasic)
	err := models.Db.Where("identity = ?", identity).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").First(&problemBasic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "The current record does not exist",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    problemBasic,
	})
}
