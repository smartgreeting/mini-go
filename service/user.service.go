/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:14:45
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-22 21:53:43
 * @Email: 17719495105@163.com
 */
package service

import (
	"fmt"
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
func (u *UserService) DelUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	err := u.svcCtx.UserDao.DelById(int64(id))
	fmt.Println(err)
	if err != nil {
		utils.ErrorResponse(ctx, "删除失败")
		return
	}
	utils.SuccessResponse(ctx, nil)

}

func (u *UserService) GetTokenByOpenId(ctx *gin.Context) {
	id := ctx.Query("openid")
	cfg := u.svcCtx.Config

	token, err := utils.GenerateToken(id, []byte(cfg.Token.Secret), cfg.Token.ExpireTime)
	if err != nil {
		utils.ErrorResponse(ctx, "生成token失败")
		return
	}
	utils.SuccessResponse(ctx, token)
}
