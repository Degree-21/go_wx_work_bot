/*
@Time : 2021/1/7 10:47 上午
@Author : 21
@File : req
@Software: GoLand
*/
package lib

import (
	"io/ioutil"
	"net/http"
)

func SendGet(url string, header map[string]string) (body []byte, err error) {
	client := &http.Client{}
	var req *http.Request

	req, _ = http.NewRequest("GET", url,nil)
	for k, v := range header {
		req.Header.Add(k,v)
	}
	resp, err := client.Do(req)
	if err != nil{
		return  nil, err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil


}