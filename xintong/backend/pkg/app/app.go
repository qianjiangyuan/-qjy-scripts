package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xxmyjk/xintong/backend/pkg/app/middleware"
)

func New(conf *viper.Viper) *gin.Engine {
	app := gin.Default()

	// cors here
	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = conf.GetStringSlice("cors.AllowOrigins")
	corsConf.AllowMethods = conf.GetStringSlice("cors.AllowMethods")
	corsConf.AllowHeaders = conf.GetStringSlice("cors.AllowHeaders")
	corsConf.AllowCredentials = conf.GetBool("cors.AllowCredentials")

	app.Use(cors.New(corsConf))
	app.Use(middleware.Stub())
	app.Use(middleware.JWT())
	// middleware shall be Used before router
	RegisterRouter(app)
	return app
}
