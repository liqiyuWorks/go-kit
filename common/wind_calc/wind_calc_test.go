package windcalc

import (
	"testing"

	"gitee.com/liqiyuworks/go-kit/base"
)

// 功能测试
func TestNewWind(t *testing.T) {
	u := -4.6951
	v := -0.1984
	wind := NewWind(u, v)
	degree := wind.WindAngle()
	direction := wind.WindDirection(degree)
	speed := wind.WindSpeed()
	knots := wind.WindKnots(speed)
	pf := wind.WindBeaufortWindForceScale(speed)
	base.Glog.Infof("degree=%f, direction=%s, speed=%f, knots=%f, pf=%d", degree, direction, speed, knots, pf)
}
