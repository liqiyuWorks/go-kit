/*
 * @Author: lisheng
 * @Date: 2022-11-23 14:24:18
 * @LastEditTime: 2022-12-07 15:25:41
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/common/wave_calc/wave_calc.go
 */
package wavecalc

type wave struct {
	U10 float64
	V10 float64
}

func NewWave(u10, v10 float64) *wave {
	return &wave{
		U10: u10,
		V10: v10,
	}
}

/**
 * @description: 计算 wave 道格拉斯量表
 * @param {float64} height wave 高度
 * @return {*}
 * @author: liqiyuWorks
 */
func (w *wave) WaveDouglasScale(height float64) string {
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
