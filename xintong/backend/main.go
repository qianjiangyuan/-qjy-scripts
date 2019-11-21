package main

import (
	"fmt"
	"github.com/spf13/viper"
	App "github.com/xxmyjk/xintong/backend/pkg/app"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/middleware"
	"os"
)

const (
	DEFAULT_APP_CONF_PATH = "config/app.conf.toml"
	DEFAULT_APP_CONF_TYPE = "toml"
)

func main() {
	// TODO: from flag
	confPath := "config/app.conf.toml"

	conf := loadConf(confPath, "toml")
	connect.Init(conf)

	app := App.New(conf)
	app.Use(middleware.Stub())

	err := app.Run(fmt.Sprintf("%s:%d",
		conf.Get("http.host"),
		conf.Get("http.port"),
	))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func loadConf(appConf ...string) *viper.Viper {
	var (
		length   int
		confPath string
		confType string
	)

	if length < 1 {
		confPath = DEFAULT_APP_CONF_PATH
		confType = DEFAULT_APP_CONF_TYPE
	} else if length < 2 {
		confPath = appConf[0]
		confType = DEFAULT_APP_CONF_TYPE
	} else {
		confPath = appConf[0]
		confType = appConf[1]
	}

	fd, err := os.Open(confPath)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	conf := viper.New()
	conf.SetConfigType(confType)

	err = conf.ReadConfig(fd)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return conf
}
