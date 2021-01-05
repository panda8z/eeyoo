package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/utils"
	"github.com/panda8z/eeyoo/utils/errors"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errors.ERROR
	}
	return token, errors.SUCCESS

}

// 验证token

func CheckToken(token string) (*MyClaims, int) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.ERROR_TOKEN_OUTTIME
			} else {
				return nil, errors.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, errors.SUCCESS
		} else {
			return nil, errors.ERROR_TOKEN_WRONG
		}
	}
	return nil, errors.ERROR_TOKEN_WRONG
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errors.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errors.Msg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errors.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errors.Msg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errors.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errors.Msg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode != errors.SUCCESS {
			code = tCode
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errors.Msg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key)
		c.Next()
	}
}
