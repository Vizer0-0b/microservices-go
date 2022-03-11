package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var jwtkey = []byte("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDI2a2EJ7m872v0afyoSDJT2o1+SitIeJSWtLJU8/Wz2m7gStexajkeD+Lka6DSTy8gt9UwfgVQo6uKjVLG5Ex7PiGOODVqAEghBuS7JzIYU5RvI543nNDAPfnJsas96mSA7L/mD7RTE2drj6hf3oZjJpMPZUQI/B1Qjb5H3K3PNwIDAQAB")

type Claims struct {
	Owner string
	Uid string
	RoleId string
	RouterList string
	jwt.StandardClaims
}

func CheckAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	// validate token format
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Access deny!!!"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	print(claims.Uid)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Access deny!!!"})
		ctx.Abort()
		return
	}

	ctx.Next()
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
