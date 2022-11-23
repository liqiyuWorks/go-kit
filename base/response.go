/*
 * @Author: lisheng
 * @Date: 2022-10-11 14:29:08
 * @LastEditTime: 2022-11-02 23:02:33
 * @LastEditors: lisheng
 * @Description: 通用json返回模块
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/base/response.go
 */
package base

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}
