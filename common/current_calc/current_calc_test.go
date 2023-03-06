package currentcalc

import (
	"testing"

	"gitee.com/liqiyuworks/go-kit/base"
)

// 功能测试
func TestNewCurrent(t *testing.T) {
	u := 0.4404
	v := 0.1136
	azimuth := 20.0

	current := NewCurrent(u, v)
	angle := current.CurrentAngle()
	direction := current.CurrentDirection(angle)
	speed := current.CurrentSpeed()
	knots := current.CurrentKnots(speed)
	factor := current.CurrentFactor(azimuth, knots, angle)
	base.Glog.Infof("angle=%f, direction=%s, speed=%f, knots=%f, factor=%f", angle, direction, speed, knots, factor)
}
