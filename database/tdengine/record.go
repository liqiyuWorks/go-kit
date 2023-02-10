/*
 * @Author: lisheng
 * @Date: 2022-11-23 14:06:29
 * @LastEditTime: 2023-02-10 22:29:56
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/database/tdengine/record.go
 */
package tdengine

import (
	"gitee.com/liqiyuworks/jf-go-kit/common/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"gitee.com/chunanyong/zorm"
)

/**
 * @description: C - 插入数据
 * @param {string} exeSql
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func RestInsertRecordPtr(engineName, exeSql string, objPtr interface{}) {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(*ctx, finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_INSERT_RECORD.Msg, err)
	}
}

/**
 * @description: RestFindRecordsPtr
 * @param {*} engineName
 * @param {string} exeSql
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func RestFindRecordsPtr(engineName, exeSql string, objPtr interface{}) error {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(*ctx, finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_FIND_RECORDS.Msg, err)
	}
	return err
}

/**
 * @description: U - 更新数据
 * @param {string} exeSql
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func RestUpdateRecordPtr(engineName, exeSql string, objPtr interface{}) error {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(*ctx, finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_UPDATE_RECORD.Msg, err)
	}
	return err
}

/**
 * @description: D - 更新数据
 * @param {string} exeSql
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func RestDeleteRecordPtr(engineName, exeSql string, objPtr interface{}) error {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(*ctx, finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_DELETE_RECORD.Msg, err)
	}
	return err
}

/**
 * @description: RestFindRecordsMap
 * @param {*} engineName
 * @param {string} exeSql
 * @return {*} 返回map
 * @author: liqiyuWorks
 */
func RestFindRecordsMap(engineName, exeSql string) ([]map[string]interface{}, error) {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	reMaps, err := zorm.QueryMap(*ctx, finder, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_FIND_RECORDS.Msg, err)
	}
	return reMaps, err
}

func RestFindSingeRecordMap(engineName, exeSql string) (map[string]interface{}, error) {
	if engineName == "" {
		engineName = "default"
	}
	ctx := GTDengineManager.CtxMap[engineName]
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	reMaps, err := zorm.QueryRowMap(*ctx, finder)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_FIND_RECORDS.Msg, err)
	}
	return reMaps, err
}
