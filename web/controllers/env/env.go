package env

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
	"os"
	"strings"
)

// 获取程序运行环境
// 根据程序运行路径后缀判断
// 如果是test就测试环境
func IsTestEnv()bool  {
	files := os.Args
	for _,v := range files{
		if strings.Contains(v,"test"){return true}
	}
	return false
}

// 接口跨域处理
func CrsAuth() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"ACCEPT","Content-Type","Content-Length","Accept-Encoding","X-CSRF-Token","Authorization"},
		AllowCredentials:   true,

	})
}
