/*
 * @Author: lisheng
 * @Date: 2022-10-24 14:57:04
 * @LastEditTime: 2022-10-28 15:38:12
 * @LastEditors: lisheng
 * @Description: 反距离权重计算模块
 * @FilePath: /gitee.com/liqiyuworks/go-kit/common/idw.go
 */
package idw

import (
	"math"
)

type IdwPoint struct {
	X      float64
	Y      float64
	Weight float64
}

/**
 * @description: 反距离权重计算
 * @param {int} x 124.41
 * @param {int} y 32.2
 * @param {[]point} points [[124.416664, 32.25, 0.34], [124.333336, 32.166668, 0.34], [124.416664, 32.166668, 0.34], [124.333336, 32.25, 0.34]]
 * @return {*} 0.34
 * @author: liqiyuWorks
 */
func IwdIstance(x float64, y float64, points []IdwPoint) float64 {
	L := len(points)
	nominator := 0.0
	denominator := 0.0
	for i := 0; i < L; i++ {
		pt := points[i]
		dist := math.Sqrt(float64((x-pt.X)*(x-pt.X) + (y-pt.Y)*(y-pt.Y)))
		if dist < 0.0000000001 {
			return pt.Weight
		}
		nominator = nominator + (pt.Weight / math.Pow(dist, 2)) // default power = 2
		denominator = denominator + (1 / math.Pow(dist, 2))
	}

	value := 0.0
	if denominator > 0 {
		value = nominator / denominator
	} else {
		value = -9999
	}
	return value
}
