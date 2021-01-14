/*
@Time : 2021/1/6 5:01 下午
@Author : 21
@File : define
@Software: GoLand
*/
package rabbitmq

var ConnUrl = "amqp://guest:guest@127.0.0.1:5672/"


// mq 的 基类结构体
type MqBase struct {
	ExchangeName string
	ExchangeType string
	QueueName string
	RoutingKey string
}
