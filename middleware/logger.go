package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/smartgreeting/mini-go/svc"
)

// 日志记录到文件
func LoggerToFile(svcCtx *svc.SvcContext) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		ctx.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		duration := endTime.Sub(startTime)

		// 请求方式
		method := ctx.Request.Method

		// 请求路由
		url := ctx.Request.RequestURI

		// 状态码
		status := ctx.Writer.Status()

		// 请求IP
		ip := ctx.ClientIP()

		// 日志格式
		svcCtx.Logger.WithFields(logrus.Fields{
			"status":   status,
			"duration": duration,
			"ip":       ip,
			"method":   method,
			"url":      url,
		}).Info("log中间件")
	}
}
