package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//	"net/http"
)

func Stub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("stub middleware echo ...")
		names := ctx.HandlerNames()

		for _, n := range names {
			fmt.Println("\t", n)
		}
	}
}
