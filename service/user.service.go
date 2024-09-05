/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:14:45
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-05 21:06:15
 * @Email: 17719495105@163.com
 */
package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/svc"
	"github.com/smartgreeting/mini-go/utils"
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
		utils.SuccessResponse(ctx, res)
	case gorm.ErrRecordNotFound:
		utils.ErrorResponse(ctx, utils.ErrRecordNotFound)

	default:

	}

}
