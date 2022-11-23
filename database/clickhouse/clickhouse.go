/*
 * @Author: lisheng
 * @Date: 2022-10-13 09:39:39
 * @LastEditTime: 2022-11-23 14:20:15
 * @LastEditors: lisheng
 * @Description: clickhouse 驱动
 * @FilePath: /jf-go-kit/database/clickhouse/clickhouse.go
 */
package clickhouse

import (
	"fmt"
	"jf-go-kit/config"
	"log"
	"os"
	"time"

	"jf-go-kit/base/statuscode"

	"jf-go-kit/base"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type CkManager struct {
	EngineMap map[string]*gorm.DB
}

var (
	GCkManager *CkManager = new(CkManager)
)

/**
 * @description: CreateDBEngnine
 * @param {string} user
 * @param {string} pwd
 * @param {string} addr
 * @param {int} port
 * @param {string} db
 * @return {*}
 * @author: liqiyuWorks
 */
func CreateDBEngnine(user string, pwd string, addr string, port int, db string) *gorm.DB {
	var err error
	var ckCli *gorm.DB
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	dsn := fmt.Sprintf("http://%s:%d?database=%s&username=%s&password=%s", addr, port, db, user, pwd)
	ckCli, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}}) // 使用单数表名

	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_CK_CONNECT.Msg, err)
	}
	if ckCli.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_CK_CONNECT.Msg, err)
	}

	// ----------------------- 连接池设置 -----------------------
	sqlDB, err := ckCli.DB()

	// // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err == nil {
		base.Glog.Infof(">>> Connected to Clickhouse")
	} else {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_CK_CLOSE.Msg, err)
	}
	return ckCli
}

/**
 * @description: 初始化CK
 * @return {*}
 * @author: liqiyuWorks
 */
func InitCkManager() func() error {
	GCkManager.EngineMap = make(map[string]*gorm.DB)
	for k, v := range config.C.Clickhouse {
		engine := CreateDBEngnine(v.User, v.Pwd, v.Addr, v.Port, v.DB)
		GCkManager.EngineMap[k] = engine
	}

	return func() error {
		for k, engine := range GCkManager.EngineMap {
			sqlDB, _ := engine.DB()
			err := sqlDB.Close()
			if err != nil {
				base.Glog.Errorf("> clickhouseName= %s: %s\n", k, statuscode.ERROR_CK_CLOSE.Msg)
			} else {
				fmt.Printf("> clickhouseName= %s close ok!\n", k)
			}
		}
		return nil
	}
}

/**
 * @description: CK AutoMigrate
 * @param {string} engineName
 * @param {interface{}} tableObj
 * @return {*}
 * @author: liqiyuWorks
 */
func RegisterTable(engineName string, tableObj interface{}) {
	if err := GCkManager.EngineMap[engineName].AutoMigrate(tableObj); err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_CK_AUTO_MIGRATE.Msg, err)
	}
}
