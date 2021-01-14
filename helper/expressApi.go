/*
@Time : 2021/1/7 11:35 上午
@Author : 21
@File : ExpressApi
@Software: GoLand
*/
package helper

import (
	"encoding/json"
	"fmt"
	"go_wx_work_bot/lib"
)

// 快递接口api

var AppCode = "35dd1a2b88e644a88469c3a84aae4fd0"

type ExpressApiRes struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
	Result map[string]interface{} `json:"result"`
}

/*
	快递api查询
*/
func QueryExpress(expressNum string) (dataModel ExpressApiRes, err error) {
	url := fmt.Sprintf("https://jisuexpress.api.bdymkt.com/express/query?type=auto&number=%s&mobile=", expressNum)
	headers := map[string]string{
		"X-Bce-Signature" : "AppCode/" + AppCode,
	}
	res , err := lib.SendGet(url, headers)
	if err != nil{
		return dataModel, err
	}

	resultModel := ExpressApiRes{}

	err = json.Unmarshal(res, &resultModel)
	if err != nil{
		return dataModel, err
	}
	return resultModel , nil
}