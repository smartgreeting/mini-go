/*
 * @Author: lihuan
 * @Date: 2024-09-13 20:04:30
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-16 20:37:32
 * @Email: 17719495105@163.com
 */
package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/constants"
	"github.com/smartgreeting/mini-go/models"
	"github.com/smartgreeting/mini-go/svc"
	"github.com/smartgreeting/mini-go/utils"
)

type WXService struct {
	svcCtx *svc.SvcContext
}

func NewWXService(svcCtx *svc.SvcContext) *WXService {
	return &WXService{
		svcCtx: svcCtx,
	}
}

// 获取 openid、unionid 和 session_key
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
func (w *WXService) GetOpenIDByCode(ctx *gin.Context) {

	code := ctx.Query("code")
	wx := w.svcCtx.Config.WX
	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", constants.GetOpenIDUrl, wx.AppID, wx.AppSecret, code)

	var openIdReplay models.OpenIdReplay

	utils.HttpHandle(ctx, utils.HttpHandleOptions{
		Url:   url,
		Reply: &openIdReplay,
	})

	if openIdReplay.Errcode != 0 {
		str := fmt.Sprintf("%d:%s", openIdReplay.Errcode, openIdReplay.Errmsg)
		utils.ErrorResponse(ctx, str)
		return
	}
	utils.SuccessResponse(ctx, openIdReplay)

}

// 获取 access_token GetAccessToken
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html

func (w *WXService) GetAccessToken(ctx *gin.Context) {

	res, _ := GetCacheToken(ctx, w.svcCtx)

	utils.SuccessResponse(ctx, res)
}

// 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (w *WXService) GetPhoneNumber(ctx *gin.Context) {

	res, _ := GetCacheToken(ctx, w.svcCtx)
	url := fmt.Sprintf("%s?access_token=%s", constants.GetUserPhoneNumber, res.AccessToken)
	var phoneReply models.GetUserPhoneNumberReplay
	utils.HttpHandle(ctx, utils.HttpHandleOptions{
		Url:   url,
		Reply: &phoneReply,
	})
	fmt.Println(phoneReply)
}

const (
	AccessTokenKey = "__Access_Token_Redis_Key__"
	Times          = 90 //缓存时间 90 分钟
)

func GetCacheToken(ctx *gin.Context, svcCtx *svc.SvcContext) (models.AccessTokenReplay, error) {
	// 从缓存中获取
	val, _ := svcCtx.RedisDB.Get(AccessTokenKey).Result()
	var accessTokenReply models.AccessTokenReplay
	json.Unmarshal([]byte(val), &accessTokenReply)

	if len(val) != 0 {
		ttl, _ := svcCtx.RedisDB.TTL(AccessTokenKey).Result()
		return models.AccessTokenReplay{
			AccessToken: accessTokenReply.AccessToken,
			ExpiresIn:   int(ttl) / int(time.Second),
		}, nil
	}
	wx := svcCtx.Config.WX
	url := fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=client_credential", constants.GetAccessToken, wx.AppID, wx.AppSecret)

	var accessTokenReplay models.AccessTokenReplay

	body, _ := utils.HttpHandle(ctx, utils.HttpHandleOptions{
		Url:   url,
		Reply: &accessTokenReplay,
	})
	svcCtx.RedisDB.Set(AccessTokenKey, body, Times*time.Minute).Err()
	return accessTokenReplay, nil
}
