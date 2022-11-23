/*
 * @Author: lisheng
 * @Date: 2022-11-15 16:32:01
 * @LastEditTime: 2022-11-20 12:25:52
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/base/request/request.go
 */
package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @description: 支持任意方式的http请求
 * @param {*} url 请求地址
 * @param {string} method 请求方式
 * @param {*} params 请求地址栏后需要拼接参数操作
 * @param {map[string]string} headers 请求header头设置
 * @param {[]byte} data
 * @return {*} 返回类型 "错误信息"
 * @author: liqiyuWorks
 */
func HttpDo(url, method string, params, headers map[string]string, data []byte) (interface{}, error) {

	//自定义cient
	client := &http.Client{
		Timeout: 5 * time.Second, // 超时时间：5秒
	}
	//http.post等方法只是在NewRequest上又封装来了一层而已
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("new request is fail: %v")
	}
	req.Header.Set("Content-type", "application/json")

	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	//add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req) //  默认的resp ,err :=  http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	temp := make(map[string]interface{}, 0)

	err = json.Unmarshal(body, &temp)

	if err != nil {
		return nil, err
	}

	return temp, nil
}
