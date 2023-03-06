package wavecalc

import (
	"testing"

	"gitee.com/liqiyuworks/go-kit/base"
)

// 功能测试
func TestNewWave(t *testing.T) {
	degree := 340.72
	height := 1.66
	period := 6.13
	wave := NewWave(degree, height, period)
	direction := wave.WaveDirection()
	douglasScale := wave.WaveDouglasScale()
	base.Glog.Infof("douglasScale=%s, direction=%s", douglasScale, direction)
}
