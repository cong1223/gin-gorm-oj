package service

import (
	"gin_orm_oj/define"
	"gin_orm_oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Description do ping
// @Tags 公共方法
// @Param page query int false "请输入page"
// @Param size query int false "请输入page size"
// @Param keyword query int false "请输入keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList page parse error", err)
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList size parse error", err)
		return
	}

	startIndex := (page - 1) * size

	var count int64

	keyword := c.Query("keyword")

	list := make([]*models.ProblemBasic, 0)

	tx := models.GetProblemList(keyword)

	err = tx.Count(&count).Omit("content").Offset(startIndex).Limit(size).Find(&list).Error

	if err != nil {
		log.Println("GetProblemList error", err)
		return
	}

	//c.String(http.StatusOK, "Get Problem List")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
