package service

import (
	"demoProject/define"
	"demoProject/modles"
	"github.com/gin-gonic/gin"
	"log"
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
	list := make([]*modles.ProblemBasic, 0)
	tx := modles.GetProblemList(keyword, categoryIdentity)
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
