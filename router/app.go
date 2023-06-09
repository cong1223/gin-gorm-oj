package router

import (
	"gin_orm_oj/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 路由规则
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)
	r.GET("/user-detail", service.GetUserDetail)
	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)
	return r
}
