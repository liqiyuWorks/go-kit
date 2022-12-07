/*
 * @Author: lisheng
 * @Date: 2022-11-23 14:24:18
 * @LastEditTime: 2022-12-07 15:55:33
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/common/wave_calc/wave_calc.go
 */
package wavecalc

import "math"

type wave struct {
	Degree float64
	Height float64
	Period float64
}

func NewWave(degree, height, period float64) *wave {
	return &wave{
		Degree: degree,
		Height: height,
		Period: period,
	}
}

/**
 * @description: 计算 wave 道格拉斯量表
 * @param {float64} height wave 高度
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wave) WaveDouglasScale() string {
	height := w.Height
	scale := "Calm"
	if height <= 0.1 {
		scale = "Calm"
	} else if height <= 0.5 {
		scale = "Smooth"
	} else if height <= 1.25 {
		scale = "Slight"
	} else if height <= 2.5 {
		scale = "Moderate"
	} else if height <= 4.0 {
		scale = "Rough"
	} else if height <= 6.0 {
		scale = "VeryRough"
	} else if height <= 9.0 {
		scale = "High"
	} else if height <= 14.0 {
		scale = "VeryHigh"
	} else {
		scale = "Precipitous"
	}
	return scale
}

/**
 * @description: convert2Direction
 * @param {float64} delta
 * @return {*}
 * @author: liqiyuWorks
 */
func convert2Direction(delta float64) string {
	dire := "N/A"
	if delta < 0 {
		delta += 2 * math.Pi
	}
	const PI16 = math.Pi / 16
	if delta < PI16 {
		// 正b
		dire = "N"
	} else if delta >= PI16 && delta < 3*PI16 {
		dire = "NNE"
	} else if delta >= 3*PI16 && delta < 5*PI16 {
		dire = "NE"
	} else if delta >= 5*PI16 && delta < 7*PI16 {
		dire = "ENE"
	} else if delta >= 7*PI16 && delta < 9*PI16 {
		dire = "E"
	} else if delta >= 9*PI16 && delta < 11*PI16 {
		dire = "ESE"
	} else if delta >= 11*PI16 && delta < 13*PI16 {
		dire = "SE"
	} else if delta >= 13*PI16 && delta < 15*PI16 {
		dire = "SSE"
	} else if delta >= 15*PI16 && delta < 17*PI16 {
		dire = "S"
	} else if delta >= 17*PI16 && delta < 19*PI16 {
		dire = "SSW"
	} else if delta >= 19*PI16 && delta < 21*PI16 {
		dire = "SW"
	} else if delta >= 21*PI16 && delta < 23*PI16 {
		dire = "WSW"
	} else if delta >= 23*PI16 && delta < 25*PI16 {
		dire = "W"
	} else if delta >= 25*PI16 && delta < 27*PI16 {
		dire = "WNW"
	} else if delta >= 27*PI16 && delta < 29*PI16 {
		dire = "NW"
	} else if delta >= 29*PI16 && delta < 31*PI16 {
		dire = "NNW"
	} else if delta >= 31*PI16 && delta < 32*PI16 {
		dire = "N"
	}
	return dire
}

/**
 * @description: calc WaveDirection
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wave) WaveDirection() string {
	direction := convert2Direction((w.Degree / 360) * 2 * math.Pi)
	return direction
}
