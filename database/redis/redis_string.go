/*
 * @Author: lisheng
 * @Date: 2022-10-25 15:11:22
 * @LastEditTime: 2022-11-04 10:19:44
 * @LastEditors: lisheng
 * @Description:  redis字符串数据操作
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/database/redis/redis_string.go
 */
package redis

import (
	"context"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/base/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

/**
 * @description: String GET
 * @param {string} key
 * @return {*}: 字典
 * @author: liqiyuWorks
 */
func StringGet(engineName string, key string) (map[string]string, error) {
	if engineName == "" {
		engineName = "default"
	}
	data := make(map[string]string)
	value, err := GRdsManager.EngineMap[engineName].Get(context.Background(), key).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_STRING_GET.Msg, err.Error())
		return nil, err
	}
	data[key] = value
	return data, nil
}

/**
 * @description: String SET
 * @param {string} key
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @return {*}
 * @author: liqiyuWorks
 */
func StringSet(engineName string, key string, value interface{}, expiration time.Duration) (bool, error) {
	if engineName == "" {
		engineName = "default"
	}
	value, err := GRdsManager.EngineMap[engineName].Set(context.Background(), key, value, expiration).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_STRING_SET.Msg, err.Error())
		return false, err
	}
	return true, nil
}

/**
 * @description: String Delete
 * @param {...string} keys
 * @return {*}
 * @author: liqiyuWorks
 */
func StringDelete(engineName string, keys ...string) (bool, error) {
	if engineName == "" {
		engineName = "default"
	}
	_, err := GRdsManager.EngineMap[engineName].Del(context.Background(), keys...).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_STRING_DELETE.Msg, err.Error())
		return false, err
	}
	return true, nil
}
