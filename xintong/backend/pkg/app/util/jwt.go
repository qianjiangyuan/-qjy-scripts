package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("adfsdfsaddfsdfsdafdsfsafdsafiwrwesfnsdafiaefdsnfsdnfssdfs")

type Claims struct {
	user.User
	jwt.StandardClaims
}

func GenerateToken(user user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
