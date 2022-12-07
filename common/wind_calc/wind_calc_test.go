package windcalc

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

// 功能测试
func TestNewWind(t *testing.T) {
	u := -0.21
	v := -0.15
	wind := NewWind(u, v)
	degree := wind.WindAngle()
	direction := wind.WindDirection(degree)
	speed := wind.WindSpeed()
	knots := wind.WindKnots(speed)
	pf := wind.WindPf(speed)
	base.Glog.Infof("degree=%f, direction=%s, speed=%f, knots=%f, pf=%d", degree, direction, speed, knots, pf)
}
