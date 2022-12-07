/*
 * @Author: lisheng
 * @Date: 2022-11-23 14:24:18
 * @LastEditTime: 2022-12-07 15:32:52
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/common/wind_calc/wind_calc.go
 */
package windcalc

import (
	"math"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

type wind struct {
	U10 float64
	V10 float64
}

func NewWind(u10, v10 float64) *wind {
	return &wind{
		U10: u10,
		V10: v10,
	}
}

/**
 * @description: 计算风速
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) WindSpeed() float64 {
	value := math.Pow(math.Pow(w.U10, 2)+math.Pow(w.V10, 2), 0.5)
	return base.Decimal1(value)
}

/**
 * @description: 获取风向角度
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) WindAngle() float64 {
	angle := 999.9
	u := w.U10
	v := w.V10
	if u > 0 && v > 0 {
		angle = 270 - math.Atan(v/u)*180/math.Pi
	} else if u < 0 && 0 < v {
		angle = 90 - math.Atan(v/u)*180/math.Pi
	} else if u < 0 && v < 0 {
		angle = 90 - math.Atan(v/u)*180/math.Pi
	} else if u > 0 && 0 > v {
		angle = 270 - math.Atan(v/u)*180/math.Pi
	} else if u == 0 && v > 0 {
		angle = 180
	} else if u == 0 && v < 0 {
		angle = 0
	} else if u > 0 && v == 0 {
		angle = 270
	} else if u < 0 && v == 0 {
		angle = 90
	} else if u == 0 && v == 0 {
		angle = 999.9
	}
	return base.Decimal2(angle)
}

/**
 * @description: 根据风向角获取风向
 * @param {float64} angle
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) WindDirection(angle float64) string {
	var direction string
	if angle < 11.25 {
		direction = "N"
	} else if angle <= 33.75 {
		direction = "NNE"

	} else if angle <= 56.25 {
		direction = "NE"

	} else if angle <= 78.75 {
		direction = "ENE"

	} else if angle <= 101.25 {
		direction = "E"

	} else if angle <= 123.75 {
		direction = "ESE"

	} else if angle <= 146.25 {
		direction = "SE"

	} else if angle <= 168.75 {
		direction = "SSE"

	} else if angle <= 191.25 {
		direction = "S"

	} else if angle <= 213.75 {
		direction = "SSW"

	} else if angle <= 236.25 {
		direction = "SW"

	} else if angle <= 258.75 {
		direction = "WSW"

	} else if angle <= 281.25 {
		direction = "W"

	} else if angle <= 303.75 {
		direction = "WNW"

	} else if angle <= 326.25 {
		direction = "NW"

	} else if angle <= 348.75 {
		direction = "NNW"

	} else {
		direction = "N"
	}

	return direction
}

/**
 * @description: convert speed to knots
 * @param {float64} speed
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) WindKnots(speed float64) float64 {
	knots := (speed * 3600) / 1852
	return base.Decimal1(knots)
}

/**
 * @description: 计算 蒲福风级
 * @param {float64} speed
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) WindBeaufortWindForceScale(speed float64) int {
	var scale int = 0
	if speed <= 0.2 {
		scale = 0
	} else if speed <= 1.5 {
		scale = 1
	} else if speed <= 3.3 {
		scale = 2
	} else if speed <= 5.4 {
		scale = 3
	} else if speed <= 7.9 {
		scale = 4
	} else if speed <= 10.7 {
		scale = 5
	} else if speed <= 13.8 {
		scale = 6
	} else if speed <= 17.1 {
		scale = 7
	} else if speed <= 20.7 {
		scale = 8
	} else if speed <= 22.4 {
		scale = 9
	} else if speed <= 28.4 {
		scale = 10
	} else if speed <= 32.6 {
		scale = 11
	} else if speed > 32.6 {
		scale = 12
	}
	return scale
}
