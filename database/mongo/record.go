package mongo

import (
	"context"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * @description: C-插入数据
 * @param {string} tableName
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func InsertRecord(engineName string, tableName string, objPtr interface{}) (int64, error) {
	if engineName == "" {
		engineName = "default"
	}
	opts := options.InsertOne()
	result, err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo["default"].DB).Collection(tableName).InsertOne(context.TODO(), objPtr, opts)
	// base.Glog.Infoln(">>> result", result)
	if err != nil {
		base.Glog.Errorf("%s: %v: , result: %v", statuscode.ERROR_MONGO_INSERT_RECORD.Msg, err, result)
		return -1, err
	}
	return 1, nil
}

/**
 * @description: R - 根据查询条件查看是否存在该条记录
 * @param {string} tableName
 * @param {*map[string]string} queryMap
 * @param {interface{}} user
 * @return {*}
 * @author: liqiyuWorks
 */
func ExistByQueryMap(engineName string, tableName string, queryMap *map[string]string, user interface{}) (bool, error) {
	if engineName == "" {
		engineName = "default"
	}
	err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo["default"].DB).Collection(tableName).FindOne(context.Background(), queryMap).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_EXIST_RECORD.Msg, err)
		return false, err
	}
	if err == nil && user != nil {
		return true, nil
	}
	return false, nil
}

/**
 * @description: R - 获取单条记录
 * @param {string} tableName
 * @param {*map[string]string} queryMap
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func QueryRecord(engineName string, tableName string, queryMap *map[string]interface{}, objPtr interface{}) error {
	if engineName == "" {
		engineName = "default"
	}
	err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo[engineName].DB).Collection(tableName).FindOne(context.Background(), queryMap).Decode(objPtr)
	if err != nil && err != mongo.ErrNoDocuments {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_QUERY_RECORD.Msg, err)
		return err
	}
	return nil
}

/**
 * @description: R - 查询多条数据
 * @param {string} tableName
 * @param {*map[string]string} queryMap: 查询map{"username": "jiufang"}
 * @param {interface{}} objPtr: 绑定查询对象
 * @param {...Pagination} pagerArg: 分页器，可传入指定的分页
 * @return {*}
 * @author: liqiyuWorks
 */
func FindRecords(engineName string, tableName string, queryMap *map[string]string, objPtr interface{}, pagerArg ...Pagination) error {
	if engineName == "" {
		engineName = "default"
	}
	var pager Pagination
	if len(pagerArg) != 0 {
		pager = pagerArg[0]
	} else {
		pager.page = 1
		pager.pageSize = 10
	}

	ctx := context.Background()
	opts := options.Find().SetLimit(pager.pageSize).SetSkip((int64(pager.page) - pager.page) * pager.pageSize) //分页
	cursor, err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo[engineName].DB).Collection(tableName).Find(ctx, queryMap, opts)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_FIND_RECORDS.Msg, err)
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, objPtr)
	return err
}

/**
 * @description: U - 更新数据
 * @param {string} tableName
 * @param {*map[string]string} queryMap
 * @param {*map[string]string} updateMap
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func UpateRecord(engineName string, tableName string, queryMap *map[string]string, updateMap *map[string]map[string]string) (int64, error) {
	if engineName == "" {
		engineName = "default"
	}
	result, err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo[engineName].DB).Collection(tableName).UpdateMany(context.TODO(), queryMap, updateMap)
	base.Glog.Infoln(">>> result ", result)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_UPDATE_RECORD.Msg, err)
		return -1, err
	}
	return result.ModifiedCount, nil
}

/**
 * @description: D - 删除数据
 * @param {string} tableName
 * @param {*map[string]string} deleteMap
 * @return {*}
 * @author: liqiyuWorks
 */
func DeleteRecord(engineName string, tableName string, deleteMap *map[string]string) (int64, error) {
	if engineName == "" {
		engineName = "default"
	}
	affectedRows, err := GMgoManager.EngineMap[engineName].Database(config.C.Mongo[engineName].DB).Collection(tableName).DeleteMany(context.TODO(), deleteMap)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_MONGO_DELETE_RECORD.Msg, err)
		return -1, err
	}
	return affectedRows.DeletedCount, nil
}
