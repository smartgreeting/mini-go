/*
 * @Author: lihuan
 * @Date: 2024-08-31 20:00:17
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 21:55:37
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
	u := service.NewUserService(svcCtx)

	r.GET("/getUserInfo", u.GetUserInfo)
	return r
}
