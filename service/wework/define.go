/*
@Time : 2021/1/7 10:28 上午
@Author : 21
@File : define
@Software: GoLand
*/
package wework

import "go_wx_work_bot/rabbitmq"

var (
	MT_APP_READY_MSG        = 11024
	MT_PARAMS_ERROR_MSG     = 11025
	MT_USER_LOGIN           = 11026
	MT_USER_LOGOUT          = 11027
	MT_LOGIN_QRCODE_MSG     = 11028
	MT_SEND_TEXT_MSG        = 11029
	MT_SEND_IMAGE_MSG       = 11030
	MT_SEND_FILE_MSG        = 11031
	MT_SEND_LINK_MSG        = 11033
	MT_SEND_VIDEO_MSG       = 11067
	MT_SEND_PERSON_CARD_MSG = 11034
	MT_RECV_TEXT_MSG        = 11041
	MT_RECV_IMG_MSG         = 11042
	MT_RECV_VIDEO_MSG       = 11043
	MT_RECV_VOICE_MSG       = 11044
	MT_RECV_FILE_MSG        = 11045
	MT_RECV_LOCATION_MSG    = 11046
	MT_RECV_LINK_CARD_MSG   = 11047
	MT_RECV_EMOTION_MSG     = 11048
	MT_RECV_RED_PACKET_MSG  = 11049
	MT_RECV_PERSON_CARD_MSG = 11050
	MT_RECV_OTHER_MSG       = 11051
)

// 企业微信消息的结构体
type Message struct {
	ClientId    int                    `json:"client_id"`
	MessageType int                    `json:"message_type"`
	MessageData map[string]interface{} `json:"message_data"`
}

// 发送到消息队列的结构体
type PushMessage struct {
	ClientId       int    `json:"client_id"`
	ConversationId string `json:"conversation_id"`
	MessageType    int    `json:"message_type"`
	Content        string `json:"content"`
	Row            string `json:"row"`
}

// 企业微信 rabbitmq 结构体
type MessageMq struct {
	rabbitmq.MqBase
}

// 获取消息的rabbitmq
func GetMessageMq() MessageMq {
	return MessageMq{rabbitmq.MqBase{
		ExchangeName: "wx_work_exchange",
		ExchangeType: "direct",
		QueueName:    "wx_word_message",
		RoutingKey:   "wx_word_message",
	}}
}

// 推送消息的rabbit mq
func PushMessageMq() MessageMq {
	return MessageMq{rabbitmq.MqBase{
		ExchangeName: "wx_work_push_exchange",
		ExchangeType: "direct",
		QueueName:    "wx_word_push_message",
		RoutingKey:   "wx_word_push_message",
	}}
}
