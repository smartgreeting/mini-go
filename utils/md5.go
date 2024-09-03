/*
 * @Author: lihuan
 * @Date: 2024-08-27 20:42:11
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-02 20:33:10
 * @Email: 17719495105@163.com
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// s:加密字符串  sum:密钥
func EncodeMd5(s string, sum []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(sum))
}
