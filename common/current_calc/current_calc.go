/*
 * @Author: lisheng
 * @Date: 2022-11-23 14:24:18
 * @LastEditTime: 2023-01-09 12:20:53
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/common/current_calc/current_calc.go
 */
package currentcalc

import (
	"math"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

type wind struct {
	U10 float64
	V10 float64
}

func NewCurrent(u10, v10 float64) *wind {
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
func (w *wind) CurrentSpeed() float64 {
	value := math.Pow(math.Pow(w.U10, 2)+math.Pow(w.V10, 2), 0.5)
	return base.Decimal1(value)
}

/**
 * @description: 获取风向角度
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wind) CurrentAngle() float64 {
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
func (w *wind) CurrentDirection(angle float64) string {
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
func (w *wind) CurrentKnots(speed float64) float64 {
	knots := (speed * 3600) / 1852
	return base.Decimal1(knots)
}

func (w *wind) CurrentFactor(azimuth int, speed, angle float64) float64 {
	var theta float64
	var factor float64
	if azimuth != 0 {
		theta = (base.ConvertToFloat64(float64(azimuth), 0) - angle) / 180 * math.Pi
	} else {
		theta = (angle / 180) * math.Pi
	}
	if theta == math.Pi/2 || theta == -math.Pi/2 {
		factor = 0
	} else {
		factor = speed * math.Cos(theta)
	}
	return base.Decimal1(factor)
}
