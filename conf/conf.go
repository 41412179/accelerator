package conf

import (
	"accelerator/entity/db"
	"accelerator/util"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load(".env.example")

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	// sql := os.Getenv("MYSQL_DSN")
	// fmt.Printf("sql=%s", sql)
	db.Database(os.Getenv("MYSQL_DSN"))

	// 预留redis，如果需要使用缓存，打开这行注释
	// cache.Redis()
}
