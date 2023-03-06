/*
 * @Author: lisheng
 * @Date: 2022-10-18 15:04:49
 * @LastEditTime: 2023-01-06 14:16:12
 * @LastEditors: lisheng
 * @Description: PostareSQL驱动
 * @FilePath: /go-kit/database/pg/pg.go
 */
package pg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/liqiyuworks/go-kit/config"

	"gitee.com/liqiyuworks/go-kit/common/statuscode"

	"gitee.com/liqiyuworks/go-kit/base"

	// mysql 数据库驱动
	"gorm.io/driver/postgres"
	"gorm.io/gorm" // 使用 gorm ，操作数据库的 orm 框架
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/*
go 访问权限：
变量名、函数名、常量名首字母大写，则可以被其他包访问,
如果首字母小写，则只能在本包中使用。
首字母大写为共有，首字母小写为私有。
*/

type PgManager struct {
	EngineMap map[string]*gorm.DB
}

var (
	GPgManager *PgManager = new(PgManager)
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
	var pgCli *gorm.DB
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", addr, user, pwd, db, port)
	pgCli, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_CONNECT.Msg, err)
	}
	if pgCli.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_CONNECT.Msg, pgCli.Error)
	}

	// ----------------------- 连接池设置 -----------------------
	sqlDB, err := pgCli.DB()

	// // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err == nil {
		base.Glog.Infof(">>> Connected to Pg")
	} else {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_CONNECT.Msg, pgCli.Error)
	}
	return pgCli
}

/**
 * @description: 初始化PG
 * @return {*}
 * @author: liqiyuWorks
 */
func InitPgManager() func() error {
	GPgManager.EngineMap = make(map[string]*gorm.DB)
	for k, v := range config.C.Pg {
		engine := CreateDBEngnine(v.User, v.Pwd, v.Addr, v.Port, v.DB)
		GPgManager.EngineMap[k] = engine
	}
	return func() error {
		for k, engine := range GPgManager.EngineMap {
			sqlDB, _ := engine.DB()
			err := sqlDB.Close()
			if err != nil {
				base.Glog.Errorf("> pgName= %s: %s \n", k, statuscode.ERROR_PG_CLOSE.Msg)
			} else {
				fmt.Printf("> pgName= %s close ok!\n", k)
			}
		}
		return nil
	}

}

/**
 * @description: RegisterTable
 * @param {string} engineName
 * @param {interface{}} tableObj
 * @return {*}
 * @author: liqiyuWorks
 */
func RegisterTable(engineName string, tableObj interface{}) {
	if engineName == "" {
		engineName = "default"
	}
	if err := GPgManager.EngineMap[engineName].AutoMigrate(tableObj); err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_AUTO_MIGRATE.Msg, err)
	}
}
