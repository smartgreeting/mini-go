/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:14:45
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 22:17:13
 * @Email: 17719495105@163.com
 */
package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/svc"
	"gorm.io/gorm"
)

type UserService struct {
	svcCtx *svc.SvcContext
}

func NewUserService(svcCtx *svc.SvcContext) *UserService {
	return &UserService{
		svcCtx: svcCtx,
	}
}

func (u *UserService) GetUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := u.svcCtx.UserDao.FindUserInfoById(int64(id))
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "12",
			"data": res,
		})
	case gorm.ErrRecordNotFound:
		ctx.JSON(http.StatusOK, gin.H{
			"code":   200,
			"msg":    "未匹配到记录",
			"data":   map[string]interface{}{},
			"res":    []interface{}{map[string]interface{}{"a": 1}},
			"resArr": []interface{}{},
		})
	default:

	}

}
