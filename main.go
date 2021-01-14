/*
@Time : 2021/1/6 10:57 上午
@Author : 21
@File : main
@Software: GoLand
*/
package main

import (
	"go_wx_work_bot/rabbitmq"
	"go_wx_work_bot/service/wework"
)

func main()  {
	runChanel := make(chan int)
	go listerMqMessages()
	<- runChanel
}

func listerMqMessages()  {
	mq := wework.GetMessageMq()
	MessageService := wework.NewMessageService()
	rabbitmq.ConsumeMessage(mq.MqBase,MessageService.MessageCallback)
}

