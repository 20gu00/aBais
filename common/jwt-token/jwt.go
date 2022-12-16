package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var JWTToken jwtToken

type jwtToken struct{}

// token中包含的自定义信息以及jwt签名信息
type CustomClaims struct {
	UserName           string `json:"username"`
	Password           string `json:"password"`
	jwt.StandardClaims        // 官方字段
}

// 加解密因子
const (
	SECRET = "adoodevops"
)

//解析token
func (*jwtToken) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	//使用jwt.ParseWithClaims方法解析token，前端传来token,获得一个*Token类型的对象
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		zap.L().Error("解析token失败", zap.Error(err))
		// 处理token解析之后的各种错误
		// 解析token token有错
		if ve, ok := err.(*jwt.ValidationError); ok {
			// token格式错误(头部header 负载payload 签名signature)
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { // token过期
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { // 未生效
				return nil, errors.New("TokenNotValidYet")
			} else {
				return nil, errors.New("TokenInvalid") // 无效
			}
		}
	}

	// 转换成*CustomClaims类型并返回
	// token.Valid 1:token有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("解析Token失败")
}

// 生成JWT
//func GenToken(userID int64, username string) (string, error) {
//	// 创建一个我们自己的声明的数据
//	c := MyClaims{
//		userID,
//		username, // 自定义字段
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(
//				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Minute).Unix(), // 过期时间
//			Issuer: "cjq",
//		},
//	}
//	// 使用指定的签名方法创建签名对象(加密算法,token配置)
//	//header payload
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	// 签名并获得完整的编码后的字符串token
//	//signature
//	return token.SignedString(mySecret)
//	//最终的base64编码
//}
