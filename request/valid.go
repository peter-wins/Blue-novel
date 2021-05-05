package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peter-wins/Blue-novel/global"
	"net/http"
)
func FormValidateException(c *gin.Context, err error){
	// 写法1，类似php数组 $a[0]
	errorMessage := err.(validator.ValidationErrors)[0].Translate(global.Translator)
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg": errorMessage,
})

}