/*
 * @Author: lisheng
 * @Date: 2022-11-02 22:59:39
 * @LastEditTime: 2023-01-06 11:07:08
 * @LastEditors: lisheng
 * @Description: http状态码
 * @FilePath: /go-kit/base/statuscode/http_code.go
 */
package statuscode

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

var (
	OK    = &Status{1, "请求成功"}
	ERROR = &Status{-1, "请求失败"}
)
