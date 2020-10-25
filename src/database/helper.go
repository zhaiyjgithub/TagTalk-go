package database

import (
	"fmt"
	"github.com/zhaiyjgithub/TagTalk-go/src/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	mysqlDBOnce sync.Once
	mysqlEngine *gorm.DB
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
