/*
 * @Author: lisheng
 * @Date: 2022-11-15 16:32:01
 * @LastEditTime: 2023-01-06 11:20:29
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /go-kit/base/request/request.go
 */
package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/**
* @description: GET方式的http请求
* @param {string} url 请求地址
* @param {map[string]string} headers 请求header头设置
* @param {*} params 请求地址栏后需要拼接参数操作
* @param {interface{}} resDataPtr 需返回绑定数据
* @return {*} 返回类型 "错误信息"
* @author: liqiyuWorks
 */
func GET(url string, headers, params map[string]string, resDataPtr interface{}) error {
	client := &http.Client{
		Timeout: 5 * time.Second, // 超时时间：5秒
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New("new request is fail: %v")
	}

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

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &resDataPtr)

	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: POST方式的http请求
 * @param {string} url 请求链接
 * @param {map[string]string} headers 请求header头设置
 * @param {map[string]string} reqDataMap 请求的data map
 * @param {interface{}} resDataPtr 需返回绑定数据
 * @return {*} 错误信息
 * @author: liqiyuWorks
 */
func POST(url string, headers map[string]string, reqDataMap map[string]string, resDataPtr interface{}) error {
	client := &http.Client{
		Timeout: 5 * time.Second, // 超时时间：5秒
	}
	reqData, _ := json.Marshal(reqDataMap)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqData))
	if err != nil {
		return errors.New("new request is fail: %v")
	}
	req.Header.Set("Content-type", "application/json")

	//add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &resDataPtr)

	if err != nil {
		return err
	}

	return nil
}

func GETStatusCode(url string) int {
	client := &http.Client{
		Timeout: 5 * time.Second, // 超时时间：5秒
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 500
	}

	resp, err := client.Do(req)
	if err != nil {
		return 500
	}

	return resp.StatusCode
}
