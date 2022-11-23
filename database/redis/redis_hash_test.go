package redis

import (
	"jf-go-kit/config"
	"testing"

	"jf-go-kit/base"
)

func InitHashRds() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitRedisClient()
}

func TestHashGetMap(t *testing.T) {
	defer InitHashRds()()
	key := "liqiyu"
	field := "lisheng"
	value, err := HashGetFields("", key, field)
	base.Glog.Infoln("value =", value)
	if err != nil {
		t.Errorf("data = %v", value)
	}
}
