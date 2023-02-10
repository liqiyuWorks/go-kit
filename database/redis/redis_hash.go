package redis

import (
	"context"

	"gitee.com/liqiyuworks/jf-go-kit/common/statuscode"
	"github.com/go-redis/redis/v8"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

/**
 * @description:
 * @param {*} key
 * @param {string} field
 * @return {*}
 * @author: liqiyuWorks
 */
func HashGet(engineName string, key string) (map[string]string, error) {
	if engineName == "" {
		engineName = "default"
	}
	data, err := GRdsManager.EngineMap[engineName].HGetAll(context.Background(), key).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_HASH_GET.Msg, err.Error())
	}
	return data, err
}

/**
 * @description: HashSet
 * @param {*} key
 * @param {string} field
 * @param {interface{}} value
 * @return {*}
 * @author: liqiyuWorks
 */
func HashSet(engineName string, key, field string, value interface{}) (bool, error) {
	if engineName == "" {
		engineName = "default"
	}
	value, err := GRdsManager.EngineMap[engineName].HSet(context.Background(), key, field, value).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_HASH_SET.Msg, err.Error())
		return false, err
	}
	return true, nil
}

/**
 * @description: HashGetFields
 * @param {*} key
 * @param {string} field
 * @return {*}
 * @author: liqiyuWorks
 */
func HashGetFields(engineName string, key, field string) (string, error) {
	if engineName == "" {
		engineName = "default"
	}
	data, err := GRdsManager.EngineMap[engineName].HGet(context.Background(), key, field).Result()
	if (err.Error()) == redis.Nil.Error() {
		return data, nil
	}
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_HASH_GET_FIELDS.Msg, err.Error())
	}
	return data, err
}

/**
 * @description: Hash Delete
 * @param {string} key
 * @param {...string} fields
 * @return {*}
 * @author: liqiyuWorks
 */
func HashDelete(engineName string, key string, fields ...string) (int64, error) {
	if engineName == "" {
		engineName = "default"
	}
	count, err := GRdsManager.EngineMap[engineName].HDel(context.Background(), key, fields...).Result()
	if err != nil {
		base.Glog.Errorf("> %s : %v", statuscode.ERROR_REDIS_HASH_DELETE.Msg, err.Error())
		return count, err
	}
	return count, nil
}
