package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gitee.com/panda8xy/gin-blog/utils"
	"gitee.com/panda8xy/gin-blog/utils/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// error code
var code int

// JwtKey define the secret key
var JwtKey = []byte(utils.JwtKey)

// KeyClaim define the claim
type KeyClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate a token
func GenerateToken(username, password string) (string, int) {

	claim := KeyClaim{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
			Issuer:    "ginblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errors.ERROR
	}
	return tokenStr, errors.SUCCESS
}

// Checktoken vertify the token
func Checktoken(tokenStr string) (*KeyClaim, int) {
	token, err := jwt.ParseWithClaims(tokenStr, &KeyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.ERROR
	}

	if key, ok := token.Claims.(*KeyClaim); ok && token.Valid {
		return key, errors.SUCCESS
	}

	return nil, errors.ERROR

}

// Jwt middleware for gin
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			code = errors.ERROR_TOKEN_NOT_EXIST
			c.Abort()
		}

		checkStr := strings.SplitN(tokenStr, " ", 2)

		if len(checkStr) != 2 && checkStr[0] == "Beare" {
			code = errors.ERROR_TOKEN_TYPE_WRONG
			c.Abort()
		}
		token, Ecode := Checktoken(checkStr[1])

		if Ecode == errors.ERROR {
			code = errors.ERROR_TOKEN_WRONG
			c.Abort()
		}

		if time.Now().Unix() > token.ExpiresAt {
			code = errors.ERROR_TOKEN_OUTTIME
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errors.Msg(code),
		})
		c.Set("username", token.Username)
		c.Next()
	}
}
