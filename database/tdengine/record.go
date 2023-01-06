package tdengine

import (
	"context"

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
func RestInsertRecordPtr(exeSql string, objPtr interface{}) {
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(context.Background(), finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_INSERT_RECORD.Msg, err)
	}
}

/**
 * @description: R - TD查询返回数组结构体
 * @param {string} exeSql
 * @param {interface{}} objPtr
 * @return {*}
 * @author: liqiyuWorks
 */
func RestFindRecordsPtr(exeSql string, objPtr interface{}) error {
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(context.Background(), finder, objPtr, nil)
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
func RestUpdateRecordPtr(exeSql string, objPtr interface{}) error {
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(context.Background(), finder, objPtr, nil)
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
func RestDeleteRecordPtr(exeSql string, objPtr interface{}) error {
	finder := zorm.NewFinder()
	finder.Append(exeSql)
	err := zorm.Query(context.Background(), finder, objPtr, nil)
	if err != nil {
		base.Glog.Errorf("%s: %v", statuscode.ERROR_TDENGINE_DELETE_RECORD.Msg, err)
	}
	return err
}
