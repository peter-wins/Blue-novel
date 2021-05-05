package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetExceptionWhereInfo()(filename string, line int, functionName string){
	pc, fileName, line, ok := runtime.Caller(3)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}
	return fileName, line, functionName
}