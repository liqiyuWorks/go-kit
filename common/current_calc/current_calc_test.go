package currentcalc

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

// 功能测试
func TestNewCurrent(t *testing.T) {
	u := -0.21
	v := -0.15
	azimuth := 20

	wind := NewCurrent(u, v)
	degree := wind.CurrentAngle()
	direction := wind.CurrentDirection(degree)
	speed := wind.CurrentSpeed()
	knots := wind.CurrentKnots(speed)
	factor := wind.CurrentFactor(azimuth, speed, degree)
	base.Glog.Infof("degree=%f, direction=%s, speed=%f, knots=%f, factor=%f", degree, direction, speed, knots, factor)
}
