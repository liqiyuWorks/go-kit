/*
 * @Author: lisheng
 * @Date: 2022-10-12 17:57:22
 * @LastEditTime: 2023-01-06 14:16:24
 * @LastEditors: lisheng
 * @Description: Redis驱动
 * @FilePath: /go-kit/database/redis/redis.go
 */
package redis

import (
	"context"
	"fmt"
	"time"

	"gitee.com/liqiyuworks/go-kit/config"

	"gitee.com/liqiyuworks/go-kit/common/statuscode"

	"gitee.com/liqiyuworks/go-kit/base"

	"github.com/go-redis/redis/v8"
)

type RdsManager struct {
	EngineMap map[string]*redis.Client
}

var (
	GRdsManager *RdsManager = new(RdsManager)
)

func CreateDBEngnine(addr, pwd string) *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rdsCli := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
		MaxRetries:   1,
		MinIdleConns: 1,
	})

	_, err := rdsCli.Ping(ctx).Result()
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_REDIS_CONNECT.Msg, err)
	} else {
		base.Glog.Infof(">>> Connected to Redis")

	}

	go DBKeepAlive() // 心跳监测
	return rdsCli

}

/**
 * @description: 初始化redisClient
 * @return {*}
 * @author: liqiyuWorks
 */
func InitRedisClient() func() error {
	GRdsManager.EngineMap = make(map[string]*redis.Client)
	for k, v := range config.C.Redis {
		engine := CreateDBEngnine(v.Addr, v.Pwd)
		GRdsManager.EngineMap[k] = engine
	}

	return func() error {
		for k, engine := range GRdsManager.EngineMap {
			err := engine.Close()
			if err != nil {
				base.Glog.Errorf("> redisName= %s: %s\n", k, statuscode.ERROR_REDIS_CLOSE.Msg)
			} else {
				fmt.Printf("> redisName= %s close ok!\n", k)
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
		for k, engine := range GRdsManager.EngineMap {
			_, err := engine.Ping(context.Background()).Result()
			if err != nil {
				base.Glog.Errorf("Ping redisSrc[%s] failed! At:%v, reason:[%s]", k, t, err.Error())
			}
		}
	}
}
