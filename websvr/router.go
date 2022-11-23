/*
 * @Author: lisheng
 * @Date: 2022-10-10 23:40:57
 * @LastEditTime: 2022-11-23 15:24:04
 * @LastEditors: lisheng
 * @Description: 路由模块
 * @FilePath: /jf-go-kit/websvr/router.go
 */
package websvr

import (
	"net/http"
	"reflect"
	"strings"

	"gitee.com/liqiyuworks/jf-go-kit/middleware"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 路由结构体
type Route struct {
	path   string         //url路径
	method reflect.Value  //方法路由
	args   []reflect.Type //参数类型
}

// 路由集合
var Routes = []Route{}

var JWT_ROUTE_ARRAY = []string{"/user/login", "/user/register", "/v1/ocean"} // 不进行JWT认证的路由

func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.Default()
	pprof.Register(r) // pprof性能监测
	// 注册中间件
	r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping ok")
	})

	Bind(r)

	return r
}

// 注册控制器
func Register(handler interface{}) bool {
	ctrlName := reflect.TypeOf(handler).String()
	// base.Glog.Infoln("> ctrlName=", ctrlName)
	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".")+1:]
	}
	base.Glog.Info("module=", module)
	v := reflect.ValueOf(handler)
	module = base.ConvertCamelToCase(module)
	// 遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		case_action := base.ConvertCamelToCase(action)

		// 遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			// base.Glog.Infoln("params-name=", method.Type().In(j))
		}
		// base.Glog.Infoln("params=", params)
		// fmt.Printf(">> action = %s -> case_action = %s \n", action, case_action)
		path := "/" + module + "/" + case_action
		route := Route{path: path, method: method, args: params}
		Routes = append(Routes, route)
	}
	// base.Glog.Infoln("Routes=", Routes)
	return true
}

// 绑定路由 m是方法GET POST等
// 绑定基本路由 e.POST(path, match(path))
func Bind(e *gin.Engine) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册 swagger
	for _, route := range Routes {
		if base.In(route.path, JWT_ROUTE_ARRAY) {
			// base.Glog.Infoln(">>> Not use jwt route path: ", route.path)
			e.GET(route.path, match(route))
			e.POST(route.path, match(route))
		} else {
			// e.GET(route.path, middleware.JWTAuth(), match(route.path, route))
			// e.POST(route.path, middleware.JWTAuth(), match(route.path, route))
			e.GET(route.path, match(route))
			e.POST(route.path, match(route))
		}
	}
}

/**
 * @description: 根据path匹配对应的方法
 * @param {string} path
 * @param {Route} route
 * @return {*}
 * @author: liqiyuWorks
 */
func match(route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		route.method.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
