/*
 * @Author: lisheng
 * @Date: 2022-10-13 14:12:44
 * @LastEditTime: 2023-01-09 15:31:10
 * @LastEditors: lisheng
 * @Description: Tdengine驱动
 * @FilePath: /jf-go-kit/database/tdengine/tdengine.go
 */
package tdengine

import (
	"context"
	"fmt"
	"time"

	"gitee.com/wuxi_jiufang/jf-go-kit/config"

	"gitee.com/wuxi_jiufang/jf-go-kit/common/statuscode"

	"gitee.com/wuxi_jiufang/jf-go-kit/base"

	"gitee.com/chunanyong/zorm"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

type TDengineManager struct {
	EngineMap map[string]*zorm.DBDao
	CtxMap    map[string]*context.Context
}

var (
	GTDengineManager *TDengineManager = new(TDengineManager)
)

func CreateDBEngnine(user, pwd, addr string, port int) (*zorm.DBDao, *context.Context) {
	var readCtx = context.Background()
	var err error
	dsn := fmt.Sprintf("%s:%s@http(%s:%d)/", user, pwd, addr, port)
	dbDaoConfig := zorm.DataSourceConfig{
		//DSN 数据库的连接字符串
		DSN: dsn,
		//数据库驱动名称:mysql,postgres,oci8,sqlserver,sqlite3,clickhouse,dm,kingbase,aci 和Dialect对应,处理数据库有多个驱动
		//sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		DriverName: "taosRestful",
		//数据库方言:mysql,postgresql,oracle,mssql,sqlite,clickhouse,dm,kingbase,shentong 和 DriverName 对应,处理数据库有多个驱动
		Dialect: "tdengine",
		//SlowSQLMillis 慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出
		SlowSQLMillis: -1,
		//MaxOpenConns 数据库最大连接数 默认50
		MaxOpenConns: 10,
		//MaxIdleConns 数据库最大空闲连接数 默认50
		MaxIdleConns: 100,
		//ConnMaxLifetimeSecond 连接存活秒时间. 默认600(10分钟)后连接被销毁重建.避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		ConnMaxLifetimeSecond: 600,
		DisableTransaction:    true, // 禁用全局事务
	}
	tdClient, err := zorm.NewDBDao(&dbDaoConfig)
	readCtx, _ = tdClient.BindContextDBConnection(readCtx)
	if err == nil {
		base.Glog.Infof(">>> Connected to Tdengine")
	} else {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_CONNECT.Msg, err)
	}
	return tdClient, &readCtx

}

// 初始化连接
func InitTdClient() func() error {
	GTDengineManager.EngineMap = make(map[string]*zorm.DBDao)
	GTDengineManager.CtxMap = make(map[string]*context.Context)
	for k, v := range config.C.Tdengine {
		engine, ctx := CreateDBEngnine(v.User, v.Pwd, v.Addr, v.Port)
		GTDengineManager.EngineMap[k] = engine
		GTDengineManager.CtxMap[k] = ctx
	}
	return func() error {
		for k, engine := range GTDengineManager.EngineMap {
			err := engine.CloseDB()
			if err != nil {
				base.Glog.Errorf("> TDName= %s: %s\n", k, statuscode.ERROR_TDENGINE_CLOSE.Msg)
			} else {
				fmt.Printf("> TDName= %s close ok!\n", k)
			}
		}
		return nil
	}

}

/**
 * @description: 心跳监测
 * @return {*}
 * @author: liqiyuWorks
 */
func DBKeepAlive() {
	for t := range time.Tick(time.Minute) {
		base.Glog.Errorf("mongo tick at %v, ping error", t)
	}
}
