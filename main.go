package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/gohub/bootstrap"
	btsConfig "github.com/vinoMamba/gohub/config"
	"github.com/vinoMamba/gohub/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	app_name := config.Get("app.name")
	fmt.Println("app_name:", app_name)
	// boot router
	r := gin.New()
	bootstrap.SetupRouter(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
