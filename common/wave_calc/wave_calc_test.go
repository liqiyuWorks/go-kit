package wavecalc

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

// 功能测试
func TestNewWave(t *testing.T) {
	u := -0.21
	v := -0.15
	height := 0.3
	wave := NewWave(u, v)
	douglasScale := wave.WaveDouglasScale(height)
	base.Glog.Infof("douglasScale=%s", douglasScale)
}
