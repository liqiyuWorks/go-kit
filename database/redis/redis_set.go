/*
 * @Author: lisheng
 * @Date: 2022-10-25 15:11:22
 * @LastEditTime: 2023-01-06 14:17:17
 * @LastEditors: lisheng
 * @Description:  redis字符串数据操作
 * @FilePath: /jf-go-kit/database/redis/redis_set.go
 */
package redis

import (
	"context"

	"gitee.com/liqiyuworks/jf-go-kit/common/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

/**
 * @description: SetMembers
 * @param {string} engineName
 * @param {string} key
 * @return {*}
 * @author: liqiyuWorks
 */
func SetMembers(engineName string, key string) (map[string][]string, error) {
	if engineName == "" {
		engineName = "default"
	}
	data := make(map[string][]string)
	value, err := GRdsManager.EngineMap[engineName].SMembers(context.Background(), key).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_SET_MEMBERS.Msg, err.Error())
		return nil, err
	}
	data[key] = value
	return data, nil
}

/**
 * @description: SetAdd
 * @param {string} engineName
 * @param {string} key
 * @param {...interface{}} value
 * @return {*}
 * @author: liqiyuWorks
 */
func SetAdd(engineName string, key string, value ...interface{}) (affected int64, err error) {
	if engineName == "" {
		engineName = "default"
	}
	affected, err = GRdsManager.EngineMap[engineName].SAdd(context.Background(), key, value...).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_SET_ADD.Msg, err.Error())
		return affected, err
	}
	return affected, nil
}

/**
 * @description: SetZange
 * @param {string} engineName
 * @param {string} key
 * @param {*} start
 * @param {int64} stop
 * @return {*}
 * @author: liqiyuWorks
 */
func SetZange(engineName string, key string, start, stop int64) (map[string][]string, error) {
	if engineName == "" {
		engineName = "default"
	}
	data := make(map[string][]string)
	value, err := GRdsManager.EngineMap[engineName].ZRange(context.Background(), key, start, stop).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_SET_ZANGE.Msg, err.Error())
		return nil, err
	}
	data[key] = value
	return data, nil
}
