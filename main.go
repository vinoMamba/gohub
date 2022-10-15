package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/gohub/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupRouter(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
