/*
 * @Author: lisheng
 * @Date: 2022-10-29 11:14:13
 * @LastEditTime: 2023-01-16 11:21:29
 * @LastEditors: lisheng
 * @Description: covert A类型 To B类型
 * @FilePath: /jf-go-kit/base/convert.go
 */
package base

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

var i64Type reflect.Type
var f64Type reflect.Type
var strType reflect.Type

func init() {
	i64Type = reflect.TypeOf(int64(0))
	f64Type = reflect.TypeOf(float64(0))
	strType = reflect.TypeOf("")
}

/**
 * @description: 转int
 * @param {interface{}} i
 * @param {int} defaultValue: 0
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertToInt(i interface{}, defaultValue int) int {
	return int(ConvertToInt64(i, int64(defaultValue)))
}

/**
 * @description: 转int64
 * @param {interface{}} i
 * @param {int64} defaultValue:0
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertToInt64(i interface{}, defaultValue int64) (ret int64) {
	defer func() {
		if r := recover(); r != nil {
			// 转换出错，设置默认值
			// Glog.Errorf("Convert %v to %T failed[%v]", i, ret, r)
			ret = defaultValue
		}
	}()

	v := reflect.ValueOf(i)

	// 字符串类型
	if v.Kind() == reflect.String {
		strValue := v.String()
		ret, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			// Glog.Errorf("Convert \"%s\" to %T failed[%v]", strValue, ret, err)
			ret = defaultValue
		}

		return ret
	}

	i64Value := v.Convert(i64Type)
	ret = i64Value.Int()
	return ret
}

/**
 * @description: 转float64
 * @param {interface{}} i
 * @param {float64} defaultValue:0
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertToFloat64(i interface{}, defaultValue float64) (ret float64) {
	defer func() {
		if r := recover(); r != nil {
			// 转换出错，设置默认值
			// Glog.Errorf("Convert %v to %T failed[%v]", i, ret, r)
			ret = defaultValue
		}
	}()

	v := reflect.ValueOf(i)

	// 字符串类型
	if v.Kind() == reflect.String {
		strValue := v.String()
		ret, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			// Glog.Errorf("Convert \"%s\" to %T failed[%v]", strValue, ret, err)
			ret = defaultValue
		}

		return ret
	}

	f64Value := v.Convert(f64Type)
	ret = f64Value.Float()
	return ret
}

/**
 * @description: 转string
 * @param {interface{}} i
 * @param {string} defaultValue: ""
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertToString(i interface{}, defaultValue string) (ret string) {
	defer func() {
		if r := recover(); r != nil {
			// 转换出错，设置默认值
			Glog.Errorf("Convert %v to %T failed[%v]", i, ret, r)
			ret = defaultValue
		}
	}()

	v := reflect.ValueOf(i)
	strValue := v.Convert(strType)
	ret = strValue.String()
	return ret
}

/**
 * @description: Struct转Map
 * @param {interface{}} i
 * @param {interface{}} defaultValue: nil
 * @param {interface{}} ret
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertStructToMap(i interface{}, defaultValue interface{}, ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			// 转换出错，设置默认值
			// Glog.Errorf("Convert %v to %T failed[%v]", i, ret, r)
			ret = defaultValue
		}
	}()
	if marshalContent, err := json.Marshal(i); err != nil {
		Glog.Infoln(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber()
		d.Decode(&ret)
	}
}

/**
 * @description: 驼峰式写法转为下划线写法
 * @param {string} name
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertCamelToCase(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

/**
 * @description: ConvertCamelToSlash
 * @param {string} name
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertCamelToSlash(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('/')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

/**
 * @description: 下划线写法转为驼峰写法
 * @param {string} name
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertCaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

/**
 * @description: ConvertStringToByte
 * @param {string} s
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertStringToByte(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))

}

/**
 * @description: ConvertByteToString
 * @param {[]byte} bytes
 * @return {*}
 * @author: liqiyuWorks
 */
func ConvertByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))

}
