package dao

import (
	"gofer/pkg/orm"
	"time"
	"github.com/leave8080/gojson/conf"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	MySQL *gorm.DB
	Redis *orm.Redis
	Mongo *mongo.Database
	c     *conf.Config
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		MySQL: orm.InitGorm(c.MySQL),
		Redis: orm.InitRedis(c.Redis),
		Mongo: orm.InitMongo(c.Mongo).Database(c.Mongo.Database),
		c:     c,
	}
	return d
}

func (d *Dao) Close() {
	if d.MySQL != nil {
		d.MySQL.Close()
	}
	if d.Mongo != nil {
		orm.MongoClose(d.Mongo.Client(), time.Duration(d.c.Mongo.ConnectTimeOut)*time.Second)
	}
}
