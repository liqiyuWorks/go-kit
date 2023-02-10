package pg

import (
	"gitee.com/wuxi_jiufang/jf-go-kit/common/statuscode"

	"gitee.com/wuxi_jiufang/jf-go-kit/base"
)

/**
 * @description: C - 插入数据
 * @param {string} tableName
 * @param {string} cond
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func InsertRecord(engineName string, tableName string, objPtr interface{}) (affectedRows int64, err error) {
	if engineName == "" {
		engineName = "default"
	}
	affectedRows = 0
	err = nil
	result := GPgManager.EngineMap[engineName].Table(tableName).Create(objPtr)
	if result.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_INSERT_RECORD.Msg, result.Error)
		return result.RowsAffected, result.Error
	}
	if result.RowsAffected == 0 {
		return
	}
	return result.RowsAffected, nil
}

/**
 * @description: R - Where 查询 根据条件查询
 * @param {string} tableName
 * @param {string} cond
 * @param {string} condValue
 * @param {int} page
 * @param {int} pageSize
 * @param {interface{}} resultSlicePtr
 * @return {*}
 * @author: liqiyuWorks
 */
func FindRecordsByCond(engineName string, tableName string, cond string, page int, pageSize int, resultSlicePtr interface{}) error {
	if engineName == "" {
		engineName = "default"
	}
	if result := GPgManager.EngineMap[engineName].Table(tableName).Where(cond).Find(resultSlicePtr); result.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_FIND_RECORDS.Msg, result.Error)
		return result.Error
	}
	return nil
}

/**
 * @description: U - Where 更新 根据条件查询
 * @param {string} tableName
 * @param {string} cond
 * @param {map[string]interface{}} fieldMap
 * @return {*}
 * @author: liqiyuWorks
 */
func UpdateRecordByCond(engineName string, tableName string, cond string, fieldMap map[string]interface{}) (affectedRows int64, err error) {
	if engineName == "" {
		engineName = "default"
	}
	affectedRows = 0
	err = nil
	if result := GPgManager.EngineMap[engineName].Table(tableName).Where(cond).Updates(fieldMap); result.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_UPDATE_RECORD.Msg, result.Error)
		return result.RowsAffected, result.Error
	}
	return
}

/**
 * @description: D - Where 删除 根据条件查询
 * @param {string} tableName
 * @param {string} cond
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func DeleteRecordByCond(engineName string, tableName string, cond string, objPtr interface{}) (affectedRows int64, err error) {
	if engineName == "" {
		engineName = "default"
	}
	affectedRows = 0
	err = nil
	if result := GPgManager.EngineMap[engineName].Table(tableName).Where(cond).Delete(objPtr); result.Error != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_PG_DELETE_RECORD.Msg, result.Error)
		return result.RowsAffected, result.Error
	}
	return
}
