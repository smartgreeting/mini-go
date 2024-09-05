/*
 * @Author: lihuan
 * @Date: 2024-09-04 20:42:11
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-05 21:17:25
 * @Email: 17719495105@163.com
 */

package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/mini-go/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var code int

		code = utils.Success

		Authorization := ctx.GetHeader("Authorization")

		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = utils.ErrInvalidToken
		} else {
			claims, err := utils.ParseToken(token[1], []byte(utils.Cfg.Token.Secret))
			if err != nil {

				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = utils.ErrTokenTimeout
				default:
					code = utils.ErrTokenParse
				}
			} else {
				id := claims.ID
				ctx.Set("userId", id)
			}
		}

		if code != utils.Success {
			utils.ErrorResponse(ctx, code)

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
