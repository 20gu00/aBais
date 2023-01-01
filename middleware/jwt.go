package middleware

import (
	"github.com/20gu00/aBais/common/jwt-token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 对登录接口放行
		loginReqUrl := "/api/v1/login"
		if len(c.Request.URL.String()) >= len(loginReqUrl) && c.Request.URL.String()[0:13] == loginReqUrl {
			c.Next()
		} else {
			// 携带Token有三种方式 1.放在请求头(header中自定义key value  token:xxx 2.放在请求体 3.放在URI
			// (authorization bear token Token)放在Header的Authorization中，并使用Bearer开头 Authorization: Bearer xxx  / X-TOKEN: xxx
			// 获取Header中的Authorization
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  "请求未携带token，无权限访问",
					"data": nil,
				})
				c.Abort()
				return
			}

			// parseToken 解析token包含的信息
			claims, err := jwt.JWTToken.ParseToken(token)
			if err != nil {
				// token延期错误
				if err.Error() == "TokenExpired" {
					c.JSON(http.StatusBadRequest, gin.H{
						"msg":  "授权已过期",
						"data": nil,
					})
					c.Abort()
					return
				}
				//  其他解析错误
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  err.Error(),
					"data": nil,
				})
				c.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set("claims", claims)

			c.Next()
		}
	}
}
