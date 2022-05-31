package conf

import (
	"accelerator/entity/db"
	"accelerator/util"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	PID           string
	AppID         string
	AppName       string
	AppPublicKey  string
	AppPrivateKey string
	AliPublicKey  string
	Pro           bool
	NotifyUrl     string
	DSN           string
}

var PayConf *Conf

// Init 初始化配置项
func (c *Conf) Init() {
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

	// 支付相关的配置
	PayConf = &Conf{
		AppID:         os.Getenv("APP_ID"),
		AppName:       os.Getenv("APP_NAME"),
		AppPublicKey:  os.Getenv("APP_PUBLIC_KEY"),
		AppPrivateKey: os.Getenv("APP_PRIVATE_KEY"),
		AliPublicKey:  os.Getenv("ALI_PUBLIC_KEY"),
		PID:           os.Getenv("PID"),
		Pro:           os.Getenv("PRO") == "true",
		NotifyUrl:     os.Getenv("Notify_url"),
		DSN:           os.Getenv("MYSQL_DSN"),
	}

	// 预留redis，如果需要使用缓存，打开这行注释
	// cache.Redis()
}
