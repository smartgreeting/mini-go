/*
 * @Author: lihuan
 * @Date: 2024-09-13 20:03:55
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-14 23:26:24
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

// 定义Watermark结构体
type Watermark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"` // 注意JSON中的字段名为appid，Go中通常使用驼峰命名法，但这里为了与JSON对应，使用全小写加下划线
}

// 定义PhoneInfo结构体
type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     int       `json:"countryCode"`
	Watermark       Watermark `json:"watermark"`
}

// 定义最外层的结构体
type GetUserPhoneNumberReplay struct {
	ErrCode   int       `json:"errcode"`
	ErrMsg    string    `json:"errmsg"`
	PhoneInfo PhoneInfo `json:"phone_info"`
}
