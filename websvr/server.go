/*
 * @Author: lisheng
 * @Date: 2022-10-13 11:38:11
 * @LastEditTime: 2022-12-05 14:01:14
 * @LastEditors: lisheng
 * @Description: web服务
 * @FilePath: /jf-go-kit/websvr/server.go
 */

package websvr

import (
	"context"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/database/clickhouse"
	"gitee.com/liqiyuworks/jf-go-kit/database/mongo"
	"gitee.com/liqiyuworks/jf-go-kit/database/mysql"
	"gitee.com/liqiyuworks/jf-go-kit/database/pg"
	"gitee.com/liqiyuworks/jf-go-kit/database/redis"
	"gitee.com/liqiyuworks/jf-go-kit/database/tdengine"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"github.com/gin-gonic/gin"
)

var (
	ResiPoints = &NewResiPoints{} // 常驻内存
)

type CloseDBFunc func() error
type WebServer struct {
	g          *gin.Engine
	server     *http.Server
	signalStop chan os.Signal
	closeDBMap map[string]CloseDBFunc
}

type NewResiPoints struct {
	MfwamLonsMap map[string][]string
	MfwamLatsMap map[string][]string
	SmocLonsMap  map[string][]string
	SmocLatsMap  map[string][]string
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
 * @description: 初始化常驻内存
 * @return {*}
 * @author: liqiyuWorks
 */
func (svr *WebServer) InitResidentMemory() {
	var err error
	ResiPoints.MfwamLonsMap, err = redis.SetZange("points", "base_data:mfwam:lons", 0, -1)
	if err != nil {
		base.Glog.Errorf("init state mfwam lons points failed, err=%v", err)
	}
	ResiPoints.MfwamLatsMap, err = redis.SetZange("points", "base_data:mfwam:lats", 0, -1)
	if err != nil {
		base.Glog.Errorf("init state mfwam lats points failed, err=%v", err)
	}
	ResiPoints.SmocLonsMap, err = redis.SetZange("points", "base_data:smoc:lons", 0, -1)
	if err != nil {
		base.Glog.Errorf("init state smoc lons points failed, err=%v", err)
	}
	ResiPoints.SmocLatsMap, err = redis.SetZange("points", "base_data:smoc:lats", 0, -1)
	if err != nil {
		base.Glog.Errorf("init state smoc lats points failed, err=%v", err)
	}
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
