package redis

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

func InitStringRds() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	base.InitLogger()
	return InitRedisClient()
}

func TestStringGet(t *testing.T) {
	defer InitStringRds()()

	key := "lisheng"
	value, err := StringGet("default", key)
	base.Glog.Infoln(value[key])
	if err != nil {
		t.Errorf("data = %v", value)
	}

}
