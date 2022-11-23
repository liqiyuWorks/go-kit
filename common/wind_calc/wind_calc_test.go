package windcalc

import (
	"testing"

	"jf-go-kit/base"
)

// 功能测试
func TestNewWind(t *testing.T) {
	u := -0.21
	v := -0.15
	wind := NewWind(u, v)
	angle := wind.WindAngle()
	direction := wind.WindDirection()
	speed := wind.WindSpeed()
	base.Glog.Infoln(angle, direction, speed)
	if angle != 54.46 && angle == 3 && speed == 0.3 {
		t.Errorf("TestNewWind test error, error...")
	}
}
