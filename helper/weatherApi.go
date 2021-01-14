/*
@Time : 2021/1/7 10:53 上午
@Author : 21
@File : weatherApi
@Software: GoLand
*/
package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_wx_work_bot/lib"
)

type WeatherApiRes struct {
	Data map[string]interface{} `json:"data"`
	Status int `json:"status"`
	Desc string `json:"desc"`
}

func GetWeather(city string) (weatherJson WeatherApiRes, err error) {
	url := fmt.Sprintf("http://wthrcdn.etouch.cn/weather_mini?city=%s", city)

	// 发送api请求
	res , err := lib.SendGet(url, map[string]string{})
	if err != nil{
		return weatherJson , err
	}
	apiRes := WeatherApiRes{}
	err = json.Unmarshal(res, &apiRes)

	if err != nil {
		return weatherJson , err
	}

	if apiRes.Status != 1000 {
		return weatherJson, errors.New("获取天气预报失败！")
	}
	return apiRes ,nil
}
