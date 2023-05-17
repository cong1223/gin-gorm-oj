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
// @Param keyword query string false "请输入keyword"
// @Param category_identity query string false "请输入category identity"
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
	keyword := c.Query("keyword")
	// 分类的唯一标识
	categoryIdentity := c.Query("category_identity")
	// 查询数据的起始索引位置
	startIndex := (page - 1) * size
	// 数据总条数
	var count int64
	// 数据列表
	list := make([]*models.ProblemBasic, 0)
	// DAO 查数据， 返回DB对象
	tx := models.GetProblemList(keyword, categoryIdentity)
	// 应用各种查询条件筛选值
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
