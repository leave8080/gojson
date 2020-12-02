package conf

import (
	"gofer/pkg/orm"
)

var Conf = &Config{}



type Config struct {
	ConcurrentNumber    int  // Goroutine 最大数量
	DeviceDataLogSwitch bool // 日志开关
	MySQL               *orm.MySQLConfig
	Redis               *orm.RedisConfig
	Mongo               *orm.MongoConfig
}
