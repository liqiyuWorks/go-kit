package redis

import (
	"testing"

	"gitee.com/liqiyuworks/go-kit/config"

	"gitee.com/liqiyuworks/go-kit/base"
)

func InitHashRds() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitRedisClient()
}

func TestHashGetMap(t *testing.T) {
	defer InitHashRds()()
	key := "gfs|547|280"
	field := "2021080400"
	value, _ := HashGetFields("", key, field)
	base.Glog.Infoln("value =", value)
}
