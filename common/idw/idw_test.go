package idw

import (
	"testing"

	"gitee.com/wuxi_jiufang/jf-go-kit/base"
)

// 功能测试
func TestIwdIstance(t *testing.T) {
	x := 124.41
	y := 32.2
	points := []IdwPoint{}
	points = append(points, IdwPoint{X: 124.416664, Y: 32.25, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.333336, Y: 32.166668, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.416664, Y: 32.166668, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.333336, Y: 32.25, Weight: 0.34})

	value := IwdIstance(x, y, points)
	base.Glog.Infoln(">>> value: ", value)
	if value != 0.34 {
		t.Errorf("Add func test error, sum = %v", value)
	}
}

// 基准测试
/**
ns/op 代表每次执行逻辑消耗的时间
B/op 代表每次执行逻辑消耗的内存
allocs/op代表每次执行逻辑申请内存的次数
*/
func BenchmarkIwdIstance(b *testing.B) {
	x := 124.41
	y := 32.2
	points := []IdwPoint{}
	points = append(points, IdwPoint{X: 124.416664, Y: 32.25, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.333336, Y: 32.166668, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.416664, Y: 32.166668, Weight: 0.34})
	points = append(points, IdwPoint{X: 124.333336, Y: 32.25, Weight: 0.34})

	for i := 0; i < b.N; i++ {
		value := IwdIstance(x, y, points)
		if value != 0.34 {
			b.Errorf("Add func test error, sum = %v", value)
		}
	}

}
