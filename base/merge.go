/*
 * @Author: lisheng
 * @Date: 2022-10-29 11:42:37
 * @LastEditTime: 2022-10-29 11:47:17
 * @LastEditors: lisheng
 * @Description: 合并各种类型模块
 * @FilePath: /jf-go-kit/base/merge.go
 */
package base

import "bytes"

/**
 * @description: 合并maps
 * @param {...map[string]interface{}} mObj
 * @return {*}
 * @author: liqiyuWorks
 */
func MergeMaps(mObj ...map[string]interface{}) map[string]interface{} {
	newObj := map[string]interface{}{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

/**
 * @description: 合并字符串
 * @param {...string} strs
 * @return {*}
 * @author: liqiyuWorks
 */
func MergeStrings(strs ...string) string {
	var bt bytes.Buffer
	for i := range strs {
		bt.WriteString(strs[i])
	}
	return bt.String()
}
