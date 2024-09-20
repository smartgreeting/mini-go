/*
 * @Author: lihuan
 * @Date: 2024-08-31 20:00:17
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-20 20:44:17
 * @Email: 17719495105@163.com
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/middleware"
	"github.com/smartgreeting/mini-go/service"
	"github.com/smartgreeting/mini-go/svc"
)

func SetupRouter(svcCtx *svc.SvcContext) *gin.Engine {

	r := gin.Default()
	r.Use(middleware.LoggerToFile(svcCtx))
	userService := service.NewUserService(svcCtx)
	wxService := service.NewWXService(svcCtx)
	v1 := r.Group("/v1")
	v1.GET("/getUserInfo", userService.GetUserInfo)
	v1.DELETE("/delUserInfo", userService.DelUserInfo)
	v1.GET("/getOpenId", wxService.GetOpenIDByCode)
	v1.GET("/getPhoneNumber", wxService.GetPhoneNumber)
	return r
}
