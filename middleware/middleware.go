/*
 * @Author: lisheng
 * @Date: 2022-10-11 00:12:42
 * @LastEditTime: 2022-11-02 11:08:10
 * @LastEditors: lisheng
 * @Description: 通用中间件
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/middleware/middleware.go
 */
package middleware

import (
	"net/http"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

// 日志记录到文件 在哪调用
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 日志格式
		base.Glog.Infof("| %3d | %13v | %15s | %s | %s |",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}
