package currentcalc

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

// 功能测试
func TestNewCurrent(t *testing.T) {
	u := 0.4404
	v := 0.1136
	azimuth := 20

	wind := NewCurrent(u, v)
	angle := wind.CurrentAngle()
	direction := wind.CurrentDirection(angle)
	speed := wind.CurrentSpeed()
	knots := wind.CurrentKnots(speed)
	factor := wind.CurrentFactor(azimuth, knots, angle)
	base.Glog.Infof("angle=%f, direction=%s, speed=%f, knots=%f, factor=%f", angle, direction, speed, knots, factor)
}
