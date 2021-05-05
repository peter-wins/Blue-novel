package global

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var (
	Gin *gin.Engine
	Gorm *gorm.DB
	Validate *validator.Validate
	Translator ut.Translator
)