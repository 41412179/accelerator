package main

import (
	"accelerator/conf"
	"accelerator/server"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	fmt.Println("GIN_MODE:", os.Getenv("GIN_MODE"))
	// 从配置文件读取配置
	new(conf.Conf).Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
