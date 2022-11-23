/*
 * @Author: lisheng
 * @Date: 2022-10-10 23:46:08
 * @LastEditTime: 2022-11-04 23:08:46
 * @LastEditors: lisheng
 * @Description: mysql驱动管理，支持管理多个Mysql数据库连接
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/database/mysql/mysql.go
 */
package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MysqlManager struct {
	EngineMap map[string]*gorm.DB
}

var (
	GMysqlManager *MysqlManager = new(MysqlManager)
)

/**
 * @description: 创建数据库引擎
 * @param {string} user
 * @param {string} pwd
 * @param {string} addr
 * @param {int} port
 * @param {string} db
 * @return {*}
 * @author: liqiyuWorks
 */
func CreateDBEngnine(user string, pwd string, addr string, port int, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	// ----------------------- 连接数据库 -----------------------
	var err error
	var mysqlCli *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s", user, pwd, addr, port, db)

	mysqlCli, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		CreateBatchSize:        1000,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		base.Glog.Errorf("%s, reason:%v", statuscode.ERROR_MYSQL_CONNECT.Msg, err)
	}
	if mysqlCli.Error != nil {
		base.Glog.Errorf("%s, reason:%v", statuscode.ERROR_MYSQL_CONNECT.Msg, mysqlCli.Error)
	}

	// // ----------------------- 连接池设置 -----------------------
	sqlDB, err := mysqlCli.DB()

	// // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err == nil {
		base.Glog.Infof(">>> Connected to Mysql")
	} else {
		base.Glog.Errorf(">>> %s", statuscode.ERROR_MYSQL_CONNECT.Msg)
	}

	return mysqlCli
}

/**
 * @description: 初始化mysqlManager
 * @return {*}
 * @author: liqiyuWorks
 */
func InitMysqlManager() func() error {
	GMysqlManager.EngineMap = make(map[string]*gorm.DB)
	for k, v := range config.C.Mysql {
		engine := CreateDBEngnine(v.User, v.Pwd, v.Addr, v.Port, v.DB)
		GMysqlManager.EngineMap[k] = engine
	}
	return func() error {
		for k, engine := range GMysqlManager.EngineMap {
			sqlDB, _ := engine.DB()
			err := sqlDB.Close()
			if err != nil {
				base.Glog.Errorf("> engine=%s, %s\n", k, statuscode.ERROR_MYSQL_CLOSE.Msg)
			} else {
				fmt.Printf("> mysqlName= %s close ok!\n", k)
			}
		}
		return nil
	}
}

/**
 * @description: mysql 自动注册表
 * @param {string} engineName
 * @param {interface{}} tableObj
 * @return {*}
 * @author: liqiyuWorks
 */
func RegisterTable(engineName string, tableObj interface{}) {
	if engineName == "" {
		engineName = "default"
	}
	if err := GMysqlManager.EngineMap[engineName].AutoMigrate(tableObj); err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MYSQL_AUTO_MIGRATE.Msg, err)
	}
}
