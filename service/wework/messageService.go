/*
@Time : 2021/1/7 10:26 上午
@Author : 21
@File : MessageService
@Software: GoLand
*/
package wework

import (
	"encoding/json"
	"fmt"
	"log"
	"go_wx_work_bot/helper"
	"go_wx_work_bot/rabbitmq"
	"runtime/debug"
	"strings"
)

type MessageService struct {

}

func NewMessageService() MessageService {
	return MessageService{}
}

// 消息回调系统
func (m *MessageService) MessageCallback(message []byte)  {
	defer func() {
		if err := recover(); err!= nil{
			log.Print(string(debug.Stack()))
			log.Print(err)
		}
	}()
	//消息解析成结构体
	messageModel := Message{}
	err := json.Unmarshal(message, &messageModel)
	if err != nil{
		fmt.Println(err)
	}
	//判断是否存在对话id
	if _ , ok := messageModel.MessageData["conversation_id"] ; !ok {
		fmt.Println("不属于聊天类型的", string(message))
		return
	}
	msg := messageModel.MessageData["content"].(string)

	//消息类型判断
	if messageModel.MessageType == MT_RECV_TEXT_MSG {
		if (strings.Index(msg, "查询天气:")) >= 0 || (strings.Index(msg, "查询天气：")) >= 0 {
			m.QueryWeather(messageModel, msg)
			return
		}
		message , err := helper.RootMessage(msg)
		if err != nil{
			m.QueryDefault(messageModel, err.Error())
			return
		}
		m.QueryDefault(messageModel, message)
		return
	}
}

func (m *MessageService) QueryWeather(message Message, content string)  {
	city  := strings.Replace(content,"查询天气：","",-1)
	city  = strings.Replace(content,"查询天气:","",-1)
	if len(city) <=0{
		m.QueryDefault(message,"城市为空")
		return
	}
	weatherApiRes, err := helper.GetWeather(city)
	if err !=  nil{
		m.QueryDefault(message,err.Error())
		return
	}

	msg := m.FormatWeather(weatherApiRes)

	m.QueryDefault(message, msg)
	return
}

func (m *MessageService) QueryDefault(message Message,displayMsg string )  {
	pushMsg := "指令有误请输入查询天气：深圳"
	if len(displayMsg) > 1{
		pushMsg = displayMsg
	}
	ConversationId := message.MessageData["conversation_id"].(string)
	msg := PushMessage{
				ClientId:       message.ClientId,
				ConversationId: ConversationId,
				MessageType:    MT_RECV_TEXT_MSG,
				Content:        pushMsg,
				Row:            "111",
	}
	m.PushMessage(msg)
}


/*
	天气格式化成字符串
*/
func(m *MessageService) FormatWeather(res helper.WeatherApiRes)  string {
	// 取出今日的天气
	yesterday := res.Data["yesterday"].(map[string]interface{})
	yesterdayStr := fmt.Sprintf("今日天气%s,%s,%s,风向为%s\r\n", yesterday["type"],yesterday["high"], yesterday["low"], yesterday["fx"])
	forecast := res.Data["forecast"].([]interface{})
	forecastStr := ""
	for _,v := range forecast{
		t := v.(map[string]interface{})
		temp := fmt.Sprintf("%s天气%s,%s,%s,风向为%s\r\n",t["date"], t["type"], t["high"], t["low"], t["fengxiang"])
		forecastStr = forecastStr + temp
	}
	return 	res.Data["city"].(string) + yesterdayStr + forecastStr + res.Data["ganmao"].(string)
}

func (m *MessageService) PushMessage(pushMsg PushMessage){
	jsonInfo , err := json.Marshal(pushMsg)
	fmt.Println(err)
	mq := PushMessageMq()
	rabbitmq.Produce(mq.MqBase, jsonInfo)
}


func (m MessageService) FormatExpress(res helper.ExpressApiRes)  {
	fmt.Println(m)
}
