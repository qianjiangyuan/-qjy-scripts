package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PAGE = 0
	DEFAULT_SIZE = 20
)

func ErrCatch(err error, ctx *gin.Context) bool {
	if err != nil {
		fmt.Println("ErrHandler ...")
		fmt.Println(err)
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  "inner service error: " + err.Error(),
			"data": err.Error(),
		})

		ctx.Abort()
		return true
	}

	return false
}
