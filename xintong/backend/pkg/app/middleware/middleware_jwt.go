package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {


	return func(c *gin.Context) {
		whiteList :=[]string{
			"/auth/login",
			"/auth/github-login",
			"/upload/upload",
		}
		var msg string
		var data interface{}
		path :=c.Request.URL.Path;
		for i :=0;i< len(whiteList);i++{
			 if whiteList[i] ==path {
			 	c.Next();
			 	return;
			 }
		}
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			msg = "user not login"
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				msg = "user not login"
			} else if time.Now().Unix() > claims.ExpiresAt {
				msg = "user not login timeout "
			}

			c.Set("claims",claims)
		}
		if len(msg) > 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  msg,
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
