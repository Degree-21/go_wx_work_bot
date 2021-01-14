/*
@Time : 2021/1/14 2:40 下午
@Author : 21
@File : tianAi
@Software: GoLand
*/
package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_wx_work_bot/lib"
	"net/url"
)

// https://www.tianapi.com/apiview/47

var KEY = "6c80a34954420f65176ec34f1261967e"
var BASE = "http://api.tianapi.com/txapi/robot/index?key=%s&question=%s"

type robotMsg struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	NewsList interface{} `json:"newslist"`
}


func RootMessage(msg string) (reply string , err error) {
	msg = url.QueryEscape(msg)
	url := fmt.Sprintf(BASE,KEY,msg)
	res , err := lib.SendGet(url, map[string]string{})
	if err != nil{
		return reply , err
	}
	resultModel := robotMsg{}
	err = json.Unmarshal(res,&resultModel)
	if err != nil{
		return reply , err
	}
	resultMessage := resultModel.NewsList.([]interface{})
	if len(resultMessage) > 0 {
		replyType := resultMessage[0].(map[string]interface{})
		return replyType["reply"].(string) , nil
	}

	return reply, errors.New("获取不到机器人返回的结果")
}

func formatMessage(msg string) string {

	return msg
}