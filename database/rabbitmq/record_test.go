package rabbitmq

import (
	"testing"

	"gitee.com/wuxi_jiufang/jf-go-kit/config"
)

func InitRecord() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitRabbitMQClient()
}

func TestPublishQueue(t *testing.T) {
	defer InitRecord()()
	PublishQueue("", "", "hello", 0, 0, "test")
	t.Logf("TestConnectRabbitMQ 成功")
}

func TestConsumeQueue(t *testing.T) {
	defer InitRecord()()
	// ConsumeQueue("", "hello")
	t.Logf("TestConsumeQueue 成功")
}
