package client

import (
	"fmt"
	"movielibrary/config"

	// "time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func ConnectMySQL(dbConf config.MysqlConfig) *gorm.DB {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Schema)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Schema)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		// TODO: replace panic
		panic(err)
	}

	// db, err := gormDB.DB()
	// if err != nil {
	// 	// TODO: replace panic
	// 	panic(err)
	// }
	// db.SetMaxIdleConns(conf.MaxIdleConn)
	// db.SetMaxOpenConns(conf.MaxOpenConn)
	// db.SetConnMaxLifetime(conf.MaxConnLifetime * time.Second)

	dbClient = gormDB

	return dbClient
}

func DB() *gorm.DB {
	return dbClient
}
