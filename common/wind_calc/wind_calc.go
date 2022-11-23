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

func (w *wind) WindDirection() int {
	angle := w.WindAngle()
	var direction int = 0
	if angle < 11.25 {
		direction = 1
	} else if angle <= 33.75 {
		direction = 2

	} else if angle <= 56.25 {
		direction = 3

	} else if angle <= 78.75 {
		direction = 4

	} else if angle <= 101.25 {
		direction = 5

	} else if angle <= 123.75 {
		direction = 6

	} else if angle <= 146.25 {
		direction = 7

	} else if angle <= 168.75 {
		direction = 8

	} else if angle <= 191.25 {
		direction = 9

	} else if angle <= 213.75 {
		direction = 10

	} else if angle <= 236.25 {
		direction = 11

	} else if angle <= 258.75 {
		direction = 12

	} else if angle <= 281.25 {
		direction = 13

	} else if angle <= 303.75 {
		direction = 14

	} else if angle <= 326.25 {
		direction = 15

	} else if angle <= 348.75 {
		direction = 16

	} else {
		direction = 1
	}

	return direction
}
