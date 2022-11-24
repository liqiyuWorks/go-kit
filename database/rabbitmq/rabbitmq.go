/*
 * @Author: lisheng
 * @Date: 2022-11-15 15:36:07
 * @LastEditTime: 2022-11-24 16:28:20
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/database/rabbitmq/rabbitmq.go
 */
package rabbitmq

import (
	"fmt"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base/statuscode"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"github.com/streadway/amqp"
)

// RabbitMQ Operate Wrapper
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
}

type RabbitMQManager struct {
	EngineMap map[string]*RabbitMQ
}

var (
	GRabbitMQManager *RabbitMQManager = new(RabbitMQManager)
)

func CreateDBEngnine(amqpuri string) (*RabbitMQ, error) {
	var rabbitmq *RabbitMQ = new(RabbitMQ)
	var err error
	rabbitmq.conn, err = amqp.Dial(amqpuri)
	if err != nil {
		base.Glog.Errorf("[amqp] connect error: %s\n", err)
		return nil, err
	}
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		base.Glog.Errorf("[amqp] get channel error: %s\n", err)
		return nil, err
	}
	fmt.Printf("rabbitmq %s connect ok ......", amqpuri)
	rabbitmq.done = make(chan error)
	return rabbitmq, nil
}

func InitRabbitMQClient() func() error {
	GRabbitMQManager.EngineMap = make(map[string]*RabbitMQ)
	for k, v := range config.C.Rabbitmq {
		engine, _ := CreateDBEngnine(v)
		GRabbitMQManager.EngineMap[k] = engine
	}

	return func() error {
		for k, engine := range GRabbitMQManager.EngineMap {
			err := engine.Close()
			if err != nil {
				base.Glog.Errorf("> rabbitmqName= %s: %s\n", k, statuscode.ERROR_REDIS_CLOSE.Msg)
			} else {
				fmt.Printf("> rabbitmqName= %s close ok!\n", k)
			}
		}
		return nil
	}
}

/**
 * @description: 关闭 rabbitMQ connection
 * @return {*}
 * @author: liqiyuWorks
 */
func (r *RabbitMQ) Close() (err error) {
	err = r.conn.Close()
	if err != nil {
		base.Glog.Errorf("[amqp] close error: %s\n", err)
		return err
	}
	return nil
}
