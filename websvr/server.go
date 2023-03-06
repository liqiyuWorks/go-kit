/*
 * @Author: lisheng
 * @Date: 2022-10-13 11:38:11
 * @LastEditTime: 2022-12-15 11:29:34
 * @LastEditors: lisheng
 * @Description: web服务
 * @FilePath: /go-kit/websvr/server.go
 */

package websvr

import (
	"context"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/liqiyuworks/go-kit/database/clickhouse"
	"gitee.com/liqiyuworks/go-kit/database/mongo"
	"gitee.com/liqiyuworks/go-kit/database/mysql"
	"gitee.com/liqiyuworks/go-kit/database/pg"
	"gitee.com/liqiyuworks/go-kit/database/redis"
	"gitee.com/liqiyuworks/go-kit/database/tdengine"

	"gitee.com/liqiyuworks/go-kit/config"

	"gitee.com/liqiyuworks/go-kit/base"

	"github.com/gin-gonic/gin"
)

type CloseDBFunc func() error
type WebServer struct {
	g          *gin.Engine
	server     *http.Server
	signalStop chan os.Signal
	closeDBMap map[string]CloseDBFunc
}

/**
 * @description: 初始化数据库
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) InitDB() {
	svr.closeDBMap = make(map[string]CloseDBFunc)
	svr.closeDBMap["mysql"] = mysql.InitMysqlManager()
	svr.closeDBMap["mongo"] = mongo.InitMgocli()
	svr.closeDBMap["clickhouse"] = clickhouse.InitCkManager()
	svr.closeDBMap["redis"] = redis.InitRedisClient()
	svr.closeDBMap["pg"] = pg.InitPgManager()
	svr.closeDBMap["tdengine"] = tdengine.InitTdClient()
}

/**
 * @description: 初始化log
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) InitLog() {
	config.Initialize("config/app.json") // 配置文件
	base.CreateLogFile()                 // 检测日志文件是否存在
	base.InitLogger()                    // 日志初始化
	base.Glog.Info(" logger init ok ...")
}

/**
 * @description: 初始化路由相关服务
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) InitWebServer() {
	g := InitRouter() // 加载自动注册路由模块
	svr.g = g
	svr.signalStop = make(chan os.Signal, 1)
	svr.server = &http.Server{
		Addr:    config.C.Server.Listen,
		Handler: g,
	}
}

/**
 * @description: 启动server
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) Run() {
	go func() {
		if err := svr.server.ListenAndServe(); err != nil {
			base.Glog.Info("http server run error: %v", err)
		}
	}()
	signal.Notify(svr.signalStop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-svr.signalStop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	svr.clean(ctx)
}

/**
 * @description: 清理server
 * @param {context.Context} ctx
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) clean(ctx context.Context) {
	if err := svr.server.Shutdown(ctx); err != nil {
		base.Glog.Info("http server shutdown error: %v", err)
	} else {
		base.Glog.Info("http server clean finish")
	}
}

/**
 * @description: 统一清理数据库连接
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) CloseDB() {
	for serverName := range svr.closeDBMap {
		if err := svr.closeDBMap[serverName](); err != nil {
			base.Glog.Errorf("DB Engine: %s clean error: %v", serverName, err)
		}
	}
}
