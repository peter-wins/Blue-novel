package response

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// Failed 请求响应出现异常的结构体
type Failed struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
type Succeed struct {
	Code int `json:"code"`
}
type ResponseJson struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}
type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}
// logger中间件中使用
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
// CustomerException 自定义异常函数
func CustomerException(message string, c *gin.Context, code int) {
	exceptionMessage := viper.GetString(message)
	if exceptionMessage == "" {
		exceptionMessage = viper.GetString("The system started the car.")
	}
	c.JSON(http.StatusOK, Failed{
		Code: code,
		Msg: exceptionMessage,
	})
	// abort 感觉没啥用，但是为了以防万一还是加上吧
	c.Abort()
}
func Json (c *gin.Context) {
	c.JSON(http.StatusOK, Succeed{Ok})
}
func JsonData(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, ResponseJson{Ok, data})
}
// Write 读取响应数据
func (w BodyLogWriter ) Write (b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
// WriteString 读取响应数据
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
