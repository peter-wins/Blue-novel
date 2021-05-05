package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/peter-wins/Blue-novel/response"
	"github.com/peter-wins/Blue-novel/utils"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

func GlobalLoggerMiddleware() gin.HandlerFunc {
	logFileName := os.Getenv("LOG_PATH") + "/" + os.Getenv("APP_ENV") + ".log"
	src, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	logerTool := logrus.New()
	logerTool.SetLevel(logrus.DebugLevel)
	logerTool.Out = src
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		logFileName+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(logFileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:	logWriter,
		logrus.FatalLevel: 	logWriter,
		logrus.DebugLevel: 	logWriter,
		logrus.WarnLevel: 	logWriter,
		logrus.ErrorLevel: 	logWriter,
		logrus.PanicLevel: 	logWriter,
	}

	logerTool.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	return func(c *gin.Context){
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		params := string(body)
		defer c.Request.Body.Close()

		// fileName, line, functionName := utils.GetExceptionWhereInfo()
		responseWrite := &response.BodyLogWriter{Body: &bytes.Buffer{},
			ResponseWriter: c.Writer}
		c.Writer = responseWrite
		c.Next()

		params = utils.TrimSpaceLine(params)
		responseJsonData := responseWrite.Body.String()

		var responseData response.ResponseData
		json.Unmarshal(responseWrite.Body.Bytes(), &responseData)

		// 日志格式
		logrusFields := logrus.Fields{
			"url":			c.Request.Host + c.Request.RequestURI,
			"params":		params,
			"method":		c.Request.Method,
			"statusCode": 	c.Writer.Status(),
			"ip":			c.ClientIP(),
			//"fileName":		fileName,
			//"line":			line,
			//"functionName":	functionName,
			// 返回的响应数据和 code,有大用
			"responseCode": responseData.Code,
			"responseData": responseJsonData,

		}
		if responseData.Code == response.Ok || responseData.Code == response.Failure {
			logerTool.WithFields(logrusFields).Infoln(responseData.Msg)

		}else {
			logerTool.WithFields(logrusFields).Errorln(responseData.Msg)
		}
	}
}
