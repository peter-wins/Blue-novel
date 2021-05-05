package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/Blue-novel/response"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	router.POST("/test/customer-exception", func(c *gin.Context) {
		panic("xxxx")
		response.CustomerException(
			"Customer exception test.",
			c,
			response.Failure,
		)
		return
	})
	return router
}