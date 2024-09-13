/*
 * @Author: lihuan
 * @Date: 2024-09-13 22:03:05
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-13 22:21:20
 * @Email: 17719495105@163.com
 */
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WXHttpGetOptions struct {
	Url   string
	Reply interface{}
}

func WXHttpGet(ctx *gin.Context, opt WXHttpGetOptions) ([]byte, error) {
	resp, err := http.Get(opt.Url)
	if err != nil {
		ErrorResponse(ctx, "调用微信接口异常")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorResponse(ctx, "读取响应失败")
		return nil, err
	}

	if err := json.Unmarshal(body, opt.Reply); err != nil {
		ErrorResponse(ctx, "解析响应失败")
		return nil, err
	}
	return body, err
}
