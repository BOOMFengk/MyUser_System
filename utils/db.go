package utils

import (
	"MyUser_System/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// openDB连接db
func openDB() {
	mysqlConf := config.GetGlobalConf().DbConfig
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	log.Info("mdb addr:" + connArgs)

	var err error
	gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch db connection err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConn)
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxIdleTime * int64(time.Second)))

}

func GetDB() *gorm.DB {
	dbOnce.Do(openDB)
	return db
}
