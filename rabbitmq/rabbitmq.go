/*
@Time : 2021/1/6 5:05 下午
@Author : 21
@File : rabbitmq
@Software: GoLand
*/
package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel
var notifyClose  chan *amqp.Error

/*
	初始化 rabbitmq channel
*/
func initChanel()  bool {
	var err error
	if channel !=nil {
		return true
	}

	// todo 可以更换成 conf 读取
	conn , err = amqp.Dial(ConnUrl)
	if err != nil{
		fmt.Println(err)
		return false
	}

	channel , err = conn.Channel()
	if err != nil{
		fmt.Println(err)
		return false
	}
	return true
}

func Produce(base MqBase,message []byte)  {
	if channel == nil{
		initChanel()
	}
	var err error
	// 创建交换机
	err = channel.ExchangeDeclare(
		base.ExchangeName,
		base.ExchangeType,
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,     // arguments
	)

	if err != nil{
		fmt.Println("创建交换机失败" + err.Error())
		return
	}
	//	创建queue
	q, err := channel.QueueDeclare(
		base.QueueName,
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil{
		fmt.Println("创建queue失败:"+err.Error())
		return
	}
	//	绑定queue
	err = channel.QueueBind(q.Name,
		base.RoutingKey,
		base.ExchangeName,
		false,
		nil,
	)
	if err != nil{
		fmt.Println("绑定queue失败")
		return
	}

	err = channel.Publish(
		base.ExchangeName,
		base.RoutingKey,
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("成功发送")


}


func ConsumeMessage(base MqBase, callback func(v []byte))  {
	if channel == nil{
		initChanel()
	}
	var err error
	// 创建交换机
	err = channel.ExchangeDeclare(
		base.ExchangeName,
		base.ExchangeType,
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,     // arguments
		)

	if err != nil{
		fmt.Println("创建交换机失败" + err.Error())
		return
	}
	//	创建queue
	q, err := channel.QueueDeclare(
		base.QueueName,
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil{
		fmt.Println("创建queue失败:"+err.Error())
		return
	}
	//	绑定queue
	err = channel.QueueBind(q.Name,
		base.RoutingKey,
		base.ExchangeName,
		false,
		nil,
	)
	if err != nil{
		fmt.Println("绑定queue失败")
		return
	}

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)
	go func() {
		for v := range msgs{
			callback(v.Body)
		}
	}()
	<-forever
}


