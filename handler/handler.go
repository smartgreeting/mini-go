/*
 * @Author: lihuan
 * @Date: 2024-08-31 20:00:17
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-13 20:24:59
 * @Email: 17719495105@163.com
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/service"
	"github.com/smartgreeting/mini-go/svc"
)

func SetupRouter(svcCtx *svc.SvcContext) *gin.Engine {

	r := gin.Default()
	userService := service.NewUserService(svcCtx)
	wxService := service.NewWXService(svcCtx)

	r.GET("/getUserInfo", userService.GetUserInfo)
	r.DELETE("/delUserInfo", userService.DelUserInfo)
	r.GET("/getOpenId", wxService.GetOpenIDByCode)
	r.GET("/getAccessToken", wxService.GetAccessToken)
	return r
}
