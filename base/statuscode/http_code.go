/*
 * @Author: lisheng
 * @Date: 2022-11-02 22:59:39
 * @LastEditTime: 2022-11-03 00:48:23
 * @LastEditors: lisheng
 * @Description: http状态码
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/base/statuscode/http_code.go
 */
package statuscode

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

var (
	OK            = &Status{1, "请求成功"}
	ERROR_DEFAULT = &Status{-1, "请求失败"}
)
