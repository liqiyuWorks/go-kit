/*
 * @Author: lisheng
 * @Date: 2022-10-29 11:20:48
 * @LastEditTime: 2022-10-29 11:58:45
 * @LastEditors: lisheng
 * @Description: 保留小数的精度
 * @FilePath: /gitee.com/wuxi_jiufang/jf-go-kit/base/round.go
 */
package base

import (
	"fmt"
	"strconv"
)

/**
 * @description: 保留float64两位小数
 * @param {float64} num
 * @return {*}
 * @author: liqiyuWorks
 */
func Decimal2(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func Decimal1(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", num), 64)
	return num
}
