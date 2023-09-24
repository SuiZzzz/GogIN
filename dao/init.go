package dao

import (
	"GoGin/conf"
	"GoGin/dao/model"
	_ "GoGin/log"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var _db *gorm.DB

func init() {
	config := conf.Conf.Mysql
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	DataBaseLog(dsn)
}

func DataBaseLog(dsn string) {
	var dbLogger logger.Interface
	if gin.Mode() == gin.DebugMode {
		dbLogger = logger.Default.LogMode(logger.Info)
	} else {
		dbLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                           dsn,
		DefaultStringSize:             256,
		DisableDatetimePrecision:      true,
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         dbLogger,
	})

	if err != nil {
		panic(err)
	}

	_db = db

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 自动建表
	var user model.User
	exist := _db.Migrator().HasTable(&user)
	if !exist {
		err := _db.AutoMigrate(&user)
		if err != nil {
			panic(err)
		}
	}
}

func NewSession(ctx context.Context) *gorm.DB {
	db := _db
	return db.Session(&gorm.Session{Context: ctx})
}
