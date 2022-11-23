/*
 * @Author: lisheng
 * @Date: 2022-11-20 00:05:12
 * @LastEditTime: 2022-11-20 08:56:29
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/database/rabbitmq/record.go
 */
package rabbitmq

import (
	"jf-go-kit/base"

	"github.com/streadway/amqp"
)

/**
 * @description: 生成消息
 * @param {*} engineName
 * @param {*} exchange
 * @param {string} key
 * @param {*} deliverymode
 * @param {uint8} priority
 * @param {string} body
 * @return {*}
 * @author: liqiyuWorks
 */
func PublishQueue(engineName, exchange, key string, deliverymode, priority uint8, body string) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	err = GRabbitMQManager.EngineMap[engineName].channel.Publish(exchange, key, false, false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    deliverymode,
			Priority:        priority,
			Body:            []byte(body),
		},
	)
	if err != nil {
		base.Glog.Errorf("[amqp] publish message error: %s\n", err)
		return err
	}
	return nil
}

func DeclareExchange(engineName, name, typ string, durable, autodelete, nowait bool) (err error) {
	err = GRabbitMQManager.EngineMap[engineName].channel.ExchangeDeclare(name, typ, durable, autodelete, false, nowait, nil)
	if err != nil {
		base.Glog.Errorf("[amqp] declare exchange error: %s\n", err)
		return err
	}
	return nil
}

func DeleteExchange(engineName, name string) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	err = GRabbitMQManager.EngineMap[engineName].channel.ExchangeDelete(name, false, false)
	if err != nil {
		base.Glog.Errorf("[amqp] delete exchange error: %s\n", err)
		return err
	}
	return nil
}

/**
 * @description: DeclareQueue 创建队列
 * @param {string} name
 * @param {*} durable
 * @param {*} autodelete
 * @param {*} exclusive
 * @param {bool} nowait
 * @return {*}
 * @author: liqiyuWorks
 */
func DeclareQueue(engineName, name string, durable, autodelete, exclusive, nowait bool) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	_, err = GRabbitMQManager.EngineMap[engineName].channel.QueueDeclare(name, durable, autodelete, exclusive, nowait, nil)
	if err != nil {
		base.Glog.Errorf("[amqp] declare queue error: %s\n", err)
		return err
	}
	return nil
}

/**
 * @description: DeleteQueue 删除队列
 * @param {*} engineName
 * @param {string} name
 * @return {*}
 * @author: liqiyuWorks
 */
func DeleteQueue(engineName, name string) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	_, err = GRabbitMQManager.EngineMap[engineName].channel.QueueDelete(name, false, false, false)
	if err != nil {
		base.Glog.Errorf("[amqp] delete queue error: %s\n", err)
		return err
	}
	return nil
}

func BindQueue(engineName, queue, exchange string, keys []string, nowait bool) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	for _, key := range keys {
		if err = GRabbitMQManager.EngineMap[engineName].channel.QueueBind(queue, key, exchange, nowait, nil); err != nil {
			base.Glog.Errorf("[amqp] bind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

func UnBindQueue(engineName, queue, exchange string, keys []string) (err error) {
	if engineName == "" {
		engineName = "default"
	}
	for _, key := range keys {
		if err = GRabbitMQManager.EngineMap[engineName].channel.QueueUnbind(queue, key, exchange, nil); err != nil {
			base.Glog.Errorf("[amqp] unbind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

/**
 * @description: ConsumeQueue - 消费者
 * @param {*} engineName
 * @param {string} queue
 * @return {*}
 * @author: liqiyuWorks
 */
func ConsumeQueue(engineName, queue string) (deliveries <-chan amqp.Delivery, err error) {
	if engineName == "" {
		engineName = "default"
	}
	deliveries, err = GRabbitMQManager.EngineMap[engineName].channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		base.Glog.Errorf("[amqp] consume queue error: %s\n", err)
		return nil, err
	}
	return deliveries, nil
}
