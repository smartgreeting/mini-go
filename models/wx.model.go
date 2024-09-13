/*
 * @Author: lihuan
 * @Date: 2024-09-13 20:03:55
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-13 20:45:48
 * @Email: 17719495105@163.com
 */
package models

type OpenIdReplay struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

type AccessTokenReplay struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
