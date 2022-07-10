package config

import (
	"os"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	//ServerHost 服务器ip
	ServerHost = os.Getenv("SERVER_HOST")
	//ServerPort 服务器端口
	ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	//SubServerHost 备用服务器ip
	SubServerHost = os.Getenv("SUB_SERVER_HOST")
	//SubServerPort 备用服务器端口
	SubServerPort, _ = strconv.Atoi(os.Getenv("SUB_SERVER_PORT"))
	//Debug 是否debug模式
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	//Logger 日志
	Logger *glog.Logger
)

func init() {
	if ServerPort == 0 {
		ServerPort = 2404
	}
	initLogger()
}

func initLogger() {
	Logger = glog.New()
	Logger.SetConfigWithMap(g.Map{
		"path": "./logs/iec104.log",
		"level": "info",
	})
}
