/*
 * @Author: lisheng
 * @Date: 2022-10-14 16:28:38
 * @LastEditTime: 2022-11-02 23:35:16
 * @LastEditors: lisheng
 * @Description: mongoDB驱动
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/database/mongo/mongo.go
 */
package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Client instance

type MgoManager struct {
	EngineMap map[string]*mongo.Client
}

var (
	GMgoManager *MgoManager = new(MgoManager)
)

type Pagination struct {
	page     int64
	pageSize int64
}

func CreateDBEngnine(addr string, port int) *mongo.Client {
	var err error
	var mgoCli *mongo.Client
	url := fmt.Sprintf("mongodb://%s:%d/?connect=direct", addr, port)
	mgoCli, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	if err = mgoCli.Ping(context.TODO(), readpref.Primary()); err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_CONNECT.Msg, err)
	} else {
		base.Glog.Infof(">>> Connected to MongoDB")
	}

	go DBKeepAlive()
	return mgoCli

}

func InitMgocli() func() error {
	GMgoManager.EngineMap = make(map[string]*mongo.Client)
	for k, v := range config.C.Mongo {
		engine := CreateDBEngnine(v.Addr, v.Port)
		GMgoManager.EngineMap[k] = engine
	}

	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		for k, engine := range GMgoManager.EngineMap {
			err := engine.Disconnect(ctx)
			if err != nil {
				base.Glog.Errorf("> mongoName= %s: %s\n", k, statuscode.ERROR_MONGO_CLOSE.Msg)
			} else {
				fmt.Printf("> mongoName= %s close ok!\n", k)
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
		for k, engine := range GMgoManager.EngineMap {
			err := engine.Ping(context.TODO(), readpref.Primary())
			if err != nil {
				base.Glog.Errorf("Ping mongoSrc[%s] failed! At:%v, reason:[%s]", k, t, err.Error())
			}
		}
	}
}
