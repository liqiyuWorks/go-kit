package point

import (
	"fmt"
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/config"
	"gitee.com/liqiyuworks/jf-go-kit/database/redis"
)

func InitPoint() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return redis.InitRedisClient()
}

// 功能测试
func TestConverMfwamPointToIndex(t *testing.T) {
	defer InitPoint()()
	var lons = make(map[string][]string)
	var lats = make(map[string][]string)
	lons, _ = redis.SetZange("points", "base_data:mfwam:lons", 0, -1)
	lats, _ = redis.SetZange("points", "base_data:mfwam:lats", 0, -1)

	lon := 124.41
	lat := 32.2
	lonIndexs, latIndexs, err := ConverMfwamPointToIndex(lons["base_data:mfwam:lons"], lats["base_data:mfwam:lats"], lon, lat)

	fmt.Println(">>> value: ", lonIndexs, latIndexs)
	if err != nil {
		t.Errorf("TestConverMfwamPointToIndex test error, err = %v", err)
	}
}
