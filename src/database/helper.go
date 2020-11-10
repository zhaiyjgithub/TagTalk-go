package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zhaiyjgithub/TagTalk-go/src/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	mysqlDBOnce sync.Once
	mysqlEngine *gorm.DB

	redisDBOnce sync.Once
	redisEngine *redis.Client
)

func InstanceMysqlDB() *gorm.DB  {
	mysqlDBOnce.Do(func() {
		var err error
		c := conf.MySQLConf
		driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.DBName)

		mysqlEngine, err = gorm.Open(mysql.Open(driveSource), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})

	return mysqlEngine
}

func InstanceRedisDB() *redis.Client {
	redisDBOnce.Do(func() {
		c := conf.RedisConf
		addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
		fmt.Printf("Connect to redis: %s", addr)
		redisEngine = redis.NewClient(&redis.Options{
			Addr: addr,
			Password: "",
			DB: 0,
		})
	})

	return  redisEngine
}