package setting

import (
	"github.com/go-ini/ini"
	"time"
)

var (
	Cfg *ini.File // 配置文件

	RunMode string // 运行模式

	HTTPPort     int           // HTTP端口
	ReadTimeout  time.Duration // 读取超时时间
	WriteTimeout time.Duration // 写入超时时间

	PageSize  int    // 分页大小
	JwtSecret string // JWT密钥
)

// 初始化配置
func init() {

}

// 加载基本配置
func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// 加载服务器配置
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		panic("Fail to get section 'server': " + err.Error())
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = sec.Key("READ_TIMEOUT").MustDuration(60) * time.Second
	WriteTimeout = sec.Key("WRITE_TIMEOUT").MustDuration(60) * time.Second
}

// 加载应用配置
func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		panic("Fail to get section 'app': " + err.Error())
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT").MustString("!@)*#)!@U#@*!@!)")
}
