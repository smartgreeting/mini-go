/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:15:52
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-05 21:19:27
 * @Email: 17719495105@163.com
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success           = 2000
	ErrParamsParse    = 2001
	ErrInvalidToken   = 2003
	ErrTokenTimeout   = 2004
	ErrTokenParse     = 2005
	ErrPhoneNotExit   = 2006
	ErrRecordNotFound = 2007
	ErrCustomtMsg     = 2008
)

// 列举通用错误
var resMap = map[int]string{
	Success:           "成功",
	ErrParamsParse:    "参数格式错误",
	ErrInvalidToken:   "无效的token",
	ErrTokenTimeout:   "token过期",
	ErrTokenParse:     "token解析失败",
	ErrPhoneNotExit:   "手机号码不正确",
	ErrRecordNotFound: "未匹配到记录",
}

func ErrorResponse(ctx *gin.Context, value ...interface{}) {
	if len(value) > 0 {
		r := value[0]

		var data interface{}

		if len(value) == 1 {
			data = gin.H{}
		}
		if len(value) == 2 {
			data = value[1]
		}

		switch v := r.(type) {

		case int:
			getJson(ctx, v, GetMsg(v), data)

		case string:
			getJson(ctx, ErrCustomtMsg, r.(string), data)
		default:
			getJson(ctx, 5000, "服务异常", data)
		}
	}

}

func getJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func SuccessResponse(ctx *gin.Context, value ...interface{}) {
	if len(value) > 0 {

		var msg interface{}

		if len(value) == 1 {
			msg = GetMsg(Success)
		}
		if len(value) == 2 {
			msg = value[1]
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": Success,
			"msg":  msg,
			"data": value[0],
		})
	}

}

func GetMsg(code int) string {

	str, ok := resMap[code]

	if ok {
		return str
	}
	return "未知错误"
}
